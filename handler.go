package main

import (
	"encoding/json"
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
func logsHandler(w http.ResponseWriter, r *http.Request){
mu.Lock()
data,err := json.Marshal(logs)
mu.Unlock()
if err != nil{
  w.WriteHeader(http.StatusInternalServerError)
  return
}
w.WriteHeader(http.StatusOK)
w.Write(data)
}