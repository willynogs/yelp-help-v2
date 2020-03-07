package main

import (
	"fmt"

	"github.com/willynogs/yelp-help-v2-server/router"
)

func main() {
	fmt.Printf("-- YELP HELP v2 --\n")
	router.Listen()
}
