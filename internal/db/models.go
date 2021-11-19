package db

/******** Generic Database Models ********/

type DatabaseConnection struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

/******** API Models ********/

/*=== Generic ===*/

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type PropertyValuePair struct {
	Id       string `json:"id"`
	Property string `json:"property"`
	Value    string `json:"value"`
}

/*=== Filter ===*/

type FilterRequest struct {
	Group                      int   `json:"group"`
	Lings                      []int `json:"lings"`
	LingProperties             []int `json:"ling_properties"`
	LingPropertiesInclusive    bool  `json:"ling_properties_inclusive"` // defaults to true
	Linglets                   []int `json:"linglets"`
	LingletProperties          []int `json:"linglet_properties"`
	LingletPropertiesInclusive bool  `json:"linglet_properties_inclusive"` // defaults to true
}

type FilterResponse struct {
	Lings []FilterResponseLing `json:"lings"`
}

type FilterResponseLing struct {
	Id                 string                  `json:"id"`
	Name               string                  `json:"name"`
	PropertyValuePairs []PropertyValuePair     `json:"property_value_pairs"`
	Linglets           []FilterResponseLinglet `json:"linglets,omitempty"`
}

type FilterResponseLinglet struct {
	Id                 string              `json:"id"`
	Name               string              `json:"name"`
	PropertyValuePairs []PropertyValuePair `json:"property_value_pairs"`
}

/******** Database Models ********/

type Ling struct {
	Id         string
	Name       string
	Properties []Property
	Linglets   []Linglet
}

type Linglet struct {
	Id         string
	Name       string
	Properties []Property
}

type Property struct {
	Id    string
	Name  string
	Value string
}
