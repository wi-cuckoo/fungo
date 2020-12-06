package zero

import (
	"github.com/wi-cuckoo/fungo/util"
)

/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png)

上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色
部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
*/

// 采用层消法，老子取的名字
// 除去首尾的 0，剩下的是有效接水的段。统计其中 0 的个数，即为最下一层可接水数
// 将所有非零的柱子长度减一，递归进行下一轮统计

func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	l, r := 0, len(height)-1
	for l < r {
		if height[l] == 0 {
			l++
		}
		if height[r] == 0 {
			r--
		}
		if height[l] > 0 && height[r] > 0 {
			break
		}
	}
	if r <= l {
		return 0
	}
	water := 0
	for i := l; i <= r; i++ {
		if height[i] == 0 {
			water++
			continue
		}
		height[i]--
	}
	// fmt.Println(water, l, r, height)
	return water + trap(height[l:r+1])
}

// 另一种思路
// 这想法有点妙呀
// 对每个柱子，看其左右最高的两根柱子，如果其中较小的高出该柱子，则表示最后接水时，这根柱子
// 上能接到的水高度就会等于较小的柱子与其高度差
func trapV2(height []int) int {
	l := len(height)
	left, right := make([]int, l), make([]int, l)
	max := 0
	for i := 0; i < l; i++ {
		if height[i] > max {
			max = height[i]
		}
		left[i] = max
	}
	max = 0
	for j := l - 1; j >= 0; j-- {
		if height[j] > max {
			max = height[j]
		}
		right[j] = max
	}
	sum := 0
	for i := 1; i < l-1; i++ {
		min := left[i-1]
		if right[i+1] < min {
			min = right[i]
		}
		m := min - height[i]
		if m > 0 {
			sum += m
		}
	}
	return sum
}

// 双指针法: 对与 trapV2 中的方法还可以进一步优化
// 对于 i 点能接多少水取决于左右两边最高柱子中较矮的那个
// 我们可以在遍历过程中维护两个变量, lmax: 到目前为止左侧最高的柱子，rmax: 到目前为止右侧最高的柱子
// 当从左往右遍历到 i 点时，对于 i 点来说，lmax 是可信的，rmax 是不可信的，如果是 lmax < rmax
// 那么可以确定 i 点能接多少水了（因为右侧 rmax 只可能更大），而如果 lmax > rmax 那么我们就可以确定
// 从右往左遍历的当前点
func trapV3(height []int) int {
	lmax, rmax := 0, 0
	l, r := 0, len(height)-1
	ans := 0
	for l <= r {
		if lmax < rmax {
			ans += util.Max(0, lmax-height[l])
			lmax = util.Max(lmax, height[l])
			l++
			continue
		}
		ans += util.Max(0, rmax-height[r])
		rmax = util.Max(rmax, height[r])
		r--
	}
	return ans
}
