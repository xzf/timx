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

func Now() Time {
	return Time{Time: time.Now()}
}

func Unix(sec int64, nsec int64) Time {
	return Time{Time: time.Unix(sec, nsec)}
}

/*
time.AddDate have some "bug"
like 2001-03-31 decrease one month will be 2001-03-03 not 2001-02
so provide AddMonth\AddYear
*/

const (
	YYYY                = "2006"
	YYYY_MM             = YYYY + "-01"
	YYYYMM              = YYYY + "01"
	YYYY_MM_DD          = YYYY_MM + "-02"
	YYYYMMDD            = YYYYMM + "02"
	YYYY_MM_DD_HH_mm_ss = YYYY_MM_DD + " 15:04:05"
	YYYYMMDDHHmmss      = YYYYMMDD + "150405"
)
