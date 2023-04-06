package timx

import (
	"time"
)

const (
	Day  = time.Hour * 24
	Week = Day * 7
	//Month = Day * 30  //not always a standard month, it may be wrong
	//Year  = Day * 365 //not always a standard year, it may be wrong
)

type Time struct {
	time.Time
}

func New(t time.Time) Time {
	return Time{Time: t}
}

/*
time.AddDate have some "bug"
like 2001-03-31 decrease one month will be 2001-03-03 not 2001-02
so provide AddMonth\AddYear
*/
