package memory

import (
	"bytes"
	"testing"
)

func BenchmarkFetch(b *testing.B) {
	store := NewStore()
	for i := 0; i < 1000; i++ {
		buffer := bytes.NewBufferString("Test Data")
		_, _ = store.Push(buffer)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := store.Fetch(i % 1000)
		if err != nil {
			b.Fatalf("Fetch failed: %v", err)
		}
	}
}

func BenchmarkPush(b *testing.B) {
	store := NewStore()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buffer := bytes.NewBufferString("Test Data")
		_, _ = store.Push(buffer)
	}
}

func BenchmarkRead(b *testing.B) {
	store := NewStore()
	for i := 0; i < 1000; i++ {
		buffer := bytes.NewBufferString("Test Data")
		_, _ = store.Push(buffer)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = store.Read(i % 1000)
	}
}
