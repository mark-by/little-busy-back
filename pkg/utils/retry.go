package utils

import "time"

func Retry(work func(), count int, delay time.Duration) {

	safeWork := func(ok *bool) {
		defer func (ok *bool) {
			if err := recover(); err != nil {
				*ok = false
			}
		} (ok)
		work()
		*ok = true
	}

	ok := false

	for i := 0; i < count; i++ {
		safeWork(&ok)
		if ok {
			break
		}
		time.Sleep(delay)
	}
}
