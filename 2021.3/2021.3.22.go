package _2021_3


func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		if num & 1 == 1 {
			cnt++
		}
		num >>= 1
	}
	return cnt
}



