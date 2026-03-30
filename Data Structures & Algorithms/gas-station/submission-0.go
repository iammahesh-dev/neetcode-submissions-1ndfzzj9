func canCompleteCircuit(gas []int, cost []int) int {
    if sum(gas) < sum(cost){
		return -1
	}
	total := 0
	start := 0
	for i, v := range gas {
		total += v - cost[i]
		if total < 0 {
			total = 0
			start = i + 1
		}
	}
	return start
}

func sum(nums []int) int {
	s := 0
	for _,v := range nums {
		s += v
	}
	return s
}