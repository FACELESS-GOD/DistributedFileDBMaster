package Structstore

type GenericReqFile struct {
	Size int64
}

type GenericReq struct {
	FileName string
}

type GenericResponseData struct {
	ChunkServer []string
	Message     string
}
