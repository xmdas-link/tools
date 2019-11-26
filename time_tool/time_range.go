package time_tool

import (
	"errors"
	"time"
)

func CheckTimeRange(beginTime, endTime string, format string) error {
	var (
		begin *time.Time
		end   *time.Time
	)
	if beginTime != "" {
		if t, err := Parse(format, beginTime); err != nil {
			return errors.New("开始时间格式错误")
		} else {
			begin = &t
		}
	}

	if endTime != "" {
		if t, err := Parse(format, endTime); err != nil {
			return errors.New("结束时间格式错误")
		} else {
			end = &t
		}

		if begin != nil {
			if end.Before(*begin) {
				return errors.New("开始时间不能早于结束时间")
			}
		}
	}

	return nil
}
