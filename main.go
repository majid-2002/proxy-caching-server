package main

import (
	"flag"
	"log"
	"caching-proxy/proxy"
	"os"
)

func main() {
	port := flag.Int("port", 8080, "Port for the proxy server")
	origin := flag.String("origin", "", "Origin server URL")
	clearCache := flag.Bool("clear-cache", false, "Clear the cache")
	flag.Parse()

	
	if *clearCache {
		proxy.ClearCache()
		log.Println("Cache cleared successfully")
		os.Exit(0)
	}

	if *origin == "" {
		log.Fatal("Error: origin server URL must be specified")
	}

	proxy.StartServer(*port, *origin)
}
