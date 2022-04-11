package poweradapter

import (
	EM "main/pkg/electricitymap"
	"os"
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
	iacProvider           string
	iacLocation           string
	carbonRateProvider    string
	electricityMapZoneKey string
	wattTimeUser          string
	wattTimePass          string
}

////////////////////////////////////////////////////
type TypAPIParams struct {
	Zone    string `json:"zone"`
	History []struct {
		CarbonIntensity int    `json:"carbonIntensity"`
		Datetime        string `json:"datetime"`
		UpdatedAt       string `json:"updatedAt"`
		CreatedAt       string `json:"createdAt"`
	}
}

////////////////////////////////////////////////////

func LiveCarbonIntensity(zoneKey string) CarbonIntensity {
	electricityMapZoneKey := os.Getenv("ELECTRICITY_MAP_AUTH_TOKEN")
	em := EM.New(electricityMapZoneKey)

	var ci CarbonIntensity
	data1, _ := em.LiveCarbonIntensity(EM.TypAPIParams{Zone: zoneKey})
	ci.LiveCarbonIntensity = data1.CarbonIntensity

	data2, _ := em.RecentCarbonIntensity(EM.TypAPIParams{Zone: zoneKey})
	var historyci []RecentCIHistory

	for _, i := range data2.History {
		historyci = append(historyci, RecentCIHistory{i.CarbonIntensity, i.Datetime})
	}
	ci.History = historyci

	return ci
	//$Env:ELECTRICITY_MAP_AUTH_TOKEN="3bhtgXSayVvgmuwEHry6zYYr"
}
