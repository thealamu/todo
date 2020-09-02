package db

// MemoryDB is an in-memory implementation of the DB interface
type MemoryDB struct{}

// NewMemDB returns a MemoryDB
func NewMemDB() *MemoryDB {
	return &MemoryDB{}
}
