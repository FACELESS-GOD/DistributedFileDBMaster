package Route

import (
	"DistributedFileDBMaster/Helper/RouterURL"
	"DistributedFileDBMaster/Package/Controller"

	"github.com/gorilla/mux"
)

func CustomRouter(Router *mux.Router) {

	Router.HandleFunc(RouterURL.GetServerList, Controller.ReturnExistingChunkList).Methods("GET")
	Router.HandleFunc(RouterURL.GetChunkList, Controller.ReturnAvailableChunkList).Methods("GET")

}
