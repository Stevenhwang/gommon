package times

import "time"

var defaultLayout = "2006-01-02 15:04:05"
var defaultTimezone = "Asia/Shanghai"

type TimeOptions struct {
	Layout   string
	Timezone string
}

// GetTimeString 获取以 offset 天数为偏移量的格式化字符串
func GetTimeString(offset int, opts TimeOptions) (string, error) {
	var layout = defaultLayout
	var timezone = defaultTimezone
	if len(opts.Layout) != 0 {
		layout = opts.Layout
	}
	if len(opts.Timezone) != 0 {
		timezone = opts.Timezone
	}
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return "", err
	}
	now := time.Now().In(location).Add(time.Duration(offset) * time.Hour * 24)
	return now.Format(layout), nil
}
