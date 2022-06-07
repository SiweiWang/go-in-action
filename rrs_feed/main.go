package main

import (
	"log"
	"os"

	"github.com/SiweiWang/go-in-action/rrs_feed/search"
)

//init is called prior to main
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
