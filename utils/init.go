package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func WaitExit(twice bool, chExit chan struct{}, f ...func()) {
	var once bool
	sigChan := make(chan os.Signal, 1)
	signal.Notify(
		sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	for {
		s := <-sigChan
		switch s {
		case syscall.SIGINT, syscall.SIGTERM:
			if twice && !once {
				once = true
				fmt.Println("Send ^C to force exit.")
			} else {
				closeCh(chExit)
				for _, v := range f {
					v()
				}
				os.Exit(0)
			}
		}
	}
}

func closeCh(ch chan struct{}) {
	select {
	case <-ch:
	default:
		close(ch)
	}
}
