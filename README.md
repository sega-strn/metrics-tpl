# Metrics Collection and Alerting Server

Сервер для сбора и мониторинга метрик с функциями алертинга.

## Описание проекта

Этот проект представляет собой систему сбора и мониторинга метрик с возможностью настройки алертов. Система разработана для отслеживания различных показателей работы приложений и инфраструктуры в режиме реального времени.

### Основные возможности

- **Сбор метрик**: поддержка различных типов метрик (счетчики, gauge)
- **Хранение данных**: персистентное хранение метрик с возможностью восстановления
- **Визуализация**: удобный веб-интерфейс с темной темой для просмотра метрик
- **Алертинг**: настраиваемая система оповещений при достижении пороговых значений
- **API**: REST API для интеграции с другими системами

### Технологии

- Go (основной язык разработки)
- PostgreSQL (хранение данных)
- HTTP REST API
- Prometheus-совместимый формат метрик

## Установка и запуск

1. Клонируйте репозиторий:
```bash
git clone <your-repo-url>
```

2. Создайте модуль (если еще не создан):
```bash
go mod init <module-name>
```

3. Установите зависимости:
```bash
go mod tidy
```

4. Запустите сервер:
```bash
go run cmd/server/main.go
```

## API Endpoints

### Метрики

- `POST /update/` - обновление значения метрики
- `GET /value/` - получение значения метрики
- `GET /` - получение всех метрик

### Формат метрик

Поддерживаются следующие типы метрик:
- Counter (счетчик) - только положительные целые числа, можно только увеличивать
- Gauge (измеритель) - число с плавающей точкой, можно как увеличивать, так и уменьшать

## Конфигурация

Сервер может быть настроен через переменные окружения или флаги командной строки:

- `ADDRESS` - адрес сервера (по умолчанию ":8080")
- `STORE_INTERVAL` - интервал сохранения метрик (по умолчанию "300s")
- `STORE_FILE` - путь к файлу для сохранения метрик
- `RESTORE` - восстанавливать ли метрики при запуске (по умолчанию true)

## Разработка

## Начало работы

1. Склонируйте репозиторий в любую подходящую директорию на вашем компьютере.
2. В корне репозитория выполните команду `go mod init <name>` (где `<name>` — адрес вашего репозитория на GitHub без префикса `https://`) для создания модуля.

## Обновление шаблона

Чтобы иметь возможность получать обновления автотестов и других частей шаблона, выполните команду:

```bash
git remote add -m main template https://github.com/Yandex-Practicum/go-musthave-metrics-tpl.git
```

Для обновления кода автотестов выполните команду:

```bash
git fetch template && git checkout template/main .github
```

Затем добавьте полученные изменения в свой репозиторий.

## Запуск автотестов

Для успешного запуска автотестов называйте ветки `iter<number>`, где `<number>` — порядковый номер инкремента. Например, в ветке с названием `iter4` запустятся автотесты для инкрементов с первого по четвёртый.

При мёрже ветки с инкрементом в основную ветку `main` будут запускаться все автотесты.

Подробнее про локальный и автоматический запуск читайте в [README автотестов](https://github.com/Yandex-Practicum/go-autotests).
