package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mick-roper/hubl/pkg/common"
	"github.com/mick-roper/hubl/pkg/web"
)

type (
	topic struct {
		Name string
	}

	topicStore interface {
		Get() []topic
	}
)

var (
	sigs chan os.Signal = make(chan os.Signal, 1)
	done chan struct{}  = make(chan struct{}, 1)

	listenAddr string
	topics     topicStore
)

func main() {
	defer close(sigs)

	flag.StringVar(&listenAddr, "listen-addr", ":8080", "server listen address")
	flag.Parse()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	router := buildNewRouter(nil)

	server := http.Server{
		Addr:         listenAddr,
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

	log.Printf("server listening on port: %v", listenAddr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server could not listen on addr %s: %v\n", listenAddr, err)
	}

	<-done
	log.Print("Hubl server stopped")
}

func buildNewRouter(store common.SubscriptionStore) http.Handler {
	if store == nil {
		panic("store is nil!")
	}

	router := http.NewServeMux()

	topicHandler := web.NewTopicHandler(store)
	subscriptionHandler := web.NewSubscriptionHandler(store)

	router.HandleFunc("/topic", topicHandler)
	router.HandleFunc("/subscription", subscriptionHandler)

	return router
}
