func multiThreadAdd(v) {
	m := runtime.numOfCPU();
	var wg sync.WaitGroup
	n := len(v)
		res := [m]int{}
	for i := 0; i < m; i++ {
		wg.Add(1)
		go func(i int, j int, []float64 arr) {
			defer wg.Done()
			for _, x := range arr[i:j] {
				res[i/(n/m)] += x
			}
		}(i * (n/m), i * (n/m) + n/m, v)
	}
	wg.Wait()
	var ans float64
	for _, x := range res {
		ans += x
	}
	return ans
}
