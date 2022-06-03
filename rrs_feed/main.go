package main

import (
	"log"
	"os"
)

//init is called prior to main
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
