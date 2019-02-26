package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		// These paths will be different on your system.
		seleniumPath     = "vendor/selenium-server-standalone-3.141.59.jar"
		port             = 8080
		chromeDriverPath = "vendor/chromedriver.exe"
	)

	for {
		fmt.Println("what?")
		time.Sleep(5 * time.Second)
	}

}
