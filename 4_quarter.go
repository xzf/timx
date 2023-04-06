package timx

var gQuarterData = [4][3]int{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
	{10, 11, 12},
}

func (t Time) Quarter() int {
	return int(t.Month()-1)/3 + 1
}

func (t Time) QuarterStart() Time {
	quarter := t.Quarter()
	thisMonth := int(t.Month())
	if thisMonth != gQuarterData[quarter-1][0] {
		t.Time = t.AddMonth(-1 * (int(t.Month()) - gQuarterData[quarter-1][0])).Time
	}
	return t.MonthStart()
}

func (t Time) QuarterEnd() Time {
	quarter := int(t.Month()-1) / 3
	thisMonth := int(t.Month())
	if thisMonth != gQuarterData[quarter][2] {
		t.Time = t.AddMonth(gQuarterData[quarter][2] - int(t.Month())).Time
	}
	return t.MonthStart()
}

func (t Time) AddQuarter(n int) Time {
	if n == 0 {
		return t
	}
	quarter := int(t.Month()-1)/3 + 1
	targetQuarter := quarter + n
	for targetQuarter < 1 {
		targetQuarter += 4
	}
	for targetQuarter > 4 {
		targetQuarter -= 4
	}
	return t.AddMonth(3 * n)
}

func (t Time) LastQuarter() Time {
	return t.AddQuarter(-1)
}

func (t Time) NextQuarter() Time {
	return t.AddQuarter(1)
}
