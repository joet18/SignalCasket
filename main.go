package main
import ("fmt"
        "bufio"
        "io"
        "time"
        "os")


func main() {
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
    for{
        
        line, err := reader.ReadString('\n')
        fmt.Println("Line:", line)
        fmt.Println("err:",err)
        if err == io.EOF {
            fmt.Println("waiting formare", err)
            time.Sleep(1 * time.Second)
            continue
        }
        
        if err != nil {
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