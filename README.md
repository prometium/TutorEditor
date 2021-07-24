# TutorEditor

## Использование

### База данных

1. Необходимо установить Docker.
2. Запуск Dgraph в Docker:
   `docker run -it -p 9080:9080 -v /mnt/dgraph:/dgraph dgraph/standalone:v20.03.0`.
   База данных будет доступна по адресу http://localhost:9080/.

### Сервер

1. Необходимо установить Golang.
2. Запуск сервера:
   `go run cmd/main.go "-http.addr" ":9000"`.
   Сервер будет доступен по адресу http://localhost:9000/.

### Клиент

1. Необхомо установить Node.js.
2. Установка зависимостей:
   `npm install`.
3. Запуск сервера для разработки:
   `npm run serve`.
   Приложение будет доступно по адресу http://localhost:3000/.
