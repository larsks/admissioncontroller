package http

import (
	"fmt"
	"net/http"

	"github.com/douglasmakey/admissioncontroller/annotations"
)

// NewServer creates and return a http.Server
func NewServer(port string) *http.Server {
	// Instances hooks
	annotationsValidation := annotations.NewValidationHook()

	// Routers
	ah := newAdmissionHandler()
	mux := http.NewServeMux()
	mux.Handle("/healthz", healthz())
	mux.Handle("/validate/annotations", ah.Serve(annotationsValidation))

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
}
