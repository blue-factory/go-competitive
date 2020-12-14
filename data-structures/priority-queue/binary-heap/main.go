package main

// https://www.geeksforgeeks.org/binary-heap/?ref=rp
// https://www.geeksforgeeks.org/max-heap-in-python/?ref=rp
// https://www.geeksforgeeks.org/max-heap-in-java/?ref=rp

const (
	max = 100
)

// BinaryHeap ...
type BinaryHeap interface {
	Insert(int)
	Max() int
	IsEmpty() bool
	Size() int
}

type bh struct {
	heap [max]int
	size int
}

// New ...
func New() BinaryHeap {
	return &bh{
		heap: [max]int{},
	}
}

// Insert ...
func (b *bh) Insert(int) {

}

// Max ...
func (b *bh) Max() int {
	return max
}

// IsEmpty ...
func (b *bh) IsEmpty() bool {
	return true
}

// Size ...
func (b *bh) Size() int {
	return b.size
}

func (b *bh) less(i, j int) bool {
	return i < j
}

func (b *bh) exchange(i, j int) {
	b.heap[i], b.heap[j] = b.heap[j], b.heap[i]
}

func (b *bh) swim(k int) {
	for k > 1 && b.less(k/2, k) {
		b.exchange(k/2, k)
		k = k / 2
	}
}

func (b *bh) sink(k int) {
	for 2*k <= b.size {
		j := 2 * k

		if j < b.size && b.less(j, j+1) {
			j++
		}
		if !b.less(k, j) {
			break
		}

		b.exchange(k, j)
		k = j
	}
}
