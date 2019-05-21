package app

import (
	"github.com/iphuket/pkt/app/http"
)

// Run app
func Run(port string) {
	http.Run(port)
}
