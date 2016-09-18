package main

import (
	// "os"
	// "fmt"
	// "github.com/codegangsta/cli"
	// "os/signal"
	// "robotics.neu.edu.tr/ra27-telemetry/ra/bus/inlets/udp"
	// "syscall"
	"encoding/json"
	"io/ioutil"
	"log"
)


func GetConfig(fPath string) Config {
	data, err := ioutil.ReadFile(fPath)
	if err != nil {
		log.Fatal(err)
	}

	var config Config

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
