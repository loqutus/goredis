package main

import (
	"fmt"
	"os"

	"github.com/loqutus/goredis/pkg/server/keys"
	"github.com/loqutus/goredis/pkg/server/server"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	keys.Dict = make(map[string]string)
	port := ":" + arguments[1]
	server.Serve(port)
}
