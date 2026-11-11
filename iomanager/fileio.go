package iomanager

import "os"

const DataFilePerm = 0644

// FileIO implements standard file IO
type FileIO struct {
	fd *os.File
}

// NewFileIOManager initializes a standard file IO manager
func NewFileIOManager(fileName string) (*FileIO, error) {
	fd, err := os.OpenFile(
		fileName,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		DataFilePerm,
	)
	if err != nil {
		return nil, err
	}
	return &FileIO{fd: fd}, nil
}

// Read reads data from the file at the given offset position
func (fio *FileIO) Read(buffer []byte, offset int64) (int, error) {
	return fio.fd.ReadAt(buffer, offset)
}

// Write writes a byte slice to the end of the file
func (fio *FileIO) Write(buffer []byte) (int, error) {
	return fio.fd.Write(buffer)
}

// Sync persists the file data to stable storage
func (fio *FileIO) Sync() error {
	return fio.fd.Sync()
}

// Close closes the file descriptor
func (fio *FileIO) Close() error {
	return fio.fd.Close()
}
