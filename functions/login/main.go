package main

import (
	"github.com/apex/go-apex/cloudwatch"
)

func main() {
	cloudwatch.HandleFunc(handler)
}
