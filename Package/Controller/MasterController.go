package Controller

import (
	"DistributedFileDBMaster/Helper"
	Structstore "DistributedFileDBMaster/Helper/StructStore"
	"DistributedFileDBMaster/Package/Utility"
	"encoding/json"
	"net/http"
)

func ReturnAvailableChunkList(writer http.ResponseWriter, Req *http.Request) {
	RequestInstance := &Structstore.SaveFileRequest{}
	ResponseInstance := Structstore.ChunkListResponse{}

	Utility.ParseBody(Req, RequestInstance)
	var chunkList map[string]int64

	var fileSize int64 = RequestInstance.Size
	var serverList []string

	for key, value := range Helper.AvailableChunk {
		if fileSize > 0 {
			if value > 0 {

				if fileSize > int64(value) {
					fileSize = fileSize - int64(value)
					serverList = append(serverList, key)
					chunkList[key] = value
				} else {
					serverList = append(serverList, key)
					chunkList[key] = value
					break
				}
			}
		} else {
			break
		}

	}

	Helper.ChunkMapping[RequestInstance.FileName] = serverList

	ResponseInstance.ChunkList = chunkList
	res, _ := json.Marshal(ResponseInstance)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)

}

func ReturnChunkList(writer http.ResponseWriter, Req *http.Request) {

	RequestInstance := &Structstore.GetFileRequest{}
	ResponseInstance := Structstore.ChunkIDResponse{}

	Utility.ParseBody(Req, RequestInstance)

	fileName := RequestInstance.FileName

	ChunkIDList, ok := Helper.ChunkMapping[fileName]

	if ok == true {
		ResponseInstance.ChunkList = ChunkIDList
		res, _ := json.Marshal(ResponseInstance)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write(res)
		return
	} else {
		genericResponse := Structstore.GenericResponseData{}
		genericResponse.Message = "File Not Found"
		res, _ := json.Marshal(genericResponse)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		writer.Write(res)
	}

	res, _ := json.Marshal(ResponseInstance)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)

}
