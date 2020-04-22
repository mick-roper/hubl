package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/mick-roper/hubl/pkg/common"
	"github.com/mick-roper/hubl/pkg/data"
	"github.com/mick-roper/hubl/pkg/web"
)

var (
	sigs chan os.Signal = make(chan os.Signal, 1)
	done chan struct{}  = make(chan struct{}, 1)

	listenPort int
)

func main() {
	defer close(sigs)

	flag.IntVar(&listenPort, "port", 8080, "server listen port")
	flag.Parse()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	store, err := data.NewMemorySubscriptionStore()
	if err != nil {
		panic(err)
	}
	defer store.Close()

	server := buildServer(store)

	log.Printf("server listening on port: %v", listenPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server could not listen on addr %v: %v\n", listenPort, err)
	}

	<-done
	log.Print("Hubl server stopped")
}

func buildServer(store common.SubscriptionStore) *http.Server {
	if store == nil {
		panic("store is nil")
	}

	router := http.NewServeMux()

	topicHandler := web.NewTopicHandler(store)
	subscriptionHandler := web.NewSubscriptionHandler(store)

	router.HandleFunc("/topic", topicHandler)
	router.HandleFunc("/subscription", subscriptionHandler)

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(listenPort),
		Handler:      router,
		ErrorLog:     log.New(os.Stdout, "webserver:", log.LstdFlags),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}
	server.SetKeepAlivesEnabled(true)

	go func() {
		<-sigs
		log.Print("server shutting down...")
		server.SetKeepAlivesEnabled(false)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("an error occured shutting down the server: %v\n", err)
		}
		close(done)
	}()

	return server
}
