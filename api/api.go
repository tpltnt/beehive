/*
 *    Copyright (C) 2015 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

// beehive's RESTful api for introspection and configuration
package api

import (
	"log"
	"net/http"
	"path"
	_ "strconv"

	"github.com/emicklei/go-restful"
	_ "github.com/emicklei/go-restful/swagger"
)

func configFromPathParam(req *restful.Request, resp *restful.Response) {
	rootdir := "./config"

	subpath := req.PathParameter("subpath")
	if len(subpath) == 0 {
		subpath = "index.html"
	}
	actual := path.Join(rootdir, subpath)
	log.Printf("serving %s ... (from %s)\n", actual, req.PathParameter("subpath"))
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		actual)
}

func imageFromPathParam(req *restful.Request, resp *restful.Response) {
	rootdir := "./assets/bees"

	subpath := req.PathParameter("subpath")
	actual := path.Join(rootdir, subpath)
	log.Printf("serving %s ... (from %s)\n", actual, req.PathParameter("subpath"))
	http.ServeFile(
		resp.ResponseWriter,
		req.Request,
		actual)
}

func Run() {
	// to see what happens in the package, uncomment the following
	//restful.TraceLogger(log.New(os.Stdout, "[restful] ", log.LstdFlags|log.Lshortfile))

	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})

	ws := new(restful.WebService)
	ws.Route(ws.GET("/config/").To(configFromPathParam))
	ws.Route(ws.GET("/config/{subpath:*}").To(configFromPathParam))
	ws.Route(ws.GET("/images/{subpath:*}").To(imageFromPathParam))
	wsContainer.Add(ws)

	b := BeeResource{}
	b.Register(wsContainer)
	h := HiveResource{}
	h.Register(wsContainer)

	log.Println("Starting JSON API on localhost:8181")
	server := &http.Server{Addr: ":8181", Handler: wsContainer}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}
