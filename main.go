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
	socks := flag.String("socks", "127.0.0.1:1081", "socks url (default)")
	port := flag.String("port", "8080", "port to listen")
	pure := flag.Bool("pure", false, "pure http/https proxy without socks")
	verbose := flag.Bool("v", false, "verbose")
	version := flag.Bool("version", false, "show version.")

	flag.Parse()

	// Display version info.
	if *version {
		fmt.Println("version=1.1, 2017-9-5, by chinglinwen")
		os.Exit(0)
	}

	dialer, err := proxy.SOCKS5("tcp", *socks, nil, proxy.Direct)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
		os.Exit(1)
	}

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose
	if !*pure {
		proxy.Tr.Dial = dialer.Dial
	}

	log.Fatal(http.ListenAndServe(":"+*port, proxy))
}
