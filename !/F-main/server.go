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
	"github.com/shirou/gopsutil/v3/mem"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ramMonitor(w http.ResponseWriter, r *http.Request, ctx context.Context) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Closing WebSocket connection")
			conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			return
		default:
			vmStat, _ := mem.VirtualMemory()
			ramPercent := vmStat.UsedPercent

			msg := map[string]float64{
				"used": ramPercent,
			}

			jsonData, _ := json.Marshal(msg)
			err := conn.WriteMessage(websocket.TextMessage, jsonData)
			if err != nil {
				fmt.Println("Write error:", err)
				return
			}

			time.Sleep(1 * time.Second)
		}
	}
}

func servePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	// Создаем контекст для graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ramMonitor(w, r, ctx)
	})
	http.HandleFunc("/", servePage)

	// Канал для обработки сигналов ОС
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Запуск сервера в отдельной горутине
	go func() {
		fmt.Println("Сервер запущен: http://localhost:8081")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Ошибка сервера: %v\n", err)
		}
	}()

	// Ожидание сигнала для завершения
	<-stop
	fmt.Println("\nПолучен сигнал завершения...")

	// Graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("Ошибка при завершении сервера: %v\n", err)
	}

	// Отменяем контекст для остановки мониторинга RAM
	cancel()
	fmt.Println("Сервер успешно остановлен")
}
