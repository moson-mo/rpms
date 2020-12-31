package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var acao string
var version = "1.0.0"

func main() {
	// we need root permissions
	if os.Getuid() != 0 {
		fmt.Println("Error: Program must be run as root user")
		os.Exit(1)
	}

	// flags
	flag.Usage = usage
	port := flag.Int("port", 8090, "Port number for the REST API server.")
	address := flag.String("address", "127.0.0.1", "The network address for the REST API server.\nDefine \"any\" to listen on all interfaces.")
	flag.StringVar(&acao, "acao", "null", "Sets the Access-Control-Allow-Origin header if you want to allow querying the API from a webserver.\nThe default value is \"null\" to allow queries from local resources like an html file.")
	interval := flag.Duration("interval", time.Second, "Query interval for reading data from the PM table.")
	flag.Parse()

	// detect pm table layout
	err := setPMTLayout()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// start go routine looping every second loading new pmt values into memory
	go func() {
		for {
			err := parsePMT()
			if err != nil {
				fmt.Println("Error getting PM table:")
				fmt.Println(err)
			}
			time.Sleep(*interval)
		}

	}()

	// start REST endpoint
	startRESTServer(*port, *address)
}

var usage = func() {
	fmt.Printf("rpms - Renoir power metrics server - v%s\n\n", version)
	fmt.Println("Creates a REST-API service exposing power metrics for AMD Renoir processors")

	fmt.Fprintf(os.Stderr, "\nUsage of %s:\n", os.Args[0])

	flag.PrintDefaults()

	fmt.Println("\nAPI endpoints:")
	fmt.Printf("/pmtab -> \t\tReturns full pm table in json format.\n\t\t\tUse URL parameter \"?format=plain\" to get a plain text version.\n\n")
	fmt.Printf("/pmval?metric=xyz -> \tReturns plain text value for a certain metric.\n\t\t\tExample: \"/pmval?metric=SOCKET POWER\"\n\t\t\tUse /pmtab to get a full list of available metrics.\n\n")
}
