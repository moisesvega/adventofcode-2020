package main

func day1(values []int, target int) (out int) {
	cache := map[int]bool{}
	var max int
	max = 0

	for _, value := range values {
		if _, ok := cache[target-value]; ok {
			tmp := value * (target - value)
			if tmp > max {
				max = tmp
			}
		}
		cache[value] = true
	}

	return max
}

func day1Pt2(values []int, target int) (out int) {
	triplets := map[int]int{}
	var max int

	for i, value := range values {
		if v, ok := triplets[target-value]; ok {
			if v > max {
				max = v
			}
			continue
		}

		a, b := tuple(remove(values, i), target-value)
		v := a * b * value
		// little nice improvement to not calculate any of these values
		triplets[target-value] = v
		triplets[target-a] = v
		triplets[target-b] = v
	}

	return max
}

func tuple(values []int, target int) (int, int) {
	cache := map[int]bool{}
	var (
		max  int
		out  int
		out2 int
	)
	out, max, out2 = 0, 0, 0

	for _, value := range values {
		if _, ok := cache[target-value]; ok {
			tmp := value * (target - value)
			if tmp > max {
				max = tmp
				out = value
				out2 = target - value
			}
		}
		cache[value] = true
	}

	return out, out2
}

func remove(in []int, s int) []int {
	slice := make([]int, len(in))
	copy(slice, in)
	if s == len(in)-1 {
		return append(slice[:s-1])
	}
	return append(slice[:s], slice[s+1:]...)
}
