package containers

type Block struct {
	// Hash of the current block
	Hash         []byte
	// Data Stored inside the Block
	Data         []byte
	// Hash of the previous Block
	PreviousHash []byte
}