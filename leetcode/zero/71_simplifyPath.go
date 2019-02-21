package zero

import (
	"path/filepath"
	"strings"
)

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
