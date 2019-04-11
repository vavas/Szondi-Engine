package main

import (
	"log"
	"net"
	"net/rpc"
	"os"

	db "github.com/excelWithBusiness/MSGO-Util/database"
)

var (
	//gStatus global channel for getting server setup status
	gStatus chan string
)

//setSetupChannel set setup channel to use for server status
func setSetupChannel(status chan string) chan string {
	gStatus = status
	return gStatus
}

//getSetupChannel get setup channel for server status
func getSetupChannel() chan string {
	return gStatus
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Project %s: build %s\n", os.Getenv("PROJECT_NAME"), os.Getenv("PROJECT_BUILD"))

	log.Println("Loading services...")
	services := new(Services)
	services.db = db.Instance().Connect()
	services.helper = new(Helper)

	log.Println("Attaching services...")
	server := rpc.NewServer()
	server.RegisterName("AK", services)

	log.Println("Starting rpc server...")
	ch := getSetupChannel()
	if ch != nil {
		ch <- "done"
	}

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	server.Accept(l)
}
