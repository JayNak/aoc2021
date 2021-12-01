package main

func Day1() (int, int) {

	vals := ReadToInts("../data/day1.txt")

	count := 0
	last := -1

	for _, v := range vals {
		if last != -1 && last < v {
			count++
		}

		last = v
	}

	count2 := 0
	last2 := -1

	for i := 0; i < len(vals)-2; i++ {
		this := vals[i] + vals[i+1] + vals[i+2]

		if last2 != -1 && last2 < this {
			count2++
		}

		last2 = this
	}

	return count, count2

}
