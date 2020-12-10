package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	errGroup, ctx := errgroup.WithContext(context.Background())
	errGroup.Go(func() (err error) {
		errChan := make(chan error)
		server := http.Server{Addr: "127.0.0.1:8080"}
		go func() {
			fmt.Println("server start")
			errChan <- server.ListenAndServe()
		}()

		select {
		case e := <-errChan:
			err = fmt.Errorf("server err: %v\n", e.Error())
		case <-ctx.Done():
			_ = server.Close()
			fmt.Println("server stopped")
		}

		return
	})

	errGroup.Go(func() (err error) {
		signalChan := make(chan os.Signal)
		signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

		select {
		case s := <-signalChan:
			err = fmt.Errorf("handle signal: %v\n", s)
		case <-ctx.Done():
			signal.Stop(signalChan)
			fmt.Println("signal notify stopped")
		}

		return
	})

	if err := errGroup.Wait(); err != nil {
		fmt.Println(err)
	}
}
