package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"kenef.online/dep"
	"kenef.online/handlers"
)

func main() {

	fmt.Println("\nSERVICE START")

	// Load Dependencies
	d, err := dep.Load()
	if err != nil {
		panic(err)
	}
	d.Log.Ok("dependencies loaded")

	// Create wait group
	var wg sync.WaitGroup

	// Create Server
	server := http.Server{
		Addr:         d.Cfg.HostAddr,
		Handler:      handlers.Mux(d, &wg),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Make channels
	ech := make(chan error)
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Start Server
	go func(ch chan error) {
		d.Log.Ok("service listening @ " + d.Cfg.HostAddr)
		ch <- server.ListenAndServe()
	}(ech)

	// Shutdown
	select {
	case err := <-ech:
		d.Log.Error(err)
		os.Exit(1)

	case <-shutdown:
		d.Log.Ok("waiting for handlers to finish")
		wg.Wait()
		d.Log.Ok("SERVICE STOP")
		os.Exit(0)
	}
}
