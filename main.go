package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// read the lines
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		ip := sc.Text()
		netIP := net.ParseIP(ip)
		if !netIP.IsPrivate() && !netIP.IsLoopback() {
			fmt.Println(ip)
		}
	}
}
