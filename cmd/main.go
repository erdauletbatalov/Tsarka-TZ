package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/erdauletbatalov/tsarka/configs"
	"github.com/erdauletbatalov/tsarka/delivery/web"
	"github.com/gin-gonic/gin"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/docker.toml", "path to config file")
}

func main() {
	flag.Parse()

	logger, err := os.Create("server.log")
	if err != nil {
		log.Println(err)
		return
	}

	config := configs.NewConfig()
	_, err = toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// db, err := postgres.NewPostgresRepository(config)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	router := gin.Default()
	router.Use(gin.LoggerWithWriter(logger))
	web.NewHandler(router)

	server := &http.Server{
		Addr:         config.BindAddr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	log.Println(server.ListenAndServe())
}
