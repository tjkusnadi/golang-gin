package main

import (
	"golang-gin/server"

	"go.uber.org/fx"
)

func main() {
	fx.New(server.Module).Run()
}
