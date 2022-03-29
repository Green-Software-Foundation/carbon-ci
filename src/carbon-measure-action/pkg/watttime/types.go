package watttime

// HEADERS FOR INIT HTTP REQUEST
type httpRequestType struct {
	Url      string
	Method   string
	Data     map[string]string
	Header   map[string]string
	Query    map[string]string
	Response interface{}
}

// WATTTIME - AUTHENTICATION
type loginResponse struct {
	Token string `json:"token"`
}

// WATTTIME -GRID EMISSIONS INFORMATION
type determineGridRegionResp struct {
	Abbrev string `json:"abbrev"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
}

type listOfGridRegionsResp struct {
	BA       string `json:"ba"`
	Name     string `json:"name"`
	Accept   bool   `json:"accept"`
	DataType string `json:"datatype"`
}

type realTimeEmissionsIndexResp struct {
	Freq      string `json:"freq"`
	BA        string `json:"ba"`
	Percent   string `json:"percent"`
	Moer      string `json:"moer"`
	PointTime string `json:"point_time"`
}

type gridEmissionsDataResp struct {
	BA        string  `json:"ba"`
	DataType  string  `json:"datatype"`
	Frequency int     `json:"frequency"`
	Market    string  `json:"market"`
	PointTime string  `json:"point_time"`
	Value     float32 `json:"value"`
	Version   string  `json:"version"`
}

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
