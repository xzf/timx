package timx

import "time"

var gWeekFirstDay time.Weekday

//SetWeekStart
//golang default set sunday as the first day of a week
//but some area use other day as first day
//so you can just set the first day you want
//does not support using multiple first day of week in one system
func (t Time) SetWeekStart(weekday time.Weekday) {
	gWeekFirstDay = weekday
}

func (t Time) WeekStart() Time {
	for t.Weekday() != gWeekFirstDay {
		t.Time = t.LastDay().Time
	}
	return t.DayStart()
}

func (t Time) WeekEnd() Time {
	for t.Weekday() != gWeekFirstDay {
		t.Time = t.NextDay().Time
	}
	return t.LastDay().DayStart()
}

func (t Time) AddWeek(n int) Time {
	return t.AddDay(7 * n)
}

func (t Time) LastWeek() Time {
	return t.AddWeek(-1)
}

func (t Time) NextWeek() Time {
	return t.AddWeek(1)
}
