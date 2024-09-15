package main

import (
	"fmt"
	"net"
	"redis/resp"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		respConn := resp.NewResp(conn, conn)

		value, err := respConn.Read()
		if err != nil {
			fmt.Println(err)
			return
		}

		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		handler, ok := Handlers[command]
		if !ok {
			err = respConn.Write(resp.Value{Typ: resp.STRING, Str: "Unknown command"})
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		response := handler(args)
		err = respConn.Write(response)
		if err != nil {
			fmt.Println(err)
		}
	}
}
