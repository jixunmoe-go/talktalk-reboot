package main

import (
	"flag"
	"fmt"
	"github.com/jixunmoe/talktalk-reboot-go/router"
	"log"
	"os"
	"time"
)

var Version = "0.0.1-unknown"
var BuildTime = "1970-01-01 (unknown)"

func main() {
	var baseURL, username, password string
	flag.StringVar(&baseURL, "url", "http://192.168.1.1", "Base URL to the router.")
	flag.StringVar(&username, "name", "admin", "admin username")
	flag.StringVar(&password, "pass", "S3CR3T", "admin password")
	flag.Parse()

	l := log.New(os.Stderr, "INFO", 0)
	l.Printf("talktalk-reboot (ver %s, build time %s)\n", Version, BuildTime)
	l.Println("")

	client := router.Client{}
	client.Init(baseURL)

	loginResp := client.Login(username, password)
	if loginResp.Error != "ok" {
		log.Fatalf("login failed: %s\n", loginResp.Error)
	}

	rebootResp := client.Reboot()
	if rebootResp.ErrorCode != 0 {
		log.Fatalf("reboot failed with error code: %d", rebootResp.ErrorCode)
	}

	now := time.Now()
	fmt.Printf("reboot ok! (%s)\n", now.String())
}
