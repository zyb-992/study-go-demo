package sort

type Heap[T int64 | int | int32 | float32 | float64 | string] []T

func (h *Heap[T]) Push(v T) {
	// 写入堆尾
	*h = append(*h, v)
	// 排序？堆化
	h.siftUp(len(*h) - 1)
}

func (h *Heap[T]) Parent(i int) int {
	// 获取父节点
	if i != 0 {
		return (i - 1) / 2
	}

	return 0
}

func (h *Heap[T]) Pop() {

}

// 自底向上
func (h *Heap[T]) siftUp(index int) {
	for {
		parentIdx := h.Parent(index)
		if index == 0 || (*h)[parentIdx] < (*h)[index] {
			return
		}

		h.Swap(parentIdx, index)
		h.siftUp(parentIdx)
	}
}

func (h *Heap[T]) siftDown() {

}

func (h *Heap[T]) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
