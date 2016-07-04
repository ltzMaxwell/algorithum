package heap

import (
	"sync"
)

const (
	capacity = 100
)

//sortable interface to be implemented by different types
type Sortable interface {
	Less(Sortable) bool
}

//Int implementation
type Int int

func (x Int) Less(than Sortable) bool {
	return x < than.(Int)
}

type Heap struct {
	lock sync.Mutex
	data []Sortable
}

func NewHeap() *Heap {
	h := &Heap{}
	h.data = make([]Sortable, 0)
	return h
}

func (h *Heap) IsEmpty() bool {
	h.lock.Lock()
	defer h.lock.Unlock()
	return len(h.data) == 0
}

func (h *Heap) getIndex(el Sortable) int {
	if h.IsEmpty() {
		return -1
	}

	for i := 0; i < len(h.data); i++ {
		if el == h.data[i] {
			return i
		}
	}
	return -1
}

func (h *Heap) siftDown(start, end int) {
	c := start
	//index of left child, as cursor downward
	l := 2*c + 1

	tmp := h.data[c]

	//cursor downward must not over end index
	for l <= end {

		//get bigger one from child nodes(left ,right)
		//l < end to avoid out of range
		if (l < end) && (h.data[l].Less(h.data[l+1])) {
			l++
		}

		if h.data[l].Less(tmp) {
			break
		} else {
			h.data[c] = h.data[l]
			//sift down c
			c = l
			//sift down l
			l = 2*l + 1
		}
	}
	h.data[c] = tmp
}

func (h *Heap) siftUp(start uint64) {
	c := start
	//if not single element
	if c > 0 {
		// p is used as upward cursor
		//make sure it won't out of range
		p := (c - 1) / 2
		tmp := h.data[c]

		for p >= 0 {
			if tmp.Less(h.data[p]) {
				break
			} else {
				//sift down parent value
				h.data[c] = h.data[p]
				//sift up c
				c = p

				//avoid p out of range
				if p == 0 {
					break
				} else {
					//sift up p
					p = (p - 1) / 2
				}
			}
		}
		h.data[c] = tmp
	}
}

func (h *Heap) Insert(el Sortable) {
	h.lock.Lock()
	defer h.lock.Unlock()

	if len(h.data) >= capacity {
		return
	}
	//append to end of slice
	h.data = append(h.data, el)
	h.siftUp(uint64(len(h.data)) - 1)
}

func (h *Heap) Remove(data Sortable) {
	if h.IsEmpty() {
		return
	}
	index := h.getIndex(data)
	h.data[index] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	//remove last element
	h.siftDown(index, len(h.data)-1)
}
