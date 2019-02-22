package looper

import (
	"time"
)

func present() {
	const (
		// These paths will be different on your system.
		seleniumPath     = "vendor/selenium-server-standalone-3.141.59.jar"
		port             = 8080
		chromeDriverPath = "vendor/chromedriver.exe"
	)

	time.Sleep(5 * time.Second)

}
