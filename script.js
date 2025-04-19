document.addEventListener('DOMContentLoaded', function() {
    // Инициализация графиков
    const ctx = document.getElementById('usage-chart').getContext('2d');
    const chart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: [],
            datasets: [
                {
                    label: 'CPU %',
                    data: [],
                    borderColor: '#e74c3c',
                    tension: 0.1,
                    yAxisID: 'y'
                },
                {
                    label: 'RAM %',
                    data: [],
                    borderColor: '#3498db',
                    tension: 0.1,
                    yAxisID: 'y'
                }
            ]
        },
        options: {
            responsive: true,
            interaction: {
                mode: 'index',
                intersect: false,
            },
            scales: {
                y: {
                    type: 'linear',
                    display: true,
                    position: 'left',
                    min: 0,
                    max: 100
                }
            }
        }
    });

    // WebSocket соединение
    const socket = new WebSocket('ws://' + window.location.host + '/ws');
    
    socket.onopen = function() {
        console.log('WebSocket соединение установлено');
    };
    
    socket.onmessage = function(event) {
        const data = JSON.parse(event.data);
        updateMetrics(data);
        updateChart(data, chart);
    };
    
    socket.onclose = function() {
        console.log('WebSocket соединение закрыто');
        document.getElementById('cpu-usage').textContent = '--% (отключено)';
        document.getElementById('ram-usage').textContent = '--% (отключено)';
    };

    function updateMetrics(data) {
        document.getElementById('cpu-usage').textContent = data.CPUUsage.toFixed(1) + '%';
        document.getElementById('cpu-temp').textContent = data.CPUTemp.toFixed(1);
        
        document.getElementById('ram-usage').textContent = data.RAMUsage.toFixed(1) + '%';
        document.getElementById('ram-used').textContent = data.RAMUsed;
        document.getElementById('ram-total').textContent = data.RAMTotal;
    }

    function updateChart(data, chart) {
        const now = new Date(data.Timestamp * 1000).toLocaleTimeString();
        
        // Добавляем новые данные
        chart.data.labels.push(now);
        chart.data.datasets[0].data.push(data.CPUUsage);
        chart.data.datasets[1].data.push(data.RAMUsage);
        
        // Ограничиваем количество точек на графике
        if (chart.data.labels.length > 15) {
            chart.data.labels.shift();
            chart.data.datasets[0].data.shift();
            chart.data.datasets[1].data.shift();
        }
        
        chart.update();
    }
});
