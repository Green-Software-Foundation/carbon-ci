package poweradapter

import (
	"fmt"
	EM "main/pkg/electricitymap"
	"os"
)

//RETURN
//TypReturn OLD
type CarbonIntensity struct {
	liveCarbonIntensity int
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

func LiveCarbonIntensity() {

	electricityMapZoneKey := os.Getenv("ELECTRICITY_MAP_AUTH_TOKEN")
	em := EM.New(electricityMapZoneKey)
	fmt.Println("Printing ZoneKey --> ", electricityMapZoneKey)

	var ci CarbonIntensity
	data1, _ := em.LiveCarbonIntensity(EM.TypAPIParams{Zone: "US-CAL-CISO"})
	ci.liveCarbonIntensity = data1.CarbonIntensity

	data2, _ := em.RecentCarbonIntensity(EM.TypAPIParams{Zone: "US-CAL-CISO"})
	var historyci []RecentCIHistory

	for _, i := range data2.History {
		historyci = append(historyci, RecentCIHistory{i.CarbonIntensity, i.Datetime})
	}
	ci.History = historyci
	//historyci.carbonIntensity = append(data2.History, )

	//ci.history = append(ci.history, data2.History)

	fmt.Println("Printing --> em....... ", em)
	fmt.Println("Printing --> LiveCarbonIntensity....... ", ci.liveCarbonIntensity)
	fmt.Println("Printing --> RecentCarbonIntensity....... ", ci.History)

	return
	//Type this in PS
	//$Env:ELECTRICITY_MAP_AUTH_TOKEN="3bhtgXSayVvgmuwEHry6zYYr"
}
