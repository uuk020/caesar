package main

import (
	"caesar/config"
	caesarValidate "caesar/controller/validate"
	"caesar/global"
	caesarInternal "caesar/internal"
	caesarRouter "caesar/router"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func initConfig() {
	v := viper.New()
	v.SetConfigFile("./settings.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	server := config.Server{}
	if err := v.Unmarshal(&server); err != nil {
		panic(err)
	}
	global.Setting = server
}

func initMySQL() {
	mysql := global.Setting.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysql.Name, mysql.Password, mysql.Host, mysql.Port, mysql.DbName)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(global.Setting.Mysql.MaxOpenConns)
	db.SetMaxIdleConns(global.Setting.Mysql.MaxIdleConns)
	global.DB = db
}

func initRedis() {
	addr := fmt.Sprintf("%s:%d", global.Setting.Redis.Host, global.Setting.Redis.Port)
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: global.Setting.Redis.Password,
		DB:       0,
	})
	_, err := global.Redis.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func init() {
	// Init config, mysql and redis
	initConfig()
	initMySQL()
	initRedis()
}

func main() {
	// Setup
	e := echo.New()

	// Debug model
	e.Debug = global.Setting.Debug

	// Save log
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:    true,
		LogURI:       true,
		LogMethod:    true,
		LogUserAgent: true,
		LogError:     true,
		LogRemoteIP:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Status != http.StatusOK {
				file := fmt.Sprintf("%s%s.log", global.Setting.LogAddr, caesarInternal.GetNowFormatTodayTime())
				fd, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					return err
				}
				s := "Request: time:" + v.StartTime.Local().String() + ", uri:" + v.URI + ", method:" + v.Method + ", user-agent:" + v.UserAgent + ", remote-ip:" + v.RemoteIP + ", error:" + v.Error.Error() + "\n"
				if _, err := fd.WriteString(s); err != nil {
					return err
				}
				defer fd.Close()
			}
			fmt.Printf("REQUEST: time: %v, uri: %v, status: %v, method: %v, user-agent: %v\n", v.StartTime.Local().String(), v.URI, v.Status, v.Method, v.UserAgent)
			return nil
		},
	}))
	e.Use(middleware.Recover())

	// validate
	caesarValidate.Register(e)

	// Register routers
	caesarRouter.Register(e)

	// Start server
	go func() {
		port := fmt.Sprintf(":%d", global.Setting.Port)
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
