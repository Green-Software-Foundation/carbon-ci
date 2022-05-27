package electricitymap

import (
	"os"
	"testing"
)

func getEnvZoneKey() string {
	return os.Getenv("ELECTRICITY_MAP_AUTH_TOKEN")
}

func TestNew(t *testing.T) {
	em := New(getEnvZoneKey())

	if em.zoneKey != getEnvZoneKey() {
		t.Errorf("ZONE KEY :  %s - EXPECTED: %s", em.zoneKey, getEnvZoneKey())
	}
}

func TestGetZone(t *testing.T) {
	em := New(getEnvZoneKey())

	data, err := em.GetZones()

	if err != nil {
		t.Errorf(err.Error())
	} else {
		t.Log(data)
	}
}

var zoneTests = []struct {
	zone string
}{
	{"US-CAL-CISO"},
	{"US-MIDA-PJM"},
	{"US-MIDW-MISO"},
}

func TestLiveCarbonIntensity(t *testing.T) {
	em := New(getEnvZoneKey())

	for _, test := range zoneTests {
		data, err := em.LiveCarbonIntensity(TypAPIParams{
			Zone: test.zone,
		})

		if err != nil {
			t.Errorf(err.Error())
		} else {
			t.Log(data)
		}
	}
}

func TestLivePowerBreakdown(t *testing.T) {
	em := New(getEnvZoneKey())

	for _, test := range zoneTests {
		data, err := em.LivePowerBreakdown(TypAPIParams{
			Zone: test.zone,
		})

		if err != nil {
			t.Errorf(err.Error())
		} else {
			t.Log(data)
		}
	}
}

func TestRecentCarbonIntensity(t *testing.T) {
	em := New(getEnvZoneKey())

	for _, test := range zoneTests {
		data, err := em.RecentCarbonIntensity(TypAPIParams{
			Zone: test.zone,
		})

		if err != nil {
			t.Errorf(err.Error())
		} else {
			t.Log(data)
		}
	}
}

func TestRecentPowerBreakdown(t *testing.T) {
	em := New(getEnvZoneKey())

	for _, test := range zoneTests {
		data, err := em.RecentPowerBreakdown(TypAPIParams{
			Zone: test.zone,
		})

		if err != nil {
			t.Errorf(err.Error())
		} else {
			t.Log(data)
		}
	}
}

var zoneDatetimeTests = []struct {
	zone     string
	datetime string
}{
	{"US-CAL-CISO", "2022-04-10T00:00:00.000Z"},
	{"US-MIDA-PJM", "2022-04-10T00:00:00.000Z"},
	{"US-MIDW-MISO", "2022-04-10T00:00:00.000Z"},
}

func TestPastCarbonIntensity(t *testing.T) {
	em := New(getEnvZoneKey())

	for _, test := range zoneDatetimeTests {
		data, err := em.PastCarbonIntensity(TypAPIParams{
			Zone:     test.zone,
			Datetime: test.datetime,
		})

		if err != nil {
			t.Errorf(err.Error())
		} else {
			t.Log(data)
		}
	}
}

func TestPastPowerBreakdown(t *testing.T) {
	em := New(getEnvZoneKey())

	for _, test := range zoneDatetimeTests {
		data, err := em.PastPowerBreakdown(TypAPIParams{
			Zone:     test.zone,
			Datetime: test.datetime,
		})

		if err != nil {
			t.Errorf(err.Error())
		} else {
			t.Log(data)
		}
	}
}

var zoneRangeTests = []struct {
	zone  string
	start string
	end   string
}{
	{"US-CAL-CISO", "2022-04-10T00:00:00.000Z", "2022-04-13T00:00:00.000Z"},
	{"US-MIDA-PJM", "2022-04-10T00:00:00.000Z", "2022-04-13T00:00:00.000Z"},
	{"US-MIDW-MISO", "2022-04-10T00:00:00.000Z", "2022-04-13T00:00:00.000Z"},
}

func TestPastPowerBreakdownRange(t *testing.T) {
	em := New(getEnvZoneKey())

	for _, test := range zoneRangeTests {
		data, err := em.PastPowerBreakdownRange(TypAPIParams{
			Zone:  test.zone,
			Start: test.start,
			End:   test.end,
		})

		if err != nil {
			t.Errorf(err.Error())
		} else {
			t.Log(data)
		}
	}
}
