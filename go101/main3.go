package main

import "sync"

type Lockable[T any] struct {
    mu sync.Mutex
    data T
}

func (l *Lockable[T]) Do(f func(*T)) {
    l.mu.Lock()
    defer l.mu.Unlock()
    f(&l.data)
}

func main() {
    var n Lockable[uint32]
    n.Do(func(v *uint32) {
        *v++
    })

    var f Lockable[float64]
    f.Do(func(v *float32) {
        *v += 1.23
    })
  
    var b Lockable[bool]
    b.Do(func(v *bool) {
        *v = !*v
    }

    var bs Lockable[[]byte]
    ds.Do(func(v *[]byte) {
          *v = append(*v, "GO"...)
    })
}
