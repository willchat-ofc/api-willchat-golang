package setup

import (
	"net/http"
)

func Server() *http.ServeMux {
	mux := http.NewServeMux()

	//db := helpers.MongoHelper()

	//config.SetupRoutes(mux, db)

	return mux
}
