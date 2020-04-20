package 性能测试

import (
	"strconv"
	"testing"
)

//性能测试
//go test -bench . -cpuprofile cpu.out

//go tool pprof cpu.out(会进入交互式操作,执行help可以看到帮助)
//可以输入web生成web图，需安装www.graphviz.org(需重启电脑)
func BenchmarkT(t *testing.B) {
	m := make(map[int]string)
	for i := 1; i <= 1000000; i++ {
		m[i] = strconv.Itoa(i)
	}
}
