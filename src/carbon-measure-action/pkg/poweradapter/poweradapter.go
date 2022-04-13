package poweradapter

import (
	EM "main/pkg/electricitymap"
	WT "main/pkg/watttime"
	"strings"
	time "time"
)

//RETURN
type CarbonIntensity struct {
	LiveCarbonIntensity int
	History             []RecentCIHistory
}

type RecentCIHistory struct {
	CarbonIntensity int
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
		ci.LiveCarbonIntensity = live.CarbonIntensity
		recent, _ := em.RecentCarbonIntensity(EM.TypAPIParams{Zone: zone})
		var historyci []RecentCIHistory
		for _, i := range recent.History {
			historyci = append(historyci, RecentCIHistory{i.CarbonIntensity, i.Datetime})
		}
		ci.History = historyci

		return

	} else if strings.ToLower(params.CarbonRateProvider) == "watttime" {
		Watttime(TypCarbonQueryParams{WattTimeUser: params.WattTimeUser, WattTimePass: params.WattTimePass}, "CAISO_NORTH")
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

// type WTCarbonIntensity struct{
// 	live string
// 	live string
// }

func Watttime(params TypCarbonQueryParams, BA string) (string, string) {

	starttime, endtime := GetTimeRange()
	//wtlogin := WT.Login(params.WattTimeUser, params.WattTimePass)
	live, _ := WT.RealTimeEmissionsIndex(BA, 0, 0, "")
	recent, _ := WT.GridEmissionsData(BA, 0, 0, endtime, starttime, "", "")

	return
}
