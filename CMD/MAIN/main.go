package main

import (
	"DistributedFileDBMaster/Package/Route"
	"DistributedFileDBMaster/Package/Utility"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	Utility.InitializeMetaData()

	Utility.InitiateSocketConnection()

	Utility.InitiateBroadCast()

	MuxRouter := mux.NewRouter()

	Route.CustomRouter(MuxRouter)

	http.Handle("/", MuxRouter)

	http.ListenAndServe("localhost:9040", MuxRouter)
}
