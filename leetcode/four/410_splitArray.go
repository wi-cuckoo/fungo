/*
给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。

注意:
数组长度 n 满足以下条件:

1 ≤ n ≤ 1000
1 ≤ m ≤ min(50, n)
示例:

输入:
nums = [7,2,5,10,8]
m = 2

输出:
18

解释:
一共有四种方法将nums分割为2个子数组。
其中最好的方式是将其分为[7,2,5] 和 [10,8]，
因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。
*/

package four

/*
 * 注：分割分数即是各子数组和的最大值中最小的辣一个
 * 1. 将 n 个数划分为 n 段，分割分数即为 n 个正整数的最大值
 * 2. 将 n 个数划分为 1 段，分割分数即为 n 个正整数的和
 * 3. 将 n 个数划分为 m 段，则分割分数一定介于最大值与和之间。因此采用二分查找，
 * 	  每次取一个值对序列进行划分，若能划分出m 段，使得每段的和都小于这个数值，则满
 *    足划分，反之不满足，直至找到一个最小的满足划分的值为止。
 */

func splitArray(nums []int, m int) int {
	l, r := 0, 0 // 分别表示划分为n段和1段时的分割分数
	for _, n := range nums {
		if n > l {
			l = n
		}
		r += n
	}

	for l < r {
		// 这里使用二分法，从最大与最小分割分数中间值开始探索
		mid := (r + l) / 2
		// 判断数组能不能被划分为m段，并且每段大小不高于 mid 值
		// 如果能，则继续降低预期值，注意不能 r = mid-1 是因为无法
		// 确保减一后的数组还满足划分不
		// 如果不能满足，则从mid到r之间寻找，即 l = mid+1
		if fuck(nums, mid, m) {
			r = mid
			continue
		}
		l = mid + 1
	}

	return l
}

func fuck(nums []int, mid, m int) bool {
	cnt, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		if sum+nums[i] > mid {
			sum = nums[i]
			cnt++
			// 当分了 m 段后sum有剩余则划分不成功
			if cnt > m-1 {
				return false
			}
			continue
		}
		sum += nums[i]
	}
	return true
}
