package main // import "golang.org/x/lint/golint"

import (
	"fmt"
	"net/http"

	"github.com/ijasmoopan/chkdin-MachineTest/routers"
)

// main is a function that takes no arguments and returns nothing. The program will not work without main().
func main(){

	const (
		addr = ":8080"
	)

	fmt.Println("Server starting on port 8080")

	http.ListenAndServe(addr, routers.Route())
}