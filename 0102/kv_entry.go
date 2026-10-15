package db0102

import (
	"encoding/binary"
	"io"
)

type Entry struct {
	key []byte
	val []byte
}

func (ent *Entry) Encode() []byte{
	//first create a slice that holds the bits
	data := make([]byte, 4 + 4 + len(ent.key) + len(ent.val))

	//put the lenghts of the key and val into the slice
    binary.LittleEndian.PutUint32(data[0:4],uint32(len(ent.key)))  
    binary.LittleEndian.PutUint32(data[4:8],uint32(len(ent.val)))  

	//cpy the real values of key and val
	copy(data[8:],ent.key)
	copy(data[8+len(ent.key):],ent.val)

	return data
}


func (ent *Entry) Decode(r io.Reader) error {

	// create an array that would hold the key and val length
	var header [8]byte

	//try writing into hearder
	if _,err := io.ReadFull(r , header[:]); err != nil{
		return err
	}

	//get the key and value lengths
	klen := int(binary.LittleEndian.Uint32(header[0:4])) 
	vlen := int(binary.LittleEndian.Uint32(header[4:8]))

	//create a slice that holds the actual key and val
	 data := make([]byte, klen + vlen)

	 //write into data
	 if _, err := io.ReadFull(r , data); err != nil{
		return err
	 }

	 //save the values of key and val in Entry struct

	 ent.key = data[:klen]
	 ent.val = data[klen:]

	 return nil
}

// QzBQWVJJOUhU https://trialofcode.org/
