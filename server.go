package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type SystemMetrics struct {
	RAMUsage   float64 `json:"ram_usage"`
	CPUUsage   float64 `json:"cpu_usage"`
	Timestamp  int64   `json:"timestamp"`
	RAMTotal   uint64  `json:"ram_total"`
	RAMUsed    uint64  `json:"ram_used"`
	CPUTemp    float64 `json:"cpu_temp"` // Добавим температуру (заглушка)
}

func getCPULoad() (float64, error) {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	if len(percentages) > 0 {
		return percentages[0], nil
	}
	return 0, fmt.Errorf("no CPU data available")
}

func monitorSystem(ctx context.Context, conn *websocket.Conn) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			vmStat, err := mem.VirtualMemory()
			if err != nil {
				fmt.Println("RAM monitoring error:", err)
				continue
			}

			cpuPercent, err := getCPULoad()
			if err != nil {
				fmt.Println("CPU monitoring error:", err)
				cpuPercent = 0
			}

			metrics := SystemMetrics{
				RAMUsage:  vmStat.UsedPercent,
				CPUUsage:  cpuPercent,
				Timestamp: time.Now().Unix(),
				RAMTotal:  vmStat.Total / (1024 * 1024), // в MB
				RAMUsed:   vmStat.Used / (1024 * 1024),  // в MB
				CPUTemp:   65.3, // Заглушка, замените на реальные данные
			}

			if err := conn.WriteJSON(metrics); err != nil {
				fmt.Println("Write error:", err)
				return
			}
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request, ctx context.Context) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	monitorSystem(ctx, conn)
}

func main() {
	// Настройка обработчиков
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(w, r, context.Background())
	})

	server := &http.Server{Addr: ":8080"}

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("Сервер запущен: http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Ошибка сервера: %v\n", err)
		}
	}()

	<-stop
	fmt.Println("\nПолучен сигнал завершения...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Ошибка при завершении сервера: %v\n", err)
	}

	fmt.Println("Сервер успешно остановлен")
}
