package test

import (
	"github.com/EricJSanchez/php2go"
	"testing"
)

func TestTime(t *testing.T) {
	// 测试基本结构体
	php2go.Pr(php2go.Time())
	t1 := php2go.StrToTime("2025-01-02")
	t11 := php2go.Date("Y-m-d", t1)
	php2go.Pr(t1, t11)
	if t11 != "2025-01-02" {
		t.Error("时间转换出错")
	}
	php2go.Pr(php2go.Date("YmdHi"))

	t2 := php2go.StrToTime(php2go.Date("YmdHi"))
	php2go.Pr("ts", t2)
}
