package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/prerec/go-final/pkg/handler"
	"github.com/prerec/go-final/pkg/repository"
	"github.com/prerec/go-final/pkg/server"
	"github.com/prerec/go-final/pkg/service"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occurred while running server: %s", err.Error())
		}
	}()

	logrus.Print("Application started")
	logrus.Printf("Server listening on port %s", viper.GetString("port"))

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Application shutting down")
	go func() {
		if err := srv.Shutdown(context.Background()); err != nil {
			logrus.Fatalf("server shutdown fail: %s", err.Error())
		}
		if err := db.Close(); err != nil {
			logrus.Fatalf("database connection close fail: %s", err.Error())
		}
	}()

	logrus.Print("Application gracefully stopped!")

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
