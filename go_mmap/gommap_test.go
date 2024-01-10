package gommap

import (
	"fmt"
	"testing"
)

func TestLoadMemoryMapFile(t *testing.T) {
	// fileName := "mmap_file"
	// // prepare mmap file
	// want := []byte("HelloWorld!!!\n")
	// os.WriteFile(fileName, want, 0644)

	// got, _ := loadMemoryMapFile(fileName, 0, 100)

	// assert.Equal(t, want, got)
}

func TestPersistHashMap(t *testing.T) {

}

func TestEncodeNumToBytes(t *testing.T) {
	var want uint32 = 65594
	got := decodeBytesToNum(encodeNumToBytes(uint32(want), 4))
	if got != want {
		t.Errorf("Input %d doesn't match encode result %d", want, got)
	}
}

func TestHashMapItemEncode(t *testing.T) {

	keyValue := "Hello There!"
	item := HashMapItem{
		keySize:   28,
		key:       keyValue,
		valueSize: 4,
		value:     123,
	}

	out := item.encode()
	fmt.Print(string(out[:28]))

}
