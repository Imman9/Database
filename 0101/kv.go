package db0101

import "bytes"

type KV struct {
	mem map[string][]byte
}

func (kv *KV) Open() error {
	kv.mem = map[string][]byte{} // empty
	return nil
}

func (kv *KV) Close() error { return nil }

func (kv *KV) Get(key []byte) (val []byte, ok bool, err error){
	val,ok = kv.mem[string(key)]
	return
}

func (kv *KV) Set(key []byte, val []byte) (updated bool, err error){

	//1. get the prev value and check whether the key exists
	prev,exists := kv.mem[string(key)]

	//2.insert the new value or replace the old one
	kv.mem[string(key)] = val

	//3.check if updated
	updated = !exists || !bytes.Equal(val,prev)

	//4.return updated and err
	return
}

func (kv *KV) Del(key []byte) (deleted bool, err error){
	
	//check if key exists
	_, deleted = kv.mem[string(key)]

	//delete the value
	delete(kv.mem, string(key))

	//naked return
	return
}

// QzBQWVJJOUhU https://trialofcode.org/
