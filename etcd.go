package main

import (
	"flag"
	"fmt"

	"gotest/discovery"
)

func main() {
	var role = flag.String("role", "", "master | worker")
	flag.Parse()
	endpoints := []string{"https://10.255.209.167:2379", "https://10.255.209.167:22379", "https://10.255.209.167:32379"}
	if *role == "master" {
		master := discovery.NewMaster(endpoints)
		master.WatchWorkers()
	} else if *role == "worker" {
		worker := discovery.NewWorker("localhost", "127.0.0.1", endpoints)
		worker.HeartBeat()
	} else {
		fmt.Println("example -h for usage")
	}
}
