package selects

import (
	"fmt"
	"net/http"
	"time"
)

var TenSecondTimeout = 10 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, TenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

}
func ping(url string) chan bool {
	ch := make(chan bool)

	go func() {
		_, _ = http.Get(url)
		ch <- true
	}()

	return ch
}
