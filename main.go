
package main

import (
	"fmt"
	"github.com/hepiska/todo-go/routers"
)

func main() {
	fmt.Print("hello go")
	router:=routers.InitRoute()
	router.Run()
}