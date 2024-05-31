package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "postgres://user:pass@host:2137/path?k=v#f"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host, port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
}
