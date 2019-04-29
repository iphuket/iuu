package app

import (
	"github.com/iphuket/iuu/app/http"
)

// Run app
func Run() {
	http.Route()
	http.Run(":80")
}
