package main

import (
	"flag"

	pkg "github.com/codescalersinternships/NTP-client-rawan/pkg"
)

func main() {
	var server string
	flag.StringVar(&server, "server", "pool.ntp.org:123", "Specifies custom NTP server")
	flag.Parse()

	resp, err := pkg.Client(server)
	if err != nil {
		panic(err)
	}

	time, err := pkg.ParseNTPResponse(resp)
	if err != nil {
		panic(err)
	}
	println("Current time:", time)
}
