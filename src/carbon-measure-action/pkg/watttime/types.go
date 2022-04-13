package watttime

// Coordinates contains the property longitude and latitude
type Coordinates struct {
	Longitude float32
	Latitude  float32
}

// loginResp represents the Login function response
type loginResp struct {
	Token string `json:"token"`
}

// determineGridRegionResp represents the DetermineGridRegion function response
type determineGridRegionResp struct {
	Abbrev string `json:"abbrev"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
}

// listOfGridRegionsResp represents the ListOfGridRegion function response
type listOfGridRegionsResp struct {
	BA       string `json:"ba"`
	Name     string `json:"name"`
	Accept   bool   `json:"accept"`
	DataType string `json:"datatype"`
}

// realTimeEmissionIndexResp represents the RealTimeEmissionsIndex function response
type realTimeEmissionsIndexResp struct {
	Freq      string `json:"freq"`
	BA        string `json:"ba"`
	Percent   string `json:"percent"`
	Moer      string `json:"moer"`
	PointTime string `json:"point_time"`
}

// gridEmissionsDataResp represents the GridEmissionsData function response
type gridEmissionsDataResp struct {
	BA        string  `json:"ba"`
	DataType  string  `json:"datatype"`
	Frequency int     `json:"frequency"`
	Market    string  `json:"market"`
	PointTime string  `json:"point_time"`
	Value     float32 `json:"value"`
	Version   string  `json:"version"`
}

// emissionForecastResp represents the EmissionForecast function response
type emissionForecastResp struct {
	GeneratedAt string     `json:"generated_at"`
	Forecast    []forecast `json:"forecast"`
}

type forecast struct {
	BA        string  `json:"ba"`
	PointTime string  `json:"point_time"`
	Value     float32 `json:"value"`
	Version   string  `json:"version"`
}
