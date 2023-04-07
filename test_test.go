package timx

import (
	"testing"
	"time"
)

func panicIfErr(err error, code string) {
	if err == nil {
		return
	}
	panic(code + " " + err.Error())
}
func testEqual(isEqual bool, code string) {
	if isEqual {
		return
	}
	panic(code)
}

const (
	dayFormat  = "2006-01-02"
	timeFormat = "2006-01-02 15:04:05"
)

func Test_All(t *testing.T) {
	startTime, err := time.ParseInLocation("2006-01-02", "2000-01-01", time.UTC)
	panicIfErr(err, "zzzdqo5hb1")
	obj := New(startTime)
	testEqual(obj.DayStart().Format(timeFormat) == "2000-01-01 00:00:00", "05sg2t8ekm")
	testEqual(obj.DayEnd().Format(timeFormat) == "2000-01-01 23:59:59", "ur6da7hs1i")
	testEqual(obj.AddDay(2).Format(timeFormat) == "2000-01-03 00:00:00", "qjmto0svw9")
	testEqual(obj.NextDay().Format(dayFormat) == "2000-01-02", "rhsz116nkr")
	testEqual(obj.LastDay().Format(dayFormat) == "1999-12-31", "22060v86zh")

	testEqual(obj.WeekStart().Format(dayFormat) == "1999-12-26", "78zd2comjl")
	testEqual(obj.WeekEnd().Format(dayFormat) == "2000-01-01", "b7qvhtet8s")
	testEqual(obj.AddWeek(2).Format(dayFormat) == "2000-01-15", "9wnqxco8zc")
	testEqual(obj.NextWeek().Format(dayFormat) == "2000-01-08", "9199bmz908")
	testEqual(obj.LastWeek().Format(dayFormat) == "1999-12-25", "j2a82ugs75")

	testEqual(obj.MonthStart().Format(dayFormat) == "2000-01-01", "eeb53ymndz")
	testEqual(obj.MonthEnd().Format(dayFormat) == "2000-01-31", "py91j1q4hf")
	testEqual(obj.AddMonth(2).Format(dayFormat) == "2000-03-01", "98w3ijcf9i")
	testEqual(obj.NextMonth().Format(dayFormat) == "2000-02-01", "ljlvrn2i3k")
	testEqual(obj.LastMonth().Format(dayFormat) == "1999-12-01", "fhxdcwtpx8")

	testEqual(obj.Quarter() == 1, "22hrs3jzal")
	testEqual(obj.QuarterStart().Format(dayFormat) == "2000-01-01", "22hrs3jzal")
	testEqual(obj.QuarterEnd().Format(dayFormat) == "2000-03-01", "qu2g543yzi")
	testEqual(obj.AddQuarter(2).Format(dayFormat) == "2000-07-01", "gzveocwjzw")
	testEqual(obj.LastQuarter().Format(dayFormat) == "1999-10-01", "zfrxgecx4g")
	testEqual(obj.NextQuarter().Format(dayFormat) == "2000-04-01", "s4mvvy2450")

	testEqual(obj.YearStart().Format(dayFormat) == "2000-01-01", "9jsch1i2ca")
	testEqual(obj.YearEnd().Format(dayFormat) == "2000-12-31", "q916pkmo75")
	testEqual(obj.AddYear(2).Format(dayFormat) == "2002-01-01", "lvhycakz9e")
	testEqual(obj.LastYear().Format(dayFormat) == "1999-01-01", "g95e001elp")
	testEqual(obj.NextYear().Format(dayFormat) == "2001-01-01", "bfilyb36fb")

	startTime, err = time.ParseInLocation("2006-01-02", "2022-12-31", time.UTC)
	panicIfErr(err, "zzzdqo5hb1")
	testEqual(New(startTime).MonthEnd().Format(dayFormat) == "2022-12-31", "py91j1q4hf")
}
