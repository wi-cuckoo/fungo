> 先说在前面，在竞赛过程中，首先要保证过然后才是优化算法，所以暴力法不是不可以，当你没有思路的时候想想暴力怎么解

### 1. 判断能否形成等差数列

> https://leetcode-cn.com/problems/can-make-arithmetic-progression-from-sequence/

题干：
```
给你一个数字数组 arr 。
如果一个数列中，任意相邻两项的差总等于同一个常数，那么这个数列就称为 等差数列 。
如果可以重新排列数组形成等差数列，请返回 true ；否则，返回 false 
```

思路：
- 这题没啥研究的，直观来讲肯定排序后的数组就可判断了，那就先排序
- 然后优化点我觉得可以选用适合的排序把等差判断和排序放一起，比如选择排序

代码：

```golang
func canMakeArithmeticProgression(arr []int) bool {
    if len(arr) < 3 {
        return true
    }
    sort.Ints(arr)
    d := arr[1] - arr[0]
    for i := 2; i < len(arr); i++ {
        if d != arr[i]-arr[i-1] {
            return false
        }
    }
    return true
}
```

### 2. 所有蚂蚁掉下来前的最后一刻

> https://leetcode-cn.com/problems/last-moment-before-all-ants-fall-out-of-a-plank/


题干：
```
有一块木板，长度为 n 个 单位 。一些蚂蚁在木板上移动，每只蚂蚁都以 每秒一个单位 的速度移动。其中，一部分蚂蚁向 左 移动，其他蚂蚁向 右 移动。

当两只向 不同 方向移动的蚂蚁在某个点相遇时，它们会同时改变移动方向并继续移动。假设更改方向不会花费任何额外时间。

而当蚂蚁在某一时刻 t 到达木板的一端时，它立即从木板上掉下来。

给你一个整数 n 和两个整数数组 left 以及 right 。两个数组分别标识向左或者向右移动的蚂蚁在 t = 0 时的位置。请你返回最后一只蚂蚁从木板上掉下来的时刻。
```

思路：
- 此题一开始想暴力的化就是模拟整个过程，但且不说模拟相对复杂，计算量我觉得也是天价
- 题目要求一个最终的耗时结果，并不care中间蚂蚁们怎么碰头掉头啥的，所以必然有简单解法
- 我想从输入的left, right数组入手，却只想明白了如果碰撞，则是对应蚂蚁的一次两数组值交换
- 看了题解才发现，这里的关键点就是碰撞掉头后如果并不区分两只蚂蚁的话，它俩就交换了生命而已，则等同于穿透了
- 理解了上面穿透的意思，题目就简单了，向左的还是一直向左，向右的还是一直向右，没啥变化

代码：
```golang
func getLastMoment(n int, left []int, right []int) int {
    m1, m2 := 0, n
    for _, m := range left {
        if m > m1 {
            m1 = m
        }
    }
    for _, m := range right {
        if m < m2 {
            m2 = m
        }
    }
    m2 = n- m2
    if m1 > m2 {
        return m1
    }
    return m2
}
```

### 3. 统计全 1 子矩形

> https://leetcode-cn.com/problems/count-submatrices-with-all-ones/

题干：
给你一个只包含 0 和 1 的 rows * columns 矩阵 mat ，请你返回有多少个 子矩形 的元素全部都是 1 

思路：
- 这题和之前一道统计正方形的题目类似，只不过把正方形换成了矩形
- 可以按行来DP，对于点 [i, j]，统计其所在行往左连续的 1 的数量记为 dp[i][j]
- 统计完后，还是对于点 [i, j]，将其设为矩形右下顶点，往上通过dp数组来寻找有效的矩形数量
- 累加到结果里面完事

代码：
```golang
func numSubmat(mat [][]int) int {
	dp := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		dp[i] = make([]int, len(mat[i]))
		cnt := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				dp[i][j], cnt = 0, 0
				continue
			}
			cnt++
			dp[i][j] = cnt
		}
	}
	total := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if dp[i][j] == 0 {
				continue
			}
			min := len(mat[i]) + 1
			for k := i; k >= 0; k-- {
				if dp[k][j] < min {
					min = dp[k][j]
				}
				total += min
			}
		}
	}
	return total
}
```

### 4. 最多 K 次交换相邻数位后得到的最小整数

> https://leetcode-cn.com/problems/minimum-possible-integer-after-at-most-k-adjacent-swaps-on-digits/

题干：
给你一个字符串 num 和一个整数 k 。其中，num 表示一个很大的整数，字符串中的每个字符依次对应整数上的各个数位 。
你可以交换这个整数相邻数位的数字，最多 k 次。
请你返回你能得到的最小整数，并以字符串形式返回。

思路：
- 肯定是在有限的次数中尽可能把较小的数字排到前面去，暴力的思路如下
    * 找到最小的数字，计算一下排到第一位去需要交换多少次，记为 k' 次
    * 那么当满足 k > k' 时，咱们就完成这笔交易，将其插入到第一位去
    * 这次移动完后，对剩下的 S[1:] 字符串也应用同样的方法，直到 k 次用完
- 暴力方法挺好理解，但有没有更好的方法呢