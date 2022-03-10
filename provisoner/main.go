package main

import (
	"flag"
	"log"
	"os/exec"
)

func main() {
	nsPtr := flag.String("ns", "default", "namespace")
	portPtr := flag.String("port", "0", "port")
	flag.Parse()

	if *nsPtr == "default" || *portPtr == "0" {
		panic("namespace/port not provided")
	}

	log.Printf("Deploying resources to [%s] namespace", *nsPtr)

	res := exec.Command("/bin/sh", "./apply.sh", *nsPtr, *portPtr)

	if err := res.Start(); err != nil {
		panic(err)
	}
}
