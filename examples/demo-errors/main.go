package main

import (
	"net/http"

	"github.com/alexisvisco/kcd/pkg/errors"
	"github.com/go-chi/chi"

	"github.com/alexisvisco/kcd"
)

func main() {
	r := chi.NewRouter()
	r.Post("/", kcd.Handler(YourHttpHandler, http.StatusOK))
	_ = http.ListenAndServe(":3000", r)
}

func YourHttpHandler() error {
	return errors.NewWithKind(errors.KindInternal, "this is an error")
}

// Test it : curl -XPOST 'localhost:3000'
