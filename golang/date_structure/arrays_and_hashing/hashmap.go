package datastruc

import "fmt"

type Pair struct {
	Key string
	Val string
}

func NewPair(key string, val string) *Pair {
	return &Pair{
		Key: key,
		Val: val,
	}
}

type HashMap struct {
	Size     int
	Capacity int
	Map      []*Pair
}

func NewHashMap() *HashMap {
	return &HashMap{
		Size:     0,
		Capacity: 2,
		Map:      make([]*Pair, 2),
	}
}

func (h *HashMap) Hash(key string) int {
	index := 0
	for i := 0; i < len(key); i++ {
		index += int(key[i])
	}
	return index % h.Capacity
}

func (h *HashMap) Get(key string) string {
	index := h.Hash(key)
	for h.Map[index] != nil {
		if h.Map[index].Key == key {
			return h.Map[index].Val
		}
		index++
		index = index % h.Capacity
	}
	return ""
}

func (h *HashMap) Put(key string, val string) {
	index := h.Hash(key)

	for {
		if h.Map[index] == nil {
			h.Map[index] = NewPair(key, val)
			h.Size++
			if h.Size >= h.Capacity/2 {
				h.Rehash()
			}
			return
		} else if h.Map[index].Key == key {
			h.Map[index].Val = val
			return
		}
		index++
		index = index % h.Capacity
	}
}

func (h *HashMap) Remove(key string) {
	if h.Get(key) == "" {
		return
	}

	index := h.Hash(key)
	for {
		if h.Map[index].Key == key {
			h.Map[index] = nil
			h.Size--
			return
		}
		index++
		index = index % h.Capacity
	}
}

func (h *HashMap) Rehash() {
	h.Capacity = 2 * h.Capacity
	newMap := make([]*Pair, h.Capacity)

	oldMap := h.Map
	h.Map = newMap
	h.Size = 0
	for _, p := range oldMap {
		if p != nil {
			h.Put(p.Key, p.Val)
		}
	}
}

func (h *HashMap) Print() {
	for _, p := range h.Map {
		if p != nil {
			fmt.Println(p.Key, p.Val)
		}
	}
}
