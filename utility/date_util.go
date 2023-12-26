package utility

import (
	"time"

	"fuya-ark/internal/consts"
)

type dateUtils struct{}

var DateUtils dateUtils

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func (u dateUtils) GetBetweenDates(sdate, edate string) []string {
	var d []string
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil { // 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil { // 时间解析，异常
		return d
	}
	if date2.Before(date) { // 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

// GetBetweenDatesRank 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func (u dateUtils) GetBetweenDatesRank(sdate, edate string) []string {
	var d []string
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil { // 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil { // 时间解析，异常
		return d
	}
	if date2.Before(date) { // 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = consts.TimeLayoutDay
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

// GetFirstDateOfWeek 获取本周周一的日期
func (u dateUtils) GetFirstDateOfWeek(t time.Time) time.Time {
	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
}

// GetNextFirstDateOfWeek 获取下周周一
func (u dateUtils) GetNextFirstDateOfWeek(t time.Time) time.Time {
	return u.GetFirstDateOfWeek(t).AddDate(0, 0, 7)
}

// GetLastDateOfWeek 获取本周周日
func (u dateUtils) GetLastDateOfWeek(t time.Time) time.Time {
	return u.GetFirstDateOfWeek(t).AddDate(0, 0, 6)
}

// GetLastWeekFirstDate 获取上周的周一日期
func (u dateUtils) GetLastWeekFirstDate(t time.Time) time.Time {
	thisWeekMonday := u.GetFirstDateOfWeek(t)
	return thisWeekMonday.AddDate(0, 0, -7)
}

// GetLastWeekLastDate 获取下周周日
func (u dateUtils) GetLastWeekLastDate(t time.Time) time.Time {
	return u.GetFirstDateOfWeek(t).AddDate(0, 0, -7)
}

// GetFirstOfMonth 获取本月开始
func (u dateUtils) GetFirstOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// GetEndOfMonth 获取本月结束
func (u dateUtils) GetEndOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, t.Location())
}

// GetEndOfDay 获取本日结束
func (u dateUtils) GetEndOfDay(timeNow time.Time) time.Time {
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 23, 59, 59, 999, timeNow.Location()) // 获取当天0点时间 time类型
}

// GetStartOfDay 获取本日开始
func (u dateUtils) GetStartOfDay(timeNow time.Time) time.Time {
	return time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())
}

func (u dateUtils) ParesTime(dateStr, timeFormat string) time.Time {
	if len(timeFormat) != len(timeFormat) {
		timeFormat = timeFormat[0:len(dateStr)]
	}
	date, _ := time.Parse(timeFormat, dateStr)
	return date
}

func (u dateUtils) GetMondayAndSunday(weekDay int, date time.Time) (monDay, sunDay time.Time) {
	switch weekDay {
	case 0: //周日
		monDay = date.AddDate(0, 0, -6)
	case 1: //周一查上周
		monDay = date.AddDate(0, 0, -7)
	default:
		monDay = date.AddDate(0, 0, -weekDay+1)
	}
	sunDay = monDay.AddDate(0, 0, 6)
	return
}
