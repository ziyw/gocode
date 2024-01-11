package gommap

import (
	"fmt"
	"testing"
)

// func TestLoadMemoryMapFile(t *testing.T) {
// 	fileName := "mmap_file"
// 	want := []byte("HelloWorld!!!\n")
// 	os.WriteFile(fileName, want, 0644)

// 	// TODO: fix this reading size to dynamic
// 	got, _ := loadMemoryMapFile(fileName, 0, 14)

// 	assert.Equal(t, want, got)
// }

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
	p := PersistHashMap{
		"mmap.file",
		nil,
	}

	p.HashMap = make(map[string]int)
	p.HashMap["hello"] = 1
	p.HashMap["world"] = 2
	p.HashMap["what"] = 3
	p.HashMap["ever"] = 4

	p.persist()
}

func TestHashMapLoad(t *testing.T) {
	p := PersistHashMap{
		"mmap.file",
		nil,
	}
	addr, _ := p.load(300)
	p.initHashMap(addr)
	fmt.Print(p.HashMap)
}
