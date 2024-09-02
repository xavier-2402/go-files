package hello

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "¡Hola desde Google Cloud Functions en Go!")
}
