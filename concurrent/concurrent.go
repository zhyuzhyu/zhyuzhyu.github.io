package concurrent

import (
	"fmt"
	"sync"
)

func PrintNum(procNum, maxNum int) {
	if procNum <= 0 || maxNum <= 0 {
		return
	}

	var wg sync.WaitGroup
	cond := sync.NewCond(&sync.Mutex{})
	cur := 1

	wg.Add(procNum)
	for i := 0; i < procNum; i++ {
		go func(id int) {
			defer wg.Done()

			for {
				cond.L.Lock()
				for cur <= maxNum && (cur-1)%procNum != id {
					cond.Wait()
				}

				if cur > maxNum {
					cond.Broadcast()
					cond.L.Unlock()
					return
				}

				fmt.Println(cur)
				cur++
				cond.Broadcast()
				cond.L.Unlock()
			}
		}(i)
	}

	wg.Wait()
}

func PrintNum2(procNum, maxNum int) {
	wg := sync.WaitGroup{}
	cond := sync.Cond{L: &sync.Mutex{}}
	cur := 1

	wg.Add(procNum)
	for i := 0; i < procNum; i++ {
		go func(id int) {
			defer wg.Done()
			for {
				cond.L.Lock()
				if cur >= maxNum {
					cond.L.Unlock()
					cond.Broadcast()
					return
				}
				for cur%procNum != id {
					cond.Wait()
				}
				fmt.Println(cur)
				cur++
				cond.L.Unlock()
				cond.Broadcast()
			}
		}(i)
	}
	wg.Wait()
}
