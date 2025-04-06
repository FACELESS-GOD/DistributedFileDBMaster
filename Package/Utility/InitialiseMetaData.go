package Utility

var ChunkMapping map[string][]string

var AvailableChunk map[string]int64

func InitializeMetaData() {
	AvailableChunk = make(map[string]int64)
	ChunkMapping = make(map[string][]string)

}
