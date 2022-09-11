package utils

import (
	"log"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	counter := 0
	unsafeWork := func() {
		counter += 1
		if counter < 2 {
			log.Println("fail")
			panic("panic error")
		} else {
			log.Println("ok")
		}
	}

	Retry(unsafeWork, 4, 1*time.Second)
}
