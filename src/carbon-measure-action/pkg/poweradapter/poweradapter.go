package poweradapter

import (
	EM "main/pkg/electricitymap"
	"os"
)

//RETURN
type CarbonIntensity struct {
	liveCarbonIntensity int
	History             []RecentCIHistory
}

type RecentCIHistory struct {
	CarbonIntensity int
	Datetime        string
}


func LiveCarbonIntensity(zoneKey string) CarbonIntensity {
	electricityMapZoneKey := os.Getenv("ELECTRICITY_MAP_AUTH_TOKEN")
	em := EM.New(electricityMapZoneKey)

	var ci CarbonIntensity
	data1, _ := em.LiveCarbonIntensity(EM.TypAPIParams{Zone: zoneKey})
	ci.liveCarbonIntensity = data1.CarbonIntensity

	data2, _ := em.RecentCarbonIntensity(EM.TypAPIParams{Zone: zoneKey})
	var historyci []RecentCIHistory

	for _, i := range data2.History {
		historyci = append(historyci, RecentCIHistory{i.CarbonIntensity, i.Datetime})
	}
	ci.History = historyci

	return ci
	
	//$Env:ELECTRICITY_MAP_AUTH_TOKEN="3bhtgXSayVvgmuwEHry6zYYr"
}
