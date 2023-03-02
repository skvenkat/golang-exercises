package main

import "sync"

type LockableUint32 struct {
	sync.Mutex
	Data uint32
}

type LockableFloat64 struct {
	sync.Mutex
	Data float64
}

type LockableBool struct {
	sync.Mutex
	Data bool
}

type LockableBytes struct {
	sync.Mutex
	Data []byte
}

func main() {
	var n LockableUint32
	n.Lock()
	n.Data++
	n.Unlock()
	
	var f LockableFloat64
	f.Lock()
	f.Data += 1.23
	f.Unlock()
	
	var b LockableBool
	b.Lock()
	b.Data = !b.Data
	b.Unlock()
	
	var bs LockableBytes
	bs.Lock()
	bs.Data = append(bs.Data, "Go"...)
	bs.Unlock()
}
