/*
 比较两个版本号 version1 和 version2。
 如果 version1 > version2 返回 1，如果 version1 < version2 返回 -1， 除此之外返回 0。

 你可以假设版本字符串非空，并且只包含数字和 . 字符。

  . 字符不代表小数点，而是用于分隔数字序列。

 例如，2.5 不是“两个半”，也不是“差一半到三”，而是第二版中的第五个小版本。

 你可以假设版本号的每一级的默认修订版号为 0。例如，版本号 3.4 的第一级（大版本）和第二级（小版本）修订号分别为 3 和 4。其第三级和第四级修订号均为 0。
*/

package one

import "strings"

func compareVersion(version1 string, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")
	min := len(v1)
	if len(v2) < len(v1) {
		min = len(v2)
	}
	for i := 0; i < min; i++ {
		if 转成数字(v1[i]) > 转成数字(v2[i]) {
			return 1
		} else if 转成数字(v1[i]) < 转成数字(v2[i]) {
			return -1
		}
	}
	for i := min; i < len(v1); i++ {
		if 转成数字(v1[i]) != 0 {
			return 1
		}
	}
	for i := min; i < len(v2); i++ {
		if 转成数字(v2[i]) != 0 {
			return -1
		}
	}
	return 0
}

func 转成数字(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		n = 10*n + int(s[i]-'0')
	}
	return n
}
