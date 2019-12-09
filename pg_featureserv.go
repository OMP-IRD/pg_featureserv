package main

/*
 Copyright 2019 Crunchy Data Solutions, Inc.
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
      http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

import (
	"fmt"
	"net/http"

	"github.com/CrunchyData/pg_featureserv/api"
	"github.com/CrunchyData/pg_featureserv/config"
	"github.com/CrunchyData/pg_featureserv/data"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// CatalogInstance mock
var catalogInstance data.Catalog

func init() {
}

func main() {
	log.Printf("%s %s\n", config.AppConfig.Name, config.AppConfig.Version)

	//catalogInstance = data.CatMockInstance()
	catalogInstance = data.CatDBInstance()

	serve()
}

func serve() {

	confServ := config.Configuration.Server
	bindAddress := fmt.Sprintf("%v:%v", confServ.BindHost, confServ.BindPort)
	log.Printf("Serving at: %v\n", bindAddress)

	router := makeRouter()
	log.Fatal(http.ListenAndServe(bindAddress, router))
}

func makeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handleRootJSON)
	router.HandleFunc("/home{.fmt}", handleHome)

	router.HandleFunc("/conformance", handleConformance)
	router.HandleFunc("/conformance.{fmt}", handleConformance)

	router.HandleFunc("/collections", handleCollections)
	router.HandleFunc("/collections.{fmt}", handleCollections)

	router.HandleFunc("/collections/{cid}", handleCollection)
	router.HandleFunc("/collections/{cid}.{fmt}", handleCollection)

	router.HandleFunc("/collections/{cid}/items", handleCollectionItems)
	router.HandleFunc("/collections/{cid}/items.{fmt}", handleCollectionItems)

	router.HandleFunc("/collections/{cid}/items/{fid}", handleItem)
	router.HandleFunc("/collections/{cid}/items/{fid}.{fmt}", handleItem)
	return router
}

func getRequestVar(varname string, r *http.Request) string {
	vars := mux.Vars(r)
	nameFull := vars[varname]
	name := api.PathStripFormat(nameFull)
	return name
}
