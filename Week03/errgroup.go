package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

)

func serverHttp(ctx context.Context, port int) func() error {
	return func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("Success"))
		})
		errChan := make(chan error)
		server := http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: mux,
		}
		go func() {
			if err := server.ListenAndServe(); err != http.ErrServerClosed {
				errChan <- err
			}
		}()
		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			fmt.Println("shutdown", port)
			_ = server.Shutdown(context.Background())
			return nil
		}
	}
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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eg, ctx1 := errgroup.WithContext(ctx)
	eg.Go(serverHttp(ctx1, 8080))
	eg.Go(serverHttp(ctx1, 8082))
	eg.Go(notifySignal(ctx1))
	err := eg.Wait()
	log.Println(err)
}
