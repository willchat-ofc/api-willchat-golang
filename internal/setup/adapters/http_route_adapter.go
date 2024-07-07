package adapters

import (
	"io"
	"net/http"

	"github.com/willchat-ofc/api-willchat-golang/internal/presentation/protocols"
)

func AdaptRoute(controller protocols.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpRequest := &protocols.HttpRequest{
			Body:      r.Body,
			Header:    r.Header,
			UrlParams: r.URL.Query(),
		}

		res := controller.Handle(*httpRequest)

		w.WriteHeader(res.StatusCode)
		_, err := io.Copy(w, res.Body)
		if err != nil {
			http.Error(w, "Failed to write response body", http.StatusInternalServerError)
			return
		}
	})

}
