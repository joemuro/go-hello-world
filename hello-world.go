package main
import "fmt"
import "os"
import "runtime"

func main() {
    fmt.Println("hello world")
    fmt.Printf("PID: %d\n", os.Getpid())
    fmt.Printf("UID: %d\n", os.Getuid())
    cwd,_ := os.Getwd()
    fmt.Printf("cwd: %s\n", cwd)
    hostname,_ := os.Hostname()
    fmt.Printf("hostname: %s\n",hostname) 
    fmt.Printf("NumCPUs: %d\n",runtime.NumCPU())
    fmt.Printf("golang-version: %s\n", runtime.Version())
}
