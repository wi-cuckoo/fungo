package seven

/*
给定一个整数数组 temperatures ，表示每天的温度，返回一个数组 answer ，其中 answer[i] 是指在第 i 天之后，才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 0 来代替。

 

示例 1:

输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
示例 2:

输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]
示例 3:

输入: temperatures = [30,60,90]
输出: [1,1,0]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/daily-temperatures
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// checkou monotonic-stack
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	ans := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		// top less than current, pop
		for n := len(stack); n > 0; n-- {
			idx := stack[n-1]
			top := temperatures[idx] 
			if top >= temperatures[i] {
				break
			}
			if top < temperatures[i] {
				stack = stack[:n-1]
				ans[idx] = i - idx
			}
		}
		stack = append(stack, i)
		// fmt.Println(stack)
	}
	return ans
}