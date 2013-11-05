// Package aux is (currently very) small set of reusable auxillary tools.
package aux

import (
	"fmt"
	"math"
	"time"
)

var sec = int64(time.Second)
var min = int64(time.Minute)
var hour = int64(time.Hour)
var day = 24 * hour
var month = 30*day + 12*hour // rough avg
var year = 365 * day

// Takes the integer division of x/y rounded towards nearest integer.
func roundDivision(x, y int64) (r int64) {
	d := float64(x) / float64(y)

	s := 1.0
	if math.Signbit(d) {
		s = -s
	}

	// right answer if diff is small enough
	r = roundToZero(d)
	diff := d - float64(r)

	if diff*s >= float64(0.5) {
		r = roundToInf(d)
	}

	return r
}

func roundToZero(d float64) int64 {
	toZero := math.Floor

	if math.Signbit(d) {
		toZero = math.Ceil
	}

	return int64(toZero(d))

}

func roundToInf(d float64) int64 {
	toInf := math.Ceil

	if math.Signbit(d) {
		toInf = math.Floor
	}

	return int64(toInf(d))
}

// RoughDurationDirection returns the RoughDuration of d decorated with an indication
// of past or present. See RoughDuration for further details.
func RoughDurationDirection(d time.Duration) string {
	desc := RoughDuration(d)
	if d.Nanoseconds() < 0 {
		return fmt.Sprintf("in %s", desc)
	} else {
		return fmt.Sprintf("%s ago", desc)
	}
}

// RoughDuration returns a string that is an estimate of the duration in d
//
// It is shamelessly copied from the Ruby-on-Rails equivalent.
//
//  http://apidock.com/rails/ActionView/Helpers/DateHelper/distance_of_time_in_words
//
// It returns something along the lines of this:
//
//  0 <-> 29 secs                                                             # => less than a minute
//  30 secs <-> 1 min, 29 secs                                                # => 1 minute
//  1 min, 30 secs <-> 44 mins, 29 secs                                       # => [2..44] minutes
//  44 mins, 30 secs <-> 89 mins, 29 secs                                     # => about 1 hour
//  89 mins, 30 secs <-> 23 hrs, 59 mins, 29 secs                             # => about [2..24] hours
//  23 hrs, 59 mins, 30 secs <-> 41 hrs, 59 mins, 29 secs                     # => 1 day
//  41 hrs, 59 mins, 30 secs  <-> 29 days, 23 hrs, 59 mins, 29 secs           # => [2..30] days
//  29 days, 23 hrs, 59 mins, 30 secs <-> 59 days, 23 hrs, 59 mins, 29 secs   # => about 1 month
//  59 days, 23 hrs, 59 mins, 30 secs <-> 1 yr minus 1 sec                    # => [2..12] months
//  1 yr <-> 1 yr, 3 months                                                   # => about 1 year
//  1 yr, 3 months <-> 1 yr, 9 months                                         # => over 1 year
//  1 yr, 9 months <-> 2 yr minus 1 sec                                       # => almost 2 years
//  2 yrs <-> max time or date                                                # => (same rules as 1 yr)
//
// Currenly no attempt is made to handle: leap-seconds, leap-years or dayligt savings time.
func RoughDuration(d time.Duration) string {

	// i guess this could be implemented with a binary search, but it would not
	// be as clear and obvious what is going on.

	dur := d.Nanoseconds()
	s := dur
	if dur < 0 {
		dur = -dur
		s = dur
	}

	if s <= 29*sec {
		return "less than a minute"

	} else if s <= 1*min+29*sec {
		return "1 minute"

	} else if s <= 44*min+29*sec {
		return fmt.Sprintf("%d minutes", roundDivision(dur, min))

	} else if s <= 89*min+29*sec {
		return "about 1 hour"

	} else if s <= 23*hour+59*min+29*sec {
		return fmt.Sprintf("about %d hours", roundDivision(dur, hour))

	} else if s <= 41*hour+59*min+29*sec {
		return "1 day"

	} else if s <= 29*day+23*hour+59*min+29*sec {
		return fmt.Sprintf("%d days", roundDivision(s, day))

	} else if s <= 59*day+23*hour+59*min+29*sec {
		return "about 1 month"

	} else if s <= 1*year-1*sec {
		return fmt.Sprintf("%d months", roundDivision(dur, month))

	} else if s <= 1*year+3*month {
		return "about 1 year"

	} else if s <= 1*year+9*month {
		return "over 1 year"

	} else if s <= 2*year-1*sec {
		return "almost 2 years"
	}
	y := roundDivision(dur, year)
	s = s - y*year

	if s <= 3*month {
		return fmt.Sprintf("about %d years", y)

	} else if s <= 9*month {
		return fmt.Sprintf("over %d years", y)

	} else {
		return fmt.Sprintf("almost %d years", y+1)
	}

}
