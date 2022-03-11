package concurrencyingo_test

import (
	"fmt"
	"testing"
	"time"
)

func TestForSelect1(t *testing.T) {
	done := make(chan struct{})
	stringStream := make(chan string)

	go func() {
		for _, s := range []string{"a", "b", "c"} {
			select {
			case <-done:
				return
			case stringStream <- s:
			}
		}
		close(stringStream)
	}()

	for s := range stringStream {
		fmt.Println(s)
	}
}

func TestForSelect2(t *testing.T) {
	done := make(chan struct{})

	time.AfterFunc(1*time.Microsecond, func() {
		close(done)
	})

	for {
		select {
		case <-done:
			return
		default:
		}
		// Do non-preemptable work
		fmt.Println("time: ", time.Now())
	}
}
