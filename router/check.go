package router

import "log"

func check(err error) {
	if err != nil {
		log.Fatalf("found an error: %s", err)
	}
}
