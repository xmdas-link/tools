package time_tool

import "time"

/**
 * t1的日期是否早于t2的日期（不是简单对比时间）
 */
func DayBefore(t1 time.Time, t2 time.Time) bool {
	dt1, _ := Parse(Layout_YYYY_MM_DD, t1.Format(Layout_YYYY_MM_DD))
	dt2, _ := Parse(Layout_YYYY_MM_DD, t2.Format(Layout_YYYY_MM_DD))
	return dt1.Before(dt2)
}
