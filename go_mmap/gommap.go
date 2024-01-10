package gommap

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"syscall"
)

const ITEM_SIZE = 32

type PersistHashMap struct {
	mmapFile string
	mmapSize int
	HashMap  map[string]int32
}

type HashMapItem struct {
	keySize   int
	valueSize int
	key       string
	value     int
}

func loadMemoryMapFile(fileName string, offset, size int) ([]byte, error) {
	file, err := os.Open(fileName)
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

// load from file to address space []byte
// persist HashMap to file

// Decode []byte to HashMap

func (p *PersistHashMap) load() {
	file, err := os.Open(p.mmapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get the size of the file.
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	filesize := stat.Size()

	var readSize int
	if filesize > int64(p.mmapSize) {
		readSize = p.mmapSize
	} else {
		readSize = int(filesize)
	}

	// Memory-map the file.
	addr, err := syscall.Mmap(int(file.Fd()), 0, readSize, syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer syscall.Munmap(addr)

	fmt.Printf("Size of address is %d", len(addr))
	p.HashMap = make(map[string]int32)
	for s := 0; s <= len(addr)-ITEM_SIZE; {
		key, value := decode(addr[s : s+ITEM_SIZE])
		p.HashMap[key] = value
		fmt.Printf("(%s, %d)", key, value)
		s = s + ITEM_SIZE
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

func (h HashMapItem) encode() []byte {
	keyBytes := []byte(h.key)
	paddedBytes := make([]byte, h.keySize)
	if len(keyBytes) > h.keySize {
		copy(paddedBytes[:], keyBytes[:h.keySize])
	} else {
		copy(paddedBytes[:len(keyBytes)], keyBytes[:])
	}

	valueBytes := encodeNumToBytes(uint32(h.value), h.valueSize)
	result := append(paddedBytes, valueBytes...)
	return result
}

func encode(key string, val int32) ([]byte, error) {
	keyBytes := make([]byte, 28)

	keyInput := []byte(key)
	if len(keyInput) > 28 {
		return nil, fmt.Errorf("error: key input exceed 28")
	}

	for i := 0; i < len(keyInput); i++ {
		keyBytes[i] = keyInput[i]
	}
	fmt.Println(len(keyBytes))

	valueBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(valueBytes, uint32(val))
	fmt.Println(len(valueBytes))

	result := append(keyBytes, valueBytes...)
	fmt.Println(len(result))
	return result, nil
}

func decode(input []byte) (string, int32) {
	key := string(input[:27])

	var val int32
	b := input[28:]
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.LittleEndian, &val)
	if err != nil {
		fmt.Println("Binary LittleEndian Failed")
	}

	return key, val
}

func (p *PersistHashMap) persist() {
	// remove old file
	old, err := os.Stat(p.mmapFile)
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

	bw := bufio.NewWriter(file)
	for k, v := range p.HashMap {
		encoded, err := encode(k, int32(v))
		if err != nil {
			log.Fatal(err)
		}
		bw.Write(encoded)
	}

	err = bw.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	persistHashMap := PersistHashMap{
		"persist.log",
		1000,
		make(map[string]int32),
	}
	persistHashMap.HashMap["What"] = 12
	persistHashMap.HashMap["Ever"] = 13
	persistHashMap.persist()
	persistHashMap.load()
	fmt.Print(persistHashMap.HashMap)

	persistHashMap.HashMap["Hello"] = 1
	persistHashMap.HashMap["World"] = 2
	persistHashMap.persist()
	persistHashMap.load()
	fmt.Print(persistHashMap.HashMap)
}
