package handlers

import (
	"fmt"
	"net/http"

	"../visits"
	"../ip"
	"../config"
)

var visitors = make(map[string]*visits.Visitors)

// IndexPage handle "/" path
func IndexPage(w http.ResponseWriter, r *http.Request) {
	visitorIP := ip.FromRequest(r)

	url := "localhost" + config.Port + r.URL.String()

	if visitorIP == "::1" {
		visitorIP = "localhost"
	}

	if _, ok := visitors[url]; ok {
		if _, ok := visitors[url].GetByIP(visitorIP); ok {
			visitors[url].IterateVisitByIP(visitorIP)
		} else {
			visitors[url].AddNewVisitor(visitorIP)
		}
	} else {
		visitors[url] = &visits.Visitors{}
		visitors[url].AddNewVisitor(visitorIP)
	}
	 

	fmt.Fprintf(w, "<h1 style='color: green'>TOTAL VISITS THIS PAGE: %d<h1></br>", visitors[url].GetTotalVisits())
	fmt.Fprintf(w, "<h1 style='color: blue'>UNIQUE VISITORS: %d<h1></br>", visitors[url].GetUniqueVisitors())

	for key, vis := range visitors {
		fmt.Fprintf(w, "<h2 s>%s</h2></br>", key)
		for i, visitor := range vis.All {
			fmt.Fprintf(w, "<span>%d) %s ... %5d</span></br>", i, visitor.IP, visitor.Counter)
		}
	}
}