package concurrent

import (
	"fmt"
	"sync"
)

func PrintNum(procNum, maxNum int) {
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
