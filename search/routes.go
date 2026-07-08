package search

import (
	"github.com/rkeplin/bible-go-api/core"
)

var handler = Handler{repo: Repository{}}

func AddRoutes() {
	routes := core.NewRoutes()

	routes.Add(core.Route{
		Name:        "Search",
		Method:      "GET",
		Pattern:     "/search",
		HandlerFunc: handler.Search,
	})

	routes.Add(core.Route{
		Name:        "SearchAggregator",
		Method:      "GET",
		Pattern:     "/searchAggregator",
		HandlerFunc: handler.SearchAggregator,
	})
}
