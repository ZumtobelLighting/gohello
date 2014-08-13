package main
import "fmt"
import "github.com/digitallumens/goexamplepackage"
func main() {
    fmt.Println("hello world")
    i := goexamplepackage.DoExample(22)
    fmt.Println("answer:", i)
}
