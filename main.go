package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

     
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    sigCh := make(chan os.Signal, 1) //creates a channel to recive the signal
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)//redirects the signal to the channel
    go func(){ //grotine blocks itself afer recciving the signal
        <-sigCh;
        fmt.Println("shutdown signal recived");cancel()
    }()
   

  fmt.Println("starting tailer..")
   path,err := checkArgs()
   if err != nil {
       fmt.Println("Error:", err)
       os.Exit(1)
   }
   fmt.Println("Tailing file:", path)
   lineCh :=make(chan string)
   go func(){
    http.HandleFunc("/status",statHandler);
    http.ListenAndServe(":8080",nil)
    if err != nil{
        fmt.Println("Http server error:", err)
    }
   }()
   go tailFile(path,lineCh,ctx)
   for{
    select{
    case line:= <-lineCh:
        mu.Lock()
        logs = append(logs, line)
        mu.Unlock()
    fmt.Println("Line:",line)
    case <-ctx.Done():
        fmt.Println("Exiting..")
        return
    }
   }
}
func tailFile(path string,lineCh chan<- string,ctx context.Context){
   file, err := os.Open(path)
   if err != nil {
       fmt.Println("Error opening file:", err)
       os.Exit(1)
   }
    defer file.Close()
    
     reader := bufio.NewReader(file)
    ProducerLoop:
     for{
        
        line, err := reader.ReadString('\n')

        if err == io.EOF {
            fmt.Println("waiting formare", err)
             select{
        case <-ctx.Done():
            fmt.Println("Exiting..")
            break ProducerLoop
        case <-time.After(1 * time.Second):
            fmt.Println("Waiting")
    }
        }else if err != nil {
            fmt.Println("Error reading line:", err)
            os.Exit(1)

        } else {
            lineCh <- line
        }
    }
}
func checkArgs()(string,error){

if  len(os.Args) < 2 {
	fmt.Println("Please provide a file path to tail")
	os.Exit(1)
}
	return os.Args[1], nil
}