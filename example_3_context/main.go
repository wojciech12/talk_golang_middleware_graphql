package main

import (
	"os"

	"github.com/spacelift-io/spcontext"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-kit/log"
)

func wrap(fn func(http.ResponseWriter, *http.Request, http.HandlerFunc)) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fn(w, r, h.ServeHTTP)
		})
	}
}

func main() {
	r := chi.NewRouter()

	ctx := spcontext.New(log.NewJSONLogger(os.Stdout))

	// inject our context to every request
	r.Use(wrap(spcontext.ContextInjector(ctx)))

	// our middleware
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := spcontext.FromStdContext(r.Context())

		/* ctx, span := ctx.StartSpan(
			spcontext.WithOperation("middleware.account"),
			spcontext.WithResource(subdomain),
			spcontext.WithTags("myapp.account.subdomain", subdomain),

			defer func() { span.Close(err) }()
		)*/
		ctx.Errorf("I am here")

		/*
			span.SetTags(
		*/
		w.Write([]byte("welcome"))

		/* next(w, r.WithContext(ctx)) */
	})
	http.ListenAndServe(":3000", r)
}
