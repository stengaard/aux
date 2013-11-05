package aux

import (
	"testing"
	"time"
)

type tc struct {
	exp string
	dur string
}

var hdTests = []tc{
	{"less than a minute", "28s"},
	{"1 minute", "32s"},
	{"2 minutes", "92s"},
	{"2 minutes", "-92s"},
	{"3 minutes", "-3m"},
	{"5 minutes", "4m59s"},
	{"6 minutes", "6m29s"},
	{"44 minutes", "44m29s"},
	{"about 1 hour", "44m29s40us"},
	{"about 2 hours", "1h30m"},
	{"about 23 hours", "22h30m"},
	{"1 day", "24h"},
	{"1 day", "41h59m29s"},
	{"2 days", "41h59m29s1us"},
	{"30 days", "719h59m29s"}, // 720 hours in a month
	{"about 1 month", "719h59m40s"},
	{"2 months", "720h719h59m40s"},
	{"4 months", "720h720h720h719h59m40s"},
	{"12 months", "720h720h720h720h720h720h720h720h720h720h720h719h59m29s"},
	{"about 1 year", "8760h"},
}

func TestRoughDuration(t *testing.T) {

	for _, test := range hdTests {
		d, err := time.ParseDuration(test.dur)
		if err != nil {
			t.Fatal(err)
			continue
		}

		s := RoughDuration(d)
		t.Logf("Input   : %s", test.dur)
		t.Logf("Expected: %s", test.exp)
		t.Logf("Got     : %s", s)
		if s != test.exp {
			t.Log("FAIL")
			t.Fail()
		} else {
			t.Log("OK")
		}

		t.Log("")

	}

}
