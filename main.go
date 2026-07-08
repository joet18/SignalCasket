package main
import ("fmt"
        "bufio"
        "io"
        "time"
        "os"
        "context"
        "os/signal"
        "syscall"
    )

     
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
    go func(){
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
   file, err := os.Open(path)
   if err != nil {
       fmt.Println("Error opening file:", err)
       os.Exit(1)
   }
    defer file.Close()
     reader := bufio.NewReader(file)
    MainLoop:
     for{
        
        line, err := reader.ReadString('\n')
        fmt.Println("Line:", line)
        fmt.Println("err:",err)
        if err == io.EOF {
            fmt.Println("waiting formare", err)
             select{
        case <-ctx.Done():
            fmt.Println("Exiting..")
            break MainLoop
        case <-time.After(1 * time.Second):
            fmt.Println("Waiting")
    }
        }else if err != nil {
            fmt.Println("Error reading line:", err)
            os.Exit(1)

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