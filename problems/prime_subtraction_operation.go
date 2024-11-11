package problems

func arrayPrimeStrictlyLessThanNum(num int) []int {
	res := []int{}
	if num <= 2 {
		return []int{0}
	}
	n := num - 1
	for n > 2 {
		flag := true
		for i := 2; i < n; i++ {
			if n%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, n)
		}
		n--
	}
	res = append(res, 2)
	return res
}
func isStrictlyIncreasingArray(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}
func findMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
func PrimeSubOperation(nums []int) bool {
	if isStrictlyIncreasingArray(nums) {
		return true
	}
	m := findMax(nums)
	arrM := arrayPrimeStrictlyLessThanNum(m)

	if nums[0] > 2 {
		i := 0
		for nums[0]-arrM[i] <= 0 {
			i++
		}
		nums[0] = nums[0] - arrM[i]
	}

	if isStrictlyIncreasingArray(nums) {
		return true
	}
	for i := 1; i < len(nums); i++ {
		for _, p := range arrM {
			if nums[i]-p > nums[i-1] {
				nums[i] = nums[i] - p
				break
			}
		}
		if isStrictlyIncreasingArray(nums) {
			return true
		}
	}
	return false
}
