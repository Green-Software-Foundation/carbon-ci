package poweradapter

import (
	"fmt"
	EM "main/pkg/electricitymap"
	WT "main/pkg/watttime"
	"strconv"
	"strings"
	time "time"
)

//RETURN
type CarbonIntensity struct {
	LiveCarbonIntensity float64
	History             []RecentCIHistory
}

type RecentCIHistory struct {
	CarbonIntensity float64
	Datetime        string
}

//GET FROM
type TypCarbonQueryParams struct {
	IacProvider           string
	IacLocation           string
	CarbonRateProvider    string
	ElectricityMapZoneKey string
	WattTimeUser          string
	WattTimePass          string
}

func LiveCarbonIntensity(params TypCarbonQueryParams) (ci CarbonIntensity) {
	zone := GetLocation(TypCloudLocationQuery{
		Provider:      params.IacProvider,
		Location:      params.IacLocation,
		Powerprovider: params.CarbonRateProvider,
	})
	if strings.ToLower(params.CarbonRateProvider) == "electricitymap" {

		em := EM.New(params.ElectricityMapZoneKey)

		live, _ := em.LiveCarbonIntensity(EM.TypAPIParams{Zone: zone})
		ci.LiveCarbonIntensity = float64(live.CarbonIntensity)
		recent, _ := em.RecentCarbonIntensity(EM.TypAPIParams{Zone: zone})
		var historyci []RecentCIHistory
		for _, i := range recent.History {
			value := float64(i.CarbonIntensity)
			historyci = append(historyci, RecentCIHistory{value, i.Datetime})
		}
		ci.History = historyci

		return

	} else if strings.ToLower(params.CarbonRateProvider) == "watttime" {

		live, recent := Watttime(TypCarbonQueryParams{WattTimeUser: params.WattTimeUser, WattTimePass: params.WattTimePass}, "CAISO_NORTH")

		ci.LiveCarbonIntensity, _ = strconv.ParseFloat(live.Moer, 64)
		if recent != nil {
			var historyci []RecentCIHistory

			for _, i := range *recent {
				value, _ := strconv.ParseFloat(i.Value, 64)
				historyci = append(historyci, RecentCIHistory{value, i.PointTime})
			}
			ci.History = historyci
		}
		fmt.Print("..................", recent)

	}

	return
}

func GetTimeRange() (starttime, endtime string) {
	t := time.Now()
	st := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	et := st.Add((time.Hour * 23) * -1)
	starttime = st.Format(time.RFC3339)
	endtime = et.Format(time.RFC3339)
	return
}

type RealTimeEmission struct {
	Freq      string
	BA        string
	Percent   string
	Moer      string
	PointTime string
}

func Watttime(params TypCarbonQueryParams, BA string) (*WT.RealTimeEmissionsIndexResp, *[]WT.GridEmissionsDataResp) {

	loginwattt := WT.Login(params.WattTimeUser, params.WattTimePass)
	fmt.Println("login", loginwattt)
	live, _ := WT.RealTimeEmissionsIndex(BA, 0, 0, "")
	starttime, endtime := GetTimeRange()
	recent, _ := WT.GridEmissionsData("CAISO_NORTH", 0, 0, endtime, starttime, "", "")

	fmt.Println("RECENT-----------", recent)
	return live, recent
}
