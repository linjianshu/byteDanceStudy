package main

import (
	"bufio"
	"github.com/bytedance/gopkg/lang/fastrand"
	"math/rand"
	"os"
	"strings"
)

func main() {

}

func HelloTom() string {
	return "Tom"
}

func JudgePassLine(score int16) bool {
	if score >= 60 {
		return true
	}

	return false
}

func ReadFirstLine() string {
	open, err := os.Open("E:\\project\\GOproject\\src\\log.txt")
	defer open.Close()
	if err != nil {
		return ""
	}
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func ProcessFirstLine() string {
	line := ReadFirstLine()
	destLine := strings.ReplaceAll(line, "11", "00")
	return destLine
}

var ServerIndex [10]int

func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i + 100
	}
}

func Select() int {
	return ServerIndex[rand.Intn(10)]
}

func FastSelect() int {
	return ServerIndex[fastrand.Intn(10)]
}
