package main

import (
	pkg "github.com/codescalersinternships/NTP-client-rawan/pkg"
)

func main() {
	resp,err:= pkg.Client()
	if err != nil {
		panic(err)
	}
	println(resp)
}