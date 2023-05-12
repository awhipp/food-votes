package structs

type Event struct {
	RawPath string `json:"rawPath"`
	Headers struct {
		XForwardedFor string `json:"x-forwarded-for"`
		Origin        string `json:"origin"`
	} `json:"headers"`
	QueryStringParameters struct {
		Query string `json:"query"`
	} `json:"queryStringParameters"`
}

type Response struct {
	Results []struct {
		FSQID      string `json:"fsq_id"`
		Categories []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Icon struct {
				Prefix string `json:"prefix"`
				Suffix string `json:"suffix"`
			} `json:"icon"`
		} `json:"categories"`
		Chains []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"chains"`
		Distance float64 `json:"distance"`
		GeoCodes struct {
			Main struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"main"`
			Roof struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"roof"`
		} `json:"geocodes"`
		Link     string `json:"link"`
		Location struct {
			Address          string `json:"address"`
			AddressExtended  string `json:"address_extended"`
			CensusBlock      string `json:"census_block"`
			Country          string `json:"country"`
			CrossStreet      string `json:"cross_street"`
			DMA              string `json:"dma"`
			FormattedAddress string `json:"formatted_address"`
			Locality         string `json:"locality"`
			Postcode         string `json:"postcode"`
			Region           string `json:"region"`
		} `json:"location"`
		Name          string `json:"name"`
		RelatedPlaces struct {
			Parent struct {
				FSQID string `json:"fsq_id"`
				Name  string `json:"name"`
			} `json:"parent"`
		} `json:"related_places"`
		Timezone string `json:"timezone"`
	} `json:"results"`
}

type Room struct {
	ID      string   `json:"id"`
	Options []Option `json:"options"`
}

type Option struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Votes   int    `json:"votes"`
}
