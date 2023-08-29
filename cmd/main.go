package main

import (
	"context"
	"github.com/AlibekDalgat/dynamic_segmentation"
	"github.com/AlibekDalgat/dynamic_segmentation/pkg/handler"
	"github.com/AlibekDalgat/dynamic_segmentation/pkg/repository"
	"github.com/AlibekDalgat/dynamic_segmentation/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Ошибка при инициализации конфигурации: %s", err.Error())
	}
	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("Ошибка при загрузке переменных окружения: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Ошибка при инициализации базы данных: %v", err)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(dynamic_segmentation.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Ошибка при запуске сервера: %s", err.Error())
		}
	}()
	logrus.Println("Запуск сервиса")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Завершение сервиса")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Ошибка при завершении работы сервера: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("Ошибка при закрытии соединения с БД: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
