package gommap

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"log"
	"os"
	"syscall"
)

const METADATA_SIZE = 4

func loadMmap(filename string, size int) ([]byte, error) {
	file, err := os.Open(filename)
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

func persist(filename string, hashmap map[string]int) {
	// remove old file
	old, _ := os.Stat(filename)
	if old != nil {
		e := os.Remove(filename)
		if e != nil {
			log.Fatal(e)
		}
	}

	// flush hashmap to file
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal()
	}
	defer file.Close()

	// create buffered writer to persist hashmap content
	bw := bufio.NewWriter(file)
	byteArr := encodeHashMap(hashmap)
	bw.Write(byteArr)

	if err = bw.Flush(); err != nil {
		log.Fatal(err)
	}
}

// encode and decode string-int hashmap to byte array
func encodeHashMap(hashMap map[string]int) []byte {
	byteArr := []byte{}
	byteArr = append(byteArr, encodeNumToBytes(uint32(len(hashMap)), METADATA_SIZE)...)

	for key, value := range hashMap {
		b := encodeStrToByteArr(key)
		b = append(b, encodeIntToByteArr(value)...)
		byteArr = append(byteArr, b...)
	}
	return byteArr
}

func decodeByteArrToHashMap(byteArr []byte) map[string]int {
	hashMap := make(map[string]int)

	n := int(decodeBytesToNum(byteArr[:METADATA_SIZE]))
	offset := METADATA_SIZE
	for i := 0; i < n; i++ {
		key, nxt := decodeByteArrToString(byteArr, offset)
		value, nxt := decodeByteArrToInt(byteArr, nxt)
		hashMap[key] = value
		offset = nxt
	}
	return hashMap
}

// Encode and decode  string to byte array
func encodeStrToByteArr(in string) []byte {
	// String Metadat: 4 bytes content length, variable content size
	content := []byte(in)
	meta := encodeNumToBytes(uint32(len(in)), METADATA_SIZE)
	return append(meta, content...)
}

func decodeByteArrToString(byteArr []byte, offset int) (string, int) {
	meta := decodeBytesToNum(byteArr[offset : offset+METADATA_SIZE])
	contentOffset := offset + METADATA_SIZE
	contentStop := contentOffset + int(meta)
	str := string(byteArr[offset+METADATA_SIZE : contentStop])
	return str, contentStop
}

// Encode and decode int to byte array
func encodeIntToByteArr(in int) []byte {
	// metadata: 4 bytes content lenght, 1 byte sign flag, 4 byte unsign int value
	// TODO: change this to use one bit
	var flag byte
	var content []byte
	if in > 0 {
		flag = 0
		content = encodeNumToBytes(uint32(in), METADATA_SIZE)
	} else {
		flag = 255
		content = encodeNumToBytes(uint32(-in), METADATA_SIZE)
	}
	meta := encodeNumToBytes(uint32(len(content)), METADATA_SIZE)
	meta = append(meta, flag)
	return append(meta, content...)
}

func decodeByteArrToInt(byteArr []byte, offset int) (int, int) {
	meta := decodeBytesToNum(byteArr[offset : offset+METADATA_SIZE])
	flag := byteArr[offset+METADATA_SIZE : offset+METADATA_SIZE+1][0]

	contentOffset := offset + METADATA_SIZE + 1
	contentStop := contentOffset + int(meta)

	content := int(decodeBytesToNum(byteArr[contentOffset:contentStop]))
	if flag == 255 {
		return -1 * content, contentStop
	}
	return content, contentStop
}

// Fundamental: encode uint32 to bytes and decode bytes to uint32
// These two are used in all meta data handling
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
