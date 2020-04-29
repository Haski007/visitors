package main

import (
	"os"
	"log"
	"fmt"
	"net/http"

	"./logger"
	"./handlers"
	"./config"
)

func main() {
	http.HandleFunc("/", logger.LogHandler(handlers.IndexPage))
	http.HandleFunc("/favicon.ico", handlers.DoNothing)

	fmt.Println("Listerning on port:", config.Port)
	err := http.ListenAndServe(config.Port, nil)
	if err != nil {
		log.Println("ERROR:", err.Error())
		os.Exit(-1)
	}
}

func checkError(err error) {
	if err != nil {
		log.Println("ERROR:", err.Error())
	}
}