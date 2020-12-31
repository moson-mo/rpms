package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func startRESTServer(port int, address string) {
	if address == "any" {
		address = ""
	}

	ipe := address + ":" + strconv.Itoa(port)
	http.HandleFunc("/pmtab", pmtab)
	http.HandleFunc("/pmval", pmval)

	fmt.Println("Server waiting for requests @ " + ipe)
	err := http.ListenAndServe(ipe, nil)
	if err != nil {
		fmt.Print("Error starting HTTP server:\n\n")
		fmt.Println(err)
		os.Exit(1)
	}
}

// return pm table
func pmtab(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", acao)
	format := req.URL.Query().Get("format")
	// plain text output
	if format == "plain" {
		mut.RLock()
		values := ""
		for _, val := range pmValues {
			values += fmt.Sprintf("%s: %f\n", val, pmt[val].Value)
		}
		mut.RUnlock()

		w.Write([]byte(values))
	} else { //json output
		mut.RLock()
		js, err := json.MarshalIndent(pmt, "", "\t")
		mut.RUnlock()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

// return single pm value
func pmval(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", acao)
	metric := req.URL.Query().Get("metric")
	if metric == "" {
		http.Error(w, "Need metric parameter", http.StatusInternalServerError)
		return
	}

	mut.RLock()
	val, found := pmt[metric]
	mut.RUnlock()
	if !found {
		http.Error(w, "Metric not found", http.StatusNotFound)
		return
	}

	w.Write([]byte(fmt.Sprint(val.Value)))
}
