package main

import (
	"fmt"
	"time"
)

/**
关于Go语言中的日期API调用
*/
func main() {
	testDate()
}

func testDate() {
	p := fmt.Println

	now := time.Now()
	_, _ = p(now)

	then := time.Date(2019, 11, 17, 20, 34, 58, 654646464, time.UTC)
	_, _ = p(then)
	_, _ = p(then.Year())
	_, _ = p(then.Month())
	_, _ = p(then.Day())
	_, _ = p(then.Hour())
	_, _ = p(then.Minute())
	_, _ = p(then.Second())
	_, _ = p(then.Nanosecond())
	_, _ = p(then.Location())

	// The Monday-Sunday `Weekday` is also available.
	_, _ = p(then.Weekday())

	// These methods compare two times, testing if the
	// first occurs before, after, or at the same time
	// as the second, respectively.
	_, _ = p(then.Before(now))
	_, _ = p(then.After(now))
	_, _ = p(then.Equal(now))

	// The `Sub` methods returns a `Duration` representing
	// the interval between two times.
	diff := now.Sub(then)
	_, _ = p(diff)

	// We can compute the length of the duration in
	// various units.
	_, _ = p(diff.Hours())
	_, _ = p(diff.Minutes())
	_, _ = p(diff.Seconds())
	_, _ = p(diff.Nanoseconds())

	// You can use `Add` to advance a time by a given
	// duration, or with a `-` to move backwards by a
	// duration.
	_, _ = p(then.Add(diff))
	_, _ = p(then.Add(-diff))
}
