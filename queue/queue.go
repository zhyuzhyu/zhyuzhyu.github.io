package queue

import (
	"container/heap"
)

type Pair struct {
	key, value int
}

type PriorityQueuePair []*Pair

func (pq PriorityQueuePair) Len() int           { return len(pq) }
func (pq PriorityQueuePair) Less(i, j int) bool { return pq[i].value < pq[j].value }
func (pq PriorityQueuePair) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueuePair) Push(x interface{}) {
	*pq = append(*pq, x.(*Pair))
}
func (pq *PriorityQueuePair) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// https://leetcode.cn/problems/sliding-window-maximum/description/
func maxSlidingWindow(nums []int, k int) []int {
	// nums = [-7, -8, 7, 5, 7, 1, 6, 0], k = 4
	pq := &PriorityQueuePair{}
	rets := []int{}
	for i := 0; i < len(nums); i++ {
		heap.Push(pq, &Pair{nums[i], i})
		if i < k-1 {
			continue
		}

		for (*pq)[0].key < i-k+1 {
			heap.Pop(pq)
		}
		rets = append(rets, (*pq)[0].value)
	}
	return rets
}

// https://leetcode.cn/problems/top-k-frequent-elements/description/
func topKFrequent(nums []int, k int) []int {
	// nums = [1,1,1,2,2,3], k = 2
	pq := &PriorityQueuePair{}
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	for num, v := range mp {
		heap.Push(pq, &Pair{num, v})
		if len(*pq) > k {
			heap.Pop(pq)
		}
	}
	rets := []int{}
	for pq.Len() > 0 {
		rets = append(rets, heap.Pop(pq).(*Pair).key)
	}
	return rets
}
