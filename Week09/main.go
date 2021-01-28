package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const address = ":8080"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eg, ctx1 := errgroup.WithContext(ctx)
	eg.Go(serverTcp(ctx1))
	eg.Go(notifySignal(ctx1))
	err := eg.Wait()
	log.Println(err)
}

func notifySignal(ctx context.Context) func() error {
	return func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGPIPE)
		select {
		case sign := <-quit:
			return errors.New("exit signal " + sign.String())
		case <-ctx.Done():
			signal.Stop(quit)
			fmt.Println("关闭信号监听")
			return nil
		}
	}
}

func serverTcp(ctx context.Context) func() error {
	return func() error {
		listen, err := net.Listen("tcp", address)
		if err != nil {
			panic(err)
		}
		errChan := make(chan error)
		go func() {
			for {
				conn, err := listen.Accept()
				if err != nil {
					errChan <- err
					break
				}
				msgChan := make(chan []byte)
				// todo recover() goroutine
				go readConn(conn, msgChan)
				go writeConn(conn, msgChan)
			}
		}()
		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			_ = listen.Close()
			return nil
		}

	}
}

func readConn(conn net.Conn, msgChan chan<- []byte) {
	reader := bufio.NewReader(conn)
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			log.Println(err)
			break
		}
		msgChan <- data
	}
	close(msgChan)
	conn.Close()

}

func writeConn(conn net.Conn, msgChan <-chan []byte) {
	for data := range msgChan {
		_, err := conn.Write(data)
		if err != nil {
			log.Println(err)
			break
		}
	}
}
