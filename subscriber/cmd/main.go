package main

import (
	"os"
	"os/signal"
	"subscriber/internal/app"
	"syscall"
)

func main() {
	application := app.New()
	go application.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	// todo: 1) сделать валидацию приходящего json
	// todo: 2) добавить конфиги в .env файл
	// todo: 3) записать видео с демонстрацией работы сервиса
	<-stop

	application.Stop()
}
