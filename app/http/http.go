package http

import (
	"github.com/iphuket/pkt/app/account"
	"github.com/iphuket/pkt/app/admin"
	"github.com/iphuket/pkt/app/component/article"
	"github.com/iphuket/pkt/app/component/psutil"
	"github.com/iphuket/pkt/app/component/shoturl"

	"github.com/iphuket/pkt/server"
)

var engine = server.Engine

// Route of all Settings
func Route() {
	art := engine.Group("article")
	{
		article.Route(art)
	}
	psu := engine.Group("psutil")
	{
		psutil.Route(psu)
	}
	adm := engine.Group("admin")
	{
		admin.Route(adm)
	}
	acc := engine.Group("account")
	{
		account.Route(acc)
	}
	stu := engine.Group("/")
	{
		shoturl.Route(stu)
	}

}

// Run ... http
func Run(addr string) {
	Route()
	engine.Run(addr)
}
