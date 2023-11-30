package memory

import (
	"bytes"
	"errors"
	"sync"
)

// TODO: compress data / logger ???

type logger interface {
	Debug(k, v, msg string)
	Info(k, v, msg string)
	Warn(k, v, msg string)
	Err(e error, msg string)
	Fatal(msg string)
}

// Store in-memory object.
type Store struct {
	mu   sync.RWMutex
	data []*bytes.Buffer
	pool sync.Pool
}

// NewStore initializes an in-memory Store.
func NewStore() *Store {
	return &Store{
		pool: sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
		data: make([]*bytes.Buffer, 0),
	}
}

// Fetch retrieves data from the memory Store.
func (s *Store) Fetch(offset int) (*bytes.Buffer, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.data)-1 < offset {
		return nil, errors.New("offset too high")
	}
	return s.data[offset], nil
}

// Push adds a buffer to the memory Store.
func (s *Store) Push(b *bytes.Buffer) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.pool.Put(b)
	newBuffer := s.pool.Get().(*bytes.Buffer)
	s.data = append(s.data, newBuffer)
	return len(s.data) - 1, nil
}

// Read retrieves and removes data from the memory Store.
func (s *Store) Read(offset int) (*bytes.Buffer, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.data) == 0 {
		return nil, errors.New("store is empty")
	}

	if len(s.data)-1 < offset {
		return nil, errors.New("offset too high")
	}

	buffer := s.data[offset]
	clonedBuffer := &bytes.Buffer{}
	clonedBuffer.Write(buffer.Bytes())

	buffer.Reset()
	s.pool.Put(buffer)

	copy(s.data[offset:], s.data[offset+1:])
	s.data = s.data[:len(s.data)-1]

	return clonedBuffer, nil
}
