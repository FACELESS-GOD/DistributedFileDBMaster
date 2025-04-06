package Utility

import (
	//Structstore "DistributedFileDBMaster/Helper/StructStore"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(Req *http.Request, ReqInterface interface{}) {
	if body, err := ioutil.ReadAll(Req.Body); err == nil {
		if err := json.Unmarshal([]byte(body), ReqInterface); err == nil {
			return
		}
	}
}

// func ParseSessionMessage(Req *Structstore.ChunkMapping, ReqInterface interface{}) {
// 	if body, err := ioutil.ReadAll(Req); err == nil {
// 		if err := json.Unmarshal([]byte(body), ReqInterface); err == nil {
// 			return
// 		}
// 	}
// }
