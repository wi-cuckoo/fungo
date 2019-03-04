/*
给定一个整数数组 A，找到 min(B) 的总和，其中 B 的范围为 A 的每个（连续）子数组。

由于答案可能很大，因此返回答案模 10^9 + 7。



示例：

输入：[3,1,2,4]
输出：17
解释：
子数组为 [3]，[1]，[2]，[4]，[3,1]，[1,2]，[2,4]，[3,1,2]，[1,2,4]，[3,1,2,4]。
最小值为 3，1，2，4，1，1，2，1，1，1，和为 17。


提示：

1 <= A <= 30000
1 <= A[i] <= 30000
*/

package nine

// 归纳法可以知道每个元素 A[i] 被加入到最终结果的次数为
// 其左邻居中大于它、右邻居中不小于它的元素长度 left 、 right （包括自身）的乘积
func sumSubarrayMins(A []int) int {
	if len(A) == 0 {
		return 0
	}
	res := A[0]
	fn := make([]int, 1, len(A))
	fn[0] = A[0]
	for i := 1; i < len(A); i++ {
		sum := 0
		for j := 0; j < len(fn); j++ {
			if A[i] < fn[j] {
				fn[j] = A[i]
			}
			sum += fn[j]
		}
		fn = append(fn, A[i])
		res += sum + A[i]
		// dp[i] = dp[i-1] + (i+1)*A[i]
	}
	return res % (1e+09 + 7)
}
