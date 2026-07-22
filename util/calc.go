package calc
import (
    "fmt"
)

//CApitalisation is important in Go. It represents sharability among packages.
func Add(a int, b int){
    fmt.Printf("%d\n", (a+b))
}
func Mul(a int, b int) int {
    return a*b
}
