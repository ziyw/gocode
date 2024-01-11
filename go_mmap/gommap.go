package gommap

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"
)

const ITEM_SIZE = 32
const METADATA_SIZE = 4 // use 2 bytes to save key lenght and value lenght meta data

type PersistHashMap struct {
	mmapFile string
	HashMap  map[string]int
}

func (p PersistHashMap) load(size int) ([]byte, error) {
	file, err := os.Open(p.mmapFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	addr, err := syscall.Mmap(int(file.Fd()), 0, size, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return nil, err
	}
	// To deallocate addr in memory
	// defer syscall.Munmap(addr)
	return addr, nil
}

func (p *PersistHashMap) initHashMap(in []byte) {
	p.HashMap = make(map[string]int)

	hashMeta := decodeBytesToNum(in[:METADATA_SIZE])

	offset := METADATA_SIZE

	for i := uint32(0); i < hashMeta; i++ {
		itemMeta := decodeBytesToNum(in[offset : offset+METADATA_SIZE])
		k, v := decode(in[offset+METADATA_SIZE : offset+METADATA_SIZE+int(itemMeta)])
		fmt.Printf("(%s: %d)", k, v)
		p.HashMap[k] = v
		offset = offset + METADATA_SIZE + int(itemMeta)
	}
}

func (p PersistHashMap) persist() {
	// remove old file
	old, _ := os.Stat(p.mmapFile)
	if old != nil {
		e := os.Remove(p.mmapFile)
		if e != nil {
			log.Fatal(e)
		}
	}

	// flush hashmap to file
	file, err := os.Create(p.mmapFile)
	if err != nil {
		log.Fatal()
	}
	defer file.Close()

	// create buffered writer to persist hashmap content
	bw := bufio.NewWriter(file)

	// write metadata first: number of items in hashmap
	fmeta := encodeNumToBytes(uint32(len(p.HashMap)), METADATA_SIZE)
	bw.Write(fmeta)

	for k, v := range p.HashMap {
		content := encode(k, v)
		meta := encodeNumToBytes(uint32(len(content)), METADATA_SIZE)
		item := append(meta, content...)
		bw.Write(item)
	}

	err = bw.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func encodeNumToBytes(num uint32, size int) []byte {
	output := make([]byte, size)
	binary.LittleEndian.PutUint32(output, uint32(num))
	return output
}

func decodeBytesToNum(input []byte) uint32 {
	var num uint32
	buf := bytes.NewReader(input)
	err := binary.Read(buf, binary.LittleEndian, &num)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func intToString(num int) string {
	return fmt.Sprintf("%d", num)
}

func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

func encode(key string, val int) []byte {
	kBytes := []byte(key)
	kmeta := encodeNumToBytes(uint32(len(kBytes)), METADATA_SIZE)
	encoded := append(kmeta, kBytes...)

	// metadata is the lenght of the value
	vBytes := []byte(intToString(val))
	vmeta := encodeNumToBytes(uint32(len(vBytes)), METADATA_SIZE)
	encoded = append(encoded, vmeta...)
	encoded = append(encoded, vBytes...)
	return encoded
}

func decode(in []byte) (string, int) {
	kmeta := decodeBytesToNum(in[:METADATA_SIZE])
	key := string(in[METADATA_SIZE : METADATA_SIZE+kmeta])

	voffset := METADATA_SIZE + kmeta
	vmeta := decodeBytesToNum(in[voffset : voffset+METADATA_SIZE])
	value := stringToInt(string(in[voffset+METADATA_SIZE : voffset+METADATA_SIZE+vmeta]))

	return key, value
}
