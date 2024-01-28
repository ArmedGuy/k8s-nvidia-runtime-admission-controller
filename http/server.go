package http

import (
	"fmt"
	"net/http"

	"github.com/ArmedGuy/k8s-nvidia-runtime-admission-controller/pods"
)

// NewServer creates and return a http.Server
func NewServer(port string) *http.Server {
	podsMutation := pods.NewMutationHook()

	// Routers
	ah := newAdmissionHandler()
	mux := http.NewServeMux()
	mux.Handle("/healthz", healthz())
	mux.Handle("/mutate/pods", ah.Serve(podsMutation))

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
}
