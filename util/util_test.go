package util

import (
	"testing"
)

// 测试用例参考

// 加入Test
// 如果 GenShortId() 返回的 shortId 为空或者 err 不为空，则调用 t.Error() 函数标明该用例测试不通过。
func TestGenShortId(t *testing.T) {
	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		t.Error("GenShortId failed!")
	}

	t.Log("GenShortId test pass")
}

// 性能测试用例 加入 Benchmark，go test -test.bench=".*"
// 查看性能并生成函数调用图 go test -bench=".*" -cpuprofile=cpu.profile ./util
// 执行 go tool pprof util.test cpu.profile 查看性能（进入交互界面后执行 top 指令）：
func BenchmarkGenShortId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}

func BenchmarkGenShortIdTimeConsuming(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	shortId, err := GenShortId()
	if shortId == "" || err != nil {
		b.Error(err)
	}

	b.StartTimer() //重新开始时间

	for i := 0; i < b.N; i++ {
		GenShortId()
	}
}

// 覆盖率命令参考
// go test -coverprofile=cover.out 覆盖率测试
// go tool cover -func=cover.out
