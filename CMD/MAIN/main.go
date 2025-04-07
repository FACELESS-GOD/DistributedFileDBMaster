package main

import (
	"DistributedFileDBMaster/Helper"
	GRPCHandler "DistributedFileDBMaster/Package/GRPC"
	"DistributedFileDBMaster/Package/Route"
	"DistributedFileDBMaster/Package/Utility"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	Helper.InitializeMetaData()

	//Utility.InitiateSocketConnection()

	go GRPCHandler.InitiateGRPCConnection()

	//go Utility.InitiateBroadCast()

	//go Utility.AcceptingIncommingRequest()

	MuxRouter := mux.NewRouter()

	Route.CustomRouter(MuxRouter)

	http.Handle("/", MuxRouter)

	go Utility.InitiateBroadCast()

	http.ListenAndServe("localhost:9040", MuxRouter)
}
