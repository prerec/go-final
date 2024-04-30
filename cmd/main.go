package main

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	todo "github.com/prerec/go-final"
	"github.com/prerec/go-final/pkg/handler"
	"github.com/prerec/go-final/pkg/repository"
	"github.com/prerec/go-final/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("init config fail: %s", err.Error())
	}

	db, err := repository.NewSqliteDB(repository.Config{
		Driver:   viper.GetString("db_driver"),
		Port:     viper.GetString("db_port"),
		Username: viper.GetString("db_username"),
		Password: viper.GetString("db_password"),
		DBName:   viper.GetString("db_name"),
		SSLMode:  viper.GetString("db_ssl_mode"),
	})
	if err != nil {
		logrus.Fatalf("init db fail: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running server: %s", err.Error())
		}
	}()

	logrus.Print("Application started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Application shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("server shutdown fail: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Fatalf("database connection close fail: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
