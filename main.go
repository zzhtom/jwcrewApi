package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	router := NewRouter()
	srv := &http.Server{
		Addr:         ":8081",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	var dbt Database
	db = createInstance(dbt.getConf())
	log.Fatal(srv.ListenAndServe())
	//	log.Fatal(http.ListenAndServe(":8081", router))
	defer closeDatabase()
}
