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

type ChunkListResponse struct {
	ChunkList map[string]int64
}

type SaveFileRequest struct {
	FileName string
	Size     int64
}

type GetFileRequest struct {
	FileName string
	Size     int64
}

type ChunkIDResponse struct {
	ChunkList []string
}

type ChunkMapping struct {
	ServerID       string
	AvailableSpace int64
	ChunkList      []string
}
