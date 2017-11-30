package main

import (
	"log"
	"net/http"
	"time"
)

func Check(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if db == nil {
			var dbt Database
			start := time.Now()
			//Database{"mysql", "root", "zzhzzhzzh", "jw_crew", "45.62.101.211", "3306", "utf8"}
			db = createInstance(dbt.getConf())
			log.Printf(
				"%s\t%s\t%s\t%s\t%s",
				r.Method,
				r.RequestURI,
				"database have been open by",
				name,
				time.Since(start),
			)
		}
		inner.ServeHTTP(w, r)
	})
}
