package time_tool

import (
	"fmt"
	"time"
)

const Layout_YYYY_MM_DD = "2006-01-02"
const Layout_YMDHIS = "20060102150405"
const Layout_YMD_HIS = "2006-01-02 15:04:05"
const Layout_YYYY_MM = "2006-01"
const Layout_YMD = "2006-01-02"

var (
	location *time.Location
)

type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	timeVar := time.Time(t)
	stamp := ""
	if timeVar.IsZero() {
		stamp = `""` // 0000-00-00 00:00:00 零时间显示空字符串
	} else {
		stamp = fmt.Sprintf("\"%s\"", timeVar.Format(Layout_YMD_HIS))
	}
	return []byte(stamp), nil
}

func init() {
	l, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	location = l
}

func Parse(layout string, timeStr string) (time.Time, error) {
	t, err := time.ParseInLocation(layout, timeStr, location)
	return t, err
}

func ParseDateYMD(timeStr string) (time.Time, error) {
	t, err := time.ParseInLocation(Layout_YYYY_MM_DD, timeStr, location)
	return t, err
}

func ParseDate(format string, date string, defaultTime time.Time) time.Time {
	if t, err := Parse(format, date); err == nil {
		return t
	}
	t, _ := Parse(format, defaultTime.Format(format))
	return t
}

func FormatTime(format string, time time.Time) time.Time {
	t, _ := Parse(format, time.Format(format))
	return t
}

func FormatTimeStamp(format string, timestamp int64) string {

	dateTime := time.Unix(timestamp, 0)

	return dateTime.Format(format)
}
