package http

import (
	"github.com/iphuket/iuu/app/admin"
	"github.com/iphuket/iuu/app/component/article"
	"github.com/iphuket/iuu/app/component/psutil"
	"github.com/iphuket/iuu/server"
)

var engine = server.New()

// Route of all Settings
func Route() {
	art := engine.Group("/article")
	{
		article.Route(art)
	}
	psu := engine.Group("/psutil")
	{
		psutil.Route(psu)
	}
	adm := engine.Group("/admin")
	{
		admin.Route(adm)
	}

}

// Run ... http
func Run(addr string) {
	engine.Run(addr)
}
