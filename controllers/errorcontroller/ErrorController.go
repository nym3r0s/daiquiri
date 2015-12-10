package errorcontroller

import (
	"net/http"

	"github.com/GokulSrinivas/daiquiri/controllers"
)

func Error404(w http.ResponseWriter, r *http.Request) {
	controllers.WriteJson(w, r, "ERR", "API route not found")
}

func Error401(w http.ResponseWriter, r *http.Request) {
	controllers.WriteJson(w, r, "AUTH", "Authentication Failed")
}
