package app

import (
	"github.com/iphuket/iuu/app/http"
)

// Run app
func Run(port string) {
	http.Route()
	http.Run(port)
}
