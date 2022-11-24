package main

import (
	_ "github.com/study-onboard/echo/src/endpoints"
	"github.com/study-onboard/echo/src/server"
)

func main() {
	server.Server.Logger.Fatal(server.Server.Start(":9090"))
}
