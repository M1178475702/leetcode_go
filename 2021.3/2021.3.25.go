package _2021_3

func qsort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	pos := setDuePos(nums)
	qsort(nums[:pos])
	qsort(nums[pos+1:])
}

func setDuePos(nums []int) int {
 	b := nums[0]
	left, right := 0, len(nums)-1

	for left < right {
		for ; left < right; right-- {
			if nums[right] < b {
				nums[left] = nums[right]
				left++
				break
			}
		}

		for ; left < right; left++ {
			if nums[left] > b {
				nums[right] = nums[left]
				right--
				break
			}
		}
	}
	nums[left] = b
	return left
}
