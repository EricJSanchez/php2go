package php2go

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")

func Time() int64 {
	return time.Now().In(SysTimeLocation).Unix()
}

func StrToTime(input string, baseTime ...int64) int64 {
	var base time.Time
	if len(baseTime) > 0 {
		base = time.Unix(baseTime[0], 0).In(SysTimeLocation)
	} else {
		base = time.Now().In(SysTimeLocation)
	}

	// 处理相对时间格式
	if t, err := parseRelativeTime(input, base); err == nil {
		return t.Unix()
	}

	// 处理绝对时间格式
	if t, err := parseAbsoluteTime(input); err == nil {
		return t.Unix()
	}

	return 0
}

func parseRelativeTime(input string, base time.Time) (time.Time, error) {
	input = strings.ToLower(strings.TrimSpace(input))

	// 常见相对时间格式
	patterns := map[string]time.Duration{
		"now":          0,
		"today":        0,
		"tomorrow":     24 * time.Hour,
		"yesterday":    -24 * time.Hour,
		"next day":     24 * time.Hour,
		"previous day": -24 * time.Hour,
	}

	if dur, ok := patterns[input]; ok {
		return base.Add(dur), nil
	}

	// 处理数字+单位格式
	re := regexp.MustCompile(`^([+-]?\d+)\s*(second|minute|hour|day|week|month|year)s?$`)
	if matches := re.FindStringSubmatch(input); matches != nil {
		num, _ := strconv.Atoi(matches[1])
		unit := matches[2]

		var dur time.Duration
		switch unit {
		case "second":
			dur = time.Duration(num) * time.Second
		case "minute":
			dur = time.Duration(num) * time.Minute
		case "hour":
			dur = time.Duration(num) * time.Hour
		case "day":
			dur = time.Duration(num*24) * time.Hour
		case "week":
			dur = time.Duration(num*24*7) * time.Hour
		case "month":
			return base.AddDate(0, num, 0), nil
		case "year":
			return base.AddDate(num, 0, 0), nil
		}
		return base.Add(dur), nil
	}

	return time.Time{}, fmt.Errorf("不是有效的相对时间格式")
}

func parseAbsoluteTime(input string) (time.Time, error) {
	// 尝试常见日期格式
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02 15:04",
		"2006-01-02 15",
		"2006-01-02",
		"20060102150405",
		"200601021504",
		"2006010215",
		"20060102",
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
	}

	for _, format := range formats {
		if t, err := time.ParseInLocation(format, input, SysTimeLocation); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("不是有效的绝对时间格式")
}

func Date(format string, timestamp ...int64) string {
	var t time.Time
	if len(timestamp) > 0 {
		t = time.Unix(timestamp[0], 0).In(SysTimeLocation)
	} else {
		t = time.Now().In(SysTimeLocation)
	}
	replacer := strings.NewReplacer(
		// 日期
		"d", "02",
		"D", "Mon",
		// 月份
		"m", "01",
		"M", "Jan",
		// 年份（修正为使用时间戳对应的年份）
		"Y", "2006",
		"y", "06",
		// 时间
		"h", "03",
		"H", "15",
		"i", "04",
		"s", "05",
	)

	return t.Format(replacer.Replace(format))
}
