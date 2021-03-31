package fan

import "sync"

func Generate() <-chan []int {
	out := make(chan []int)
	go func() {
		defer close(out)
		data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		for _, v := range data {
			out <- v
		}
	}()
	return out
}

func Average(in <-chan []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- avg(v)
		}
	}()
	return out
}

func avg(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	result := sum / len(nums)
	return result
}

func CaMerge(ch ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(ch))

	for _, c := range ch {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
