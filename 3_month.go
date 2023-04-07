package timx

func (t Time) MonthStart() Time {
	t.Time = t.AddDay(t.Day()*-1 + 1).Time
	return t.DayStart()
}

func (t Time) MonthEnd() Time {
	n := 28 - t.Day()
	if n > 0 {
		t.Time = t.AddDay(n).Time
	}
	oldMonth := t.Month()
	for oldMonth == t.Month() {
		t.Time = t.NextDay().Time
	}
	return t.LastDay().DayStart()
}

func (t Time) AddMonth(n int) Time {
	if n == 0 {
		return t
	}
	if t.Day() <= 28 {
		t.Time = t.AddDate(0, n, 0)
		return t
	}
	targetMonth := int(t.Month()) + n
	t.Time = t.AddDate(0, n, 0)
	for targetMonth < 1 {
		targetMonth += 12
	}
	for targetMonth > 12 {
		targetMonth -= 12
	}
	for int(t.Month()) != targetMonth {
		t.Time = t.LastDay().Time
	}
	return t
}

func (t Time) LastMonth() Time {
	return t.AddMonth(-1)
}

func (t Time) NextMonth() Time {
	return t.AddMonth(1)
}

func (t Time) MonthFormat() string {
	return t.Time.Format(YYYY_MM)
}
