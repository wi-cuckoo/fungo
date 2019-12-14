package zero

/*
以 Unix 风格给出一个文件的绝对路径，你需要简化它。或者换句话说，将其转换为规范路径。

在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录）；两者都可以是复杂相对路径的组成部分。更多信息请参阅：Linux / Unix中的绝对路径 vs 相对路径

请注意，返回的规范路径必须始终以斜杠 / 开头，并且两个目录名之间必须只有一个斜杠 /。最后一个目录名（如果存在）不能以 / 结尾。此外，规范路径必须是表示绝对路径的最短字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/simplify-path
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

import (
	"path/filepath"
	"strings"
)

// 用栈来做
// 1. 先将字符串以 / 分割成数组
// 2. 遍历数组，遇到非 '', '.', '..'的情况入栈，遇到 '..' 则出栈（如果栈不为空）
// 3. 完事后检查栈是否为空，为空返回 '/', 否则用 '/' 将栈里面元素拼接起来即可

func simplifyPath(path string) string {
	filepath.Abs(path)
	segs := strings.Split(path, "/")
	stack := make([]string, 0, len(segs))
	abs := ""
	for _, seg := range segs {
		if seg != "" && seg != "." && seg != ".." {
			stack = append(stack, seg)
			continue
		}
		if seg == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	abs = strings.Join(stack, "/")
	return "/" + abs
}
