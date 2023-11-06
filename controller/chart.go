package controller

func CreateGauge() (Gauge) {
	gauge := Gauge{
		VisType:    "gauge",
		Dimensions: []string{},
		Measurements: []string{
			"% of Patients Admitted to Hospital Emergency Department",
		},
		MeasureAxis: struct {
			Min int `json:"min"`
			Max int `json:"max"`
		}{
			Min: 0,
			Max: 7935345,
		},
		SegmentInfo: struct {
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
		}{
			Limits: []struct {
				Value    int  `json:"value"`
				Gradient bool `json:"gradient"`
			}{
				{
					Value:    2645115,
					Gradient: false,
				},
				{
					Value:    5290230,
					Gradient: false,
				},
			},
			PaletteColors: []struct {
				Color string `json:"color"`
				Index int    `json:"index"`
			}{
				{
					Color: "#83af9b",
					Index: 3,
				},
				{
					Color: "#10cfc9",
					Index: 5,
				},
				{
					Color: "#006580",
					Index: 6,
				},
			},
			Colors: []struct {
				Color int `json:"color"`
			}{
				{
					Color: 3,
				},
			},
		},
		ExactData: [][]int{
			{
				2440522,
			},
		},
	}
	return gauge
}
