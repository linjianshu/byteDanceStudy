package main

import (
	"bou.ke/monkey"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectedoutput := "Tom"
	if output != expectedoutput {
		t.Errorf("Expected %s do not match actual %s", expectedoutput, output)
	}
}

func TestJudgePassLineTrue(t *testing.T) {
	isPass := JudgePassLine(70)
	assert.Equal(t, true, isPass)
}

//覆盖率测试 对各个分支进行测试
func TestJudgePassLineFalse(t *testing.T) {
	isPass := JudgePassLine(50)
	assert.Equal(t, false, isPass)
}

func TestProcessFirstLine(t *testing.T) {
	firstLine := ProcessFirstLine()
	assert.Equal(t, "line00", firstLine)
}

//mock 对readline 进行打桩测试 不再依赖本地文件
func TestProcessFirstLineWithMock(t *testing.T) {
	monkey.Patch(ReadFirstLine, func() string {
		return "line110"
	})
	defer monkey.Unpatch(ReadFirstLine)

	line := ProcessFirstLine()
	assert.Equal(t, "line000", line)
}

//基准测试 串行
func BenchmarkSelect(b *testing.B) {
	InitServerIndex()
	//重置 时间
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Select()
	}
}

func BenchmarkSelectParallel(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Select()
		}
	})
}

func BenchmarkFastSelectParallel(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			FastSelect()
		}
	})
}
