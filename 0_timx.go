package timx

import (
	"time"
)

/*
time.AddDate have some "bug"
like 2001-03-31 decrease one month will be 2001-03-03 not 2001-02
so provide AddMonth\AddYear
*/

type Time struct {
	time.Time
}

const (
	Day  = time.Hour * 24
	Week = Day * 7
	//Month = Day * 30  //not always a standard month, it may be wrong
	//Year  = Day * 365 //not always a standard year, it may be wrong
)

const (
	_YYYY = "2006"
	_MM   = "01" //
	_DD   = "02"
	_HH   = "15"
	_mm   = "04"
	_SS   = "05"

	YYYY_MM = "2006-01"

	YYYYMM = "200601"

	YYYY_MM_DD = "2006-01-02"

	YYYYMMDD = "20060102"

	HH_mm = "15:04"
	HHmm  = "1504"

	HH_mm_SS = "15:04:05"
	HHmmSS   = "150405"

	YYYY_MM_DD__HH_mm_SS = "2006-01-02 15:04:05"
	YYYYMMDDHHmmSS       = "20060102150405"

	YYYY_MM_DDTHH_mm_SSZ = "2006-01-02T15:04:05Z"
)

func New(t time.Time) Time {
	return Time{Time: t}
}

func Now() Time {
	return Time{Time: time.Now()}
}

func NewUnix(sec int64, nsec int64) Time {
	return Time{Time: time.Unix(sec, nsec)}
}

func NewUnixSec(sec int64) Time {
	return Time{Time: time.Unix(sec, 0)}
}

func (t Time) TimeFormat() string {
	return t.Time.Format(YYYY_MM_DD__HH_mm_SS)
}
