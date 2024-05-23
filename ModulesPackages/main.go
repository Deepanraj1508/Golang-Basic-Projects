package main
import (
	"fmt"
	"demo/core"
	"demo/core/add"
	"demo/hello"
)

func main() {
	fmt.Println("Main Module Hello World")
	fmt.Println("")
	fmt.Println("----------------------------")
	core.Printfunc()
	add.Addfunc()
	hello.Datefunc()
	core.Printcore()
	PrintCalc()
}

