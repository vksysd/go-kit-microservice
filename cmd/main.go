package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"nepodate"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	var(
		httpAddr = flag.String("http",":8080","http listen address")
	)
	flag.Parse()

	ctx := context.Background()

	//nepodate service
	srv := nepodate.NewService()

	errChan := make (chan error)

	go func() {
		c := make(chan os.Signal,1)
		signal.Notify(c, syscall.SIGINT,syscall.SIGTERM)
		errChan <- fmt.Errorf("%s",<-c)
	}()

	//mapping endpoints
	endpoints := nepodate.Endpoints{
		GetEndpoint:      nepodate.MakeGetEndpoint(srv),
		StatusEndpoint:   nepodate.MakeStatusEndpoint(srv),
		ValidateEndpoint: nepodate.MakeValidateEndpoint(srv),
	}

	go func() {
		log.Printf("nepodate is listening on port: %s", *httpAddr)
		handler := nepodate.NewHTTPServer(ctx,endpoints)
		errChan <- http.ListenAndServe(*httpAddr,handler)
	}()

	log.Fatalln(<- errChan)
}