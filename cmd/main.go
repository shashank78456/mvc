package main

import (
	"fmt"
	"github.com/shashank78456/mvc/pkg/api"
)

func main() {
	fmt.Println("Started Server at localhost on port 3000")
	api.Start()
}