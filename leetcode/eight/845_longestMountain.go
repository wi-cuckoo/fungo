/*
我们把数组 A 中符合下列属性的任意连续子数组 B 称为 “山脉”：

B.length >= 3
存在 0 < i < B.length - 1 使得 B[0] < B[1] < ... B[i-1] < B[i] > B[i+1] > ... > B[B.length - 1]
（注意：B 可以是 A 的任意子数组，包括整个数组 A。）

给出一个整数数组 A，返回最长 “山脉” 的长度。

如果不含有 “山脉” 则返回 0。



示例 1：

输入：[2,1,4,7,3,2,5]
输出：5
解释：最长的 “山脉” 是 [1,4,7,3,2]，长度为 5。
示例 2：

输入：[2,2,2]
输出：0
解释：不含 “山脉”。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-mountain-in-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package eight

// 左右数组方法：
// 对A中每个数字计算出其左边降序数组长度放入左数组 left
// 对A中每个数字计算出其右边降序数组长度放入右数组 right
// 那么对于每个数字 A[i],当 left[i] 与 right[i] 均不为 0 时, i 点成为峰顶。则以其为峰顶的山脉
// 长度为 left[i]+right[i]+1。最后更新该长度得到最大山脉长度

func longestMountain(A []int) int {
	left := make([]int, len(A))
	cnt := 0
	for i := 1; i < len(A); i++ {
		if A[i] > A[i-1] {
			cnt++
		} else {
			cnt = 0
		}
		left[i] = cnt
	}
	cnt = 0
	right := make([]int, len(A))
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] > A[i+1] {
			cnt++
		} else {
			cnt = 0
		}
		right[i] = cnt
	}
	max := 0
	for i := 1; i < len(A)-1; i++ {
		if left[i] == 0 || right[i] == 0 {
			continue
		}
		cur := left[i] + 1 + right[i]
		if cur > max {
			max = cur
		}
	}
	return max
}
