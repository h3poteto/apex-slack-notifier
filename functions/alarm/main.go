package main

import (
	"github.com/apex/go-apex/sns"
)

func main() {
	sns.HandleFunc(handler)
}
