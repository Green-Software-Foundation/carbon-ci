package poweradapter

import (
	"fmt"
	EM "main/pkg/electricitymap"
	WT "main/pkg/watttime"
	"strings"
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
	fmt.Println("-- Printing Zone >>> ", zone)

	if strings.ToLower(params.CarbonRateProvider) == "electricitymap" {

		em := EM.New(params.ElectricityMapZoneKey)

		data1, _ := em.LiveCarbonIntensity(EM.TypAPIParams{Zone: zone})
		ci.LiveCarbonIntensity = data1.CarbonIntensity

		data2, _ := em.RecentCarbonIntensity(EM.TypAPIParams{Zone: zone})
		var historyci []RecentCIHistory

		for _, i := range data2.History {
			historyci = append(historyci, RecentCIHistory{i.CarbonIntensity, i.Datetime})
		}
		ci.History = historyci

		return
	}

	if strings.ToLower(params.CarbonRateProvider) == "watttime" {
		Watttime(params.WattTimeUser, params.WattTimePass, params.IacLocation)
	}

	return
}

func Watttime(userName string, passWord string, Region string) {

	wtlog := WT.Login(userName, passWord)
	fmt.Println(">>> Logging in WattTime Account", wtlog)

	liveEmissions, _ := WT.RealTimeEmissionsIndex(Region, 0, 0, "")
	fmt.Println("-- Getting Real Time Emissions Index >>>", liveEmissions)

}