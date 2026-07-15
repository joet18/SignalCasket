package main

import (
	"net/http"
  "sync"
)

var (
  logs []string;
  mu sync.Mutex
)


func statHandler(w http.ResponseWriter, r *http.Request){
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Signalcasket is running"))
}