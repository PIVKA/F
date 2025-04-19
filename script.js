document.addEventListener('DOMContentLoaded', function() {
    const container = document.getElementById('metrics-container');
    
    // Показываем красивый индикатор загрузки
    container.innerHTML = `
        <div class="loading-message">
            <span class="loading-spinner"></span>
            Загрузка данных мониторинга...
        </div>
    `;

    function updateMetrics() {
        fetch('/api/metrics')
            .then(response => {
                if (!response.ok) throw new Error(`Ошибка сервера: ${response.status}`);
                return response.json();
            })
            .then(data => {
                container.innerHTML = `
                    <div class="metric">
                        <h2><span class="cpu-indicator">CPU:</span> ${data.cpu.usage}%</h2>
                        <p><span class="temp-indicator">Температура:</span> ${data.cpu.temperature}°C</p>
                    </div>
                    <div class="metric">
                        <h2><span class="gpu-indicator">GPU:</span> ${data.gpu.usage}%</h2>
                        <p><span class="temp-indicator">Температура:</span> ${data.gpu.temperature}°C</p>
                    </div>
                `;
            })
            .catch(error => {
                console.error('Ошибка:', error);
                container.innerHTML = `
                    <div class="error-message">
                        <p>Ошибка загрузки данных</p>
                        <p><small>${error.message}</small></p>
                        <button onclick="window.location.reload()">Обновить</button>
                    </div>
                `;
            });
    }

    setInterval(updateMetrics, 2000);
    updateMetrics();
});