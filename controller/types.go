package controller

type Gauge struct {
	VisType      string   `json:"visType"`
	Dimensions   []string `json:"dimensions"`
	Measurements []string `json:"measurements"`
	MeasureAxis  struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"measureAxis"`
	SegmentInfo struct {
		Limits []struct {
			Value    int  `json:"value"`
			Gradient bool `json:"gradient"`
		} `json:"limits"`
		PaletteColors []struct {
			Color string `json:"color"`
			Index int    `json:"index"`
		} `json:"paletteColors"`
		Colors []struct {
			Color int `json:"color"`
		} `json:"colors"`
	} `json:"segmentInfo"`
	ExactData [][]int `json:"exactData"`
}
