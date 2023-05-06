package timx

import "time"

func (t Time) YearStart() Time {
	t.Time = t.AddDay(t.YearDay()*-1 + 1).Time
	return t.DayStart()
}

func (t Time) YearEnd() Time {
	n := 365 - t.YearDay()
	if n < 0 {
		n *= -1
	}
	t.Time = t.AddDay(365 - t.YearDay()).Time
	oldYear := t.Year()
	for oldYear == t.Year() {
		t.Time = t.NextDay().Time
	}
	return t.LastDay().DayStart()
}

func (t Time) AddYear(n int) Time {
	if n == 0 {
		return t
	}
	t.Time = t.Time.AddDate(n, 0, 0)
	//this year not leap year
	if t.Year()%4 != 0 {
		return t
	}
	//day not 02-29
	if t.Month() != time.February && t.Day() != 29 {
		return t
	}
	targetYear := t.Year() + n
	//target year is leap year
	if targetYear%4 == 0 {
		return t
	}
	//now day will be xxxx-03-01
	return t.LastDay() //return xxxx-02-28
}

func (t Time) LastYear() Time {
	return t.AddYear(-1)
}

func (t Time) NextYear() Time {
	return t.AddYear(1)
}

func (t Time) YearFormat() string {
	return t.Time.Format(_YYYY)
}
