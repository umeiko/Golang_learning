package main

import "fmt"

func if_k_in_keys(in_map map[int]int, k int) bool {
	for key := range in_map {
		if k == key {
			return true
		}
	}
	return false
}
func if_v_in_values(in_map map[int]int, v int) bool {
	for _, value := range in_map {
		if v == value {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func maxEqualFreq_slow(nums []int) (result int) {
	if len(nums) < 4 {
		return len(nums)
	}
	result = 0
	nums_freq := map[int]int{} // 数字:频次
	freqs := map[int]int{}     // 频次:数量
	for k, i := range nums {
		// 增加该数字的频次统计, 该数字之前所处的频次减少一个
		if if_k_in_keys(nums_freq, i) {
			freqs[nums_freq[i]] -= 1
			if freqs[nums_freq[i]] == 0 {
				delete(freqs, nums_freq[i])
			}
			nums_freq[i] += 1
		} else {
			nums_freq[i] = 1
		}
		// 统计频次，该数字目前的频次增加一个
		if if_k_in_keys(freqs, nums_freq[i]) {
			freqs[nums_freq[i]] += 1
		} else {
			freqs[nums_freq[i]] = 1
		}
		// 判断是否满足要求
		if len(freqs) == 1 {
			// 只有1种频次，且这个频次就是1（不同的数都出现一次）
			if if_k_in_keys(freqs, 1) {
				result = k + 1
			}
			// 这个频次不是1，但只有一个（所有的都是同一个数）
			if if_v_in_values(freqs, 1) {
				result = k + 1
			}
		}
		if len(freqs) == 2 {
			// 有只出现一次的数，且这个数只有一个
			if if_k_in_keys(freqs, 1) {
				if freqs[1] == 1 {
					result = k + 1
					continue
				}
			}
			// 只有出现n次的数和出现n-1次的数，且n次的数只有一个
			freq_keys := make([]int, 0, 2)
			for freq_key := range freqs {
				freq_keys = append(freq_keys, freq_key)
			}
			if (freq_keys[0]-freq_keys[1]) == 1 || (freq_keys[0]-freq_keys[1]) == -1 {
				if freqs[max(freq_keys[0], freq_keys[1])] == 1 {
					result = k + 1
				}
			}

		}
	}
	return
}

func maxEqualFreq(nums []int) (result int) {
	if len(nums) < 4 {
		return len(nums)
	}
	result = 0
	max_freq := 0
	nums_freq := map[int]int{} // 数字:频次
	freqs := map[int]int{}     // 频次:数量

	for k, i := range nums {
		if nums_freq[i] > 0 {
			freqs[nums_freq[i]]--
		}
		nums_freq[i]++
		freqs[nums_freq[i]]++
		max_freq = max(nums_freq[i], max_freq)
		if max_freq == 1 || // 所有数都只有一个
			freqs[max_freq]*max_freq == k && freqs[1] == 1 || // 除1个数是1次以外，别的数都是n次
			freqs[max_freq]*max_freq+freqs[max_freq-1]*(max_freq-1) == k+1 && freqs[max_freq] == 1 {
			// 所有数都是n和n-1次，且n次只有一组
			result = k + 1
		}
	}
	return
}

func main() {
	a := []int{1, 1, 1, 1, 1, 1, 1}
	fmt.Print(maxEqualFreq(a))
	fmt.Print(maxEqualFreq_slow(a))
	// a := map[string]string{}
	// a["a"] += "hello"
	// fmt.Print(a["c"] == "")

}

// 在Go的map中，可以直接枚举并进行运算,不会出现“not defined key”的问题。
