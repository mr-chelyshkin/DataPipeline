package memory

import (
	"bytes"
	"testing"
)

func TestStore(t *testing.T) {
	store := NewStore()

	buffer := bytes.NewBufferString("Test Data")
	index, err := store.Push(buffer)
	if err != nil {
		t.Fatalf("PushBuffer failed: %v", err)
	}

	fetchedBuffer, err := store.Fetch(index)
	if err != nil {
		t.Fatalf("Fetch failed: %v", err)
	}
	if fetchedBuffer.String() != "Test Data" {
		t.Fatalf("Expected fetched buffer to be 'Test Data', got '%s'", fetchedBuffer.String())
	}

	removedBuffer, err := store.Read(index)
	if err != nil {
		t.Fatalf("ReadAndRemove failed: %v", err)
	}
	if removedBuffer.String() != "Test Data" {
		t.Fatalf("Expected removed buffer to be 'Test Data', got '%s'", removedBuffer.String())
	}

	_, err = store.Fetch(index)
	if err == nil {
		t.Fatalf("Fetch should have returned an error after removal")
	}
}

func TestEmptyStore(t *testing.T) {
	store := NewStore()

	_, err := store.Fetch(0)
	if err == nil {
		t.Fatalf("Fetch on an empty store should have returned an error")
	}

	_, err = store.Read(0)
	if err == nil {
		t.Fatalf("ReadAndRemove on an empty store should have returned an error")
	}
}
