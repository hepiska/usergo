
package main

import (
	"github.com/hepiska/todo-go/routers"
	"github.com/hepiska/todo-go/utils"



)

func main() {
	port := utils.EnvVar("PORT")

	router:=routers.InitRoute()
	router.Run(port)
}