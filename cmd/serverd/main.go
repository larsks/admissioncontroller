package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/douglasmakey/admissioncontroller/http"

	log "k8s.io/klog/v2"
)

var (
	tlscert, tlskey, port string
)

func main() {
	flag.StringVar(&tlscert, "tlscert", "/etc/certs/tls.crt", "Path to the TLS certificate")
	flag.StringVar(&tlskey, "tlskey", "/etc/certs/tls.key", "Path to the TLS key")
	flag.StringVar(&port, "port", "8443", "The port to listen")
	flag.Parse()

	server := http.NewServer(port)
	quitChan := make(chan struct{})
	go func() {
		if err := server.ListenAndServeTLS(tlscert, tlskey); err != nil {
			log.Errorf("Failed to listen and serve: %v", err)
		}
		close(quitChan)
	}()

	log.Infof("Server running in port: %s", port)

	// listen shutdown signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	exitCode := 0
	for {
		select {
		case <-signalChan:
		case <-quitChan:
		}
		exitCode = 1
		break
	}

	log.Infof("Shutdown gracefully...")
	if err := server.Shutdown(context.Background()); err != nil {
		log.Errorf("Failed to shutdown: %v", err)
		exitCode = 1
	}

	os.Exit(exitCode)
}
