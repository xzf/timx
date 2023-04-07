package timx

import "time"

func (t Time) DayStart() Time {
	t.Time = t.Add(-1 * (time.Duration(t.Hour())*time.Hour +
		time.Duration(t.Minute())*time.Minute +
		time.Duration(t.Second())*time.Second +
		time.Duration(t.Nanosecond())*time.Nanosecond))
	return t
}

func (t Time) DayEnd() Time {
	t.Time = t.NextDay().DayStart().Add(time.Nanosecond * -1)
	return t
}

func (t Time) AddDay(n int) Time {
	if n == 0 {
		return t
	}
	t.Time = t.Add(Day * time.Duration(n))
	return t
}

func (t Time) LastDay() Time {
	return t.AddDay(-1)
}

func (t Time) NextDay() Time {
	return t.AddDay(1)
}

func (t Time) DayFormat() string {
	return t.Time.Format(YYYY_MM_DD)
}
