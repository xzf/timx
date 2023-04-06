package main

import (
	"fmt"
	"github.com/xzf/timx"
	"time"
)

func main() {
	t := timx.New(time.Now())
	fmt.Println(t.MonthStart())
	fmt.Println(t.MonthEnd())
}
