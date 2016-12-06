package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/elazarl/goproxy"
	"golang.org/x/net/proxy"
)

func main() {
	socks := flag.String("socks", "127.0.0.1:1081", "socks url")
	port := flag.String("port", "8080", "port to listen")
	quiet := flag.Bool("q", false, "silent the output")
	version := flag.Bool("v", false, "show version.")

	flag.Parse()

	//Display version info
	if *version {
		fmt.Println("version=1.0, 2016-12-5, by chinglinwen")
		os.Exit(0)
	}

	dialer, err := proxy.SOCKS5("tcp", *socks, nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *quiet
	proxy.Tr.Dial = dialer.Dial

	log.Fatal(http.ListenAndServe(":"+*port, proxy))
}
