package four

/*
给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。

进阶：你可以在 O(n) 的时间解决这个问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 方法1：暴力需要n*n复杂度，但 nums[i]^nums[j] 与 nums[j]^nums[i] 属于重复计算可以优化掉
// 代码略

// 方法2：使用字典树来解决
