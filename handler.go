package main

import (
	"encoding/json"
	"net/http"
	"strconv"
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
  limitStr := r.URL.Query().Get("limit")
  limit := -1
  if limitStr != ""{
    n,err := strconv.Atoi(limitStr)
  if err != nil{
    w.WriteHeader(http.StatusBadRequest)
    return
  }
 limit = n
  }
  
mu.Lock()
result := logs
if limit > 0{
  start := len(logs)-limit
  if start <0 {
    start = 0
  }
  result = logs[start:]
}
data,err := json.Marshal(result)
mu.Unlock()
if err != nil{
  w.WriteHeader(http.StatusInternalServerError)
  return
}
w.WriteHeader(http.StatusOK)
w.Write(data)
}