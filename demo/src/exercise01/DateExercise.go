package main

import (
	"fmt"
	"time"
)

/**
关于Go语言中的日期API调用
*/
func main() {
	//testDate()
	//testDate2()
	testDate3()
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

func testDate2() {
	// Use `time.Now` with `Unix` or `UnixNano` to get
	// elapsed time since the Unix epoch in seconds or
	// nanoseconds, respectively.
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// Note that there is no `UnixMillis`, so to get the
	// milliseconds since epoch you'll need to manually
	// divide from nanoseconds.
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// You can also convert integer seconds or nanoseconds
	// since the epoch into the corresponding `time`.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

func testDate3() {
	p := fmt.Println

	// Here's a basic example of formatting a time
	// according to RFC3339, using the corresponding layout
	// constant.
	t := time.Now()
	_, _ = p(t.Format(time.RFC3339))

	// Time parsing uses the same layout values as `Format`.
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	_, _ = p(t1)

	// `Format` and `Parse` use example-based layouts. Usually
	// you'll use a constant from `time` for these layouts, but
	// you can also supply custom layouts. Layouts must use the
	// reference time `Mon Jan 2 15:04:05 MST 2006` to show the
	// pattern with which to format/parse a given time/string.
	// The example time must be exactly as shown: the year 2006,
	// 15 for the hour, Monday for the day of the week, etc.
	_, _ = p(t.Format("3:04PM"))
	_, _ = p(t.Format("Mon Jan _2 15:04:05 2006"))
	_, _ = p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	_, _ = p(t2)

	// For purely numeric representations you can also
	// use standard string formatting with the extracted
	// components of the time value.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse` will return an error on malformed input
	// explaining the parsing problem.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	_, _ = p(e)
}
