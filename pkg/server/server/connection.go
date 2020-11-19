package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/loqutus/goredis/pkg/server/keys"
)

func handleConnection(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		tempString := strings.TrimSpace(string(netData))
		cmdSplit := strings.Split(tempString, " ")
		switch cmdSplit[0] {
		case "stop":
			break
		case "get":
			if len(cmdSplit) < 2 {
				c.Write([]byte("missing key to get"))
				break
			} else {
				value, err := keys.Get(cmdSplit[1])
				if err != nil {
					c.Write([]byte(value))
				} else {
					c.Write([]byte(err.Error()))
				}
			}
		case "set":
			if len(cmdSplit) < 3 {
				c.Write([]byte("not enough arguments for set"))
				break
			} else {
				err := keys.Set(cmdSplit[1], cmdSplit[2])
				if err != nil {
					c.Write([]byte(err.Error()))
				} else {
					c.Write([]byte("ok"))
				}
			}
		}
	}
}
