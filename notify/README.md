# Notify Service

HTTP-сервис на Go, который слушает порт **30000** и отвечает `200 OK` на эндпоинт `/health`.

## Запуск в Docker

### 1. Сборка образа

```bash
docker build -t notify .
```

### 2. Запуск контейнера с прокидыванием порта

```bash
docker run -d -p 30000:30000 --name notify notify
```

Флаг `-p 30000:30000` пробрасывает порт 30000 контейнера на порт 30000 хоста.

### 3. Проверка

```bash
curl http://localhost:30000/health
```

Ожидаемый ответ: `OK` со статус-кодом `200`.

### 4. Остановка и удаление контейнера

```bash
docker stop notify
docker rm notify
```
