package datastruc

type Heap struct {
	heap []int
}

func NewHeap() *Heap {
	return &Heap{
		heap: []int{0},
	}
}

func (h *Heap) Push(val int) {
	h.heap = append(h.heap, val)
	i := len(h.heap) - 1

	// Percolate up
	for i > 1 && h.heap[i] < h.heap[i/2] {
		// Swap values
		tmp := h.heap[i]
		h.heap[i] = h.heap[i/2]
		h.heap[i/2] = tmp
		i = i / 2
	}
}

func (h *Heap) Pop() int {
	if len(h.heap) == 1 {
		return -1 // In Go, we can't return nil for int type. So, we return -1 to denote error.
	}
	if len(h.heap) == 2 {
		popVal := h.heap[len(h.heap)-1]
		h.heap = h.heap[:len(h.heap)-1]
		return popVal
	}

	res := h.heap[1]
	// Move last value to root
	h.heap[1] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	i := 1
	// Percolate down
	for 2*i < len(h.heap) {
		if 2*i+1 < len(h.heap) && h.heap[2*i+1] < h.heap[2*i] && h.heap[i] > h.heap[2*i+1] {
			// Swap right child
			tmp := h.heap[i]
			h.heap[i] = h.heap[2*i+1]
			h.heap[2*i+1] = tmp
			i = 2*i + 1
		} else if h.heap[i] > h.heap[2*i] {
			// Swap left child
			tmp := h.heap[i]
			h.heap[i] = h.heap[2*i]
			h.heap[2*i] = tmp
			i = 2 * i
		} else {
			break
		}
	}
	return res
}

func (h *Heap) Heapify(arr []int) {
	// 0-th position is moved to the end
	arr = append(arr, arr[0])
	h.heap = arr
	cur := (len(h.heap) - 1) / 2

	for cur > 0 {
		// Percolate Down
		i := cur
		for 2*i < len(h.heap) {
			if 2*i+1 < len(h.heap) && h.heap[2*i+1] < h.heap[2*i] && h.heap[i] > h.heap[2*i+1] {
				// Swap right child
				h.heap[i], h.heap[2*i+1] = h.heap[2*i+1], h.heap[i]
				i = 2*i + 1
			} else if h.heap[i] > h.heap[2*i] {
				// Swap left child
				h.heap[i], h.heap[2*i] = h.heap[2*i], h.heap[i]
				i = 2 * i
			} else {
				break
			}
		}
		cur--
	}
}
