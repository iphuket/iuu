package http

import (
	"github.com/iphuket/iuu/app/component/article"
	"github.com/iphuket/iuu/app/component/psutil"
	"github.com/iphuket/iuu/server"
)

var engine = server.New()

// Route of all Settings
func Route() {
	art := engine.Group("/article")
	psu := engine.Group("/psutil")
	psutil.Route(psu)
	article.Route(art)
}

// Run ... http
func Run(addr string) {
	engine.Run(addr)
}
