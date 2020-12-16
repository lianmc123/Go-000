package main

import (
	"directory/internal/greeter/di"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_, closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}
	exitSignals := []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGHUP}
	sig := make(chan os.Signal, len(exitSignals))
	signal.Notify(sig, exitSignals...)
	for {
		s := <-sig
		log.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeFunc()
			log.Println("exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
