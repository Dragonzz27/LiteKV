package iomanager

// IOManager is an abstract IO interface that supports different IO types
type IOManager interface {
	// Read reads data from the file at the given position
	Read([]byte, int64) (int, error)

	// Write writes a byte slice to the file
	Write([]byte) (int, error)

	// Sync persists data to stable storage
	Sync() error

	// Close closes the file
	Close() error
}
