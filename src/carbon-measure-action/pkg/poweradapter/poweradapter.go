package poweradapter

import (
	"fmt"
	EM "main/pkg/electricitymap"
	"os"
)

//RETURN
type TypReturn struct {
	liveCarbonIntensity int
	history             []struct {
		carbonIntensity int
		datetime        string
	}
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

////////////////////////////////////////////////////////////////////////
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
	///
	electricityMapZoneKey := os.Getenv("ELECTRICITY_MAP_AUTH_TOKEN")
	em := EM.New(electricityMapZoneKey)
	data, _ := em.LiveCarbonIntensity(EM.TypAPIParams{Zone: "US-CAL-CIS"})
	data2, _ := em.RecentCarbonIntensity(EM.TypAPIParams{Zone: "US-CAL-CIS"})
	fmt.Println("Printing ZoneKey --> ", electricityMapZoneKey)
	fmt.Println("Printing --> em....... ", em)

	fmt.Println("Printing --> LiveCarbonIntensity....... ", data)
	fmt.Println("Printing --> RecentCarbonIntensity....... ", data2)

	//Type this in PS
	//$Env:ELECTRICITY_MAP_AUTH_TOKEN="3bhtgXSayVvgmuwEHry6zYYr"

}
