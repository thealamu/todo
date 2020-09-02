package db

// MemoryDB is an in-memory implementation of the DB interface
type MemoryDB struct{}

// NewInMem returns an in-memory DB
func NewInMem() *MemoryDB {
	return &MemoryDB{}
}
