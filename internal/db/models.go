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

type NameValuePair struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

/*=== Filter ===*/

// Filter Lings

type FilterLingsRequest struct {
	Group                   int   `json:"group"`
	Lings                   []int `json:"lings"`
	LingProperties          []int `json:"ling_properties"`
	LingPropertiesInclusive bool  `json:"ling_properties_inclusive"`
}

type FilterLingsResponse struct {
	Type  string                    `json:"type"`
	On    string                    `json:"on"`
	Lings []FilterLingsResponseLing `json:"lings"`
}

type FilterLingsResponseLing struct {
	Id                 string          `json:"id"`
	Name               string          `json:"name"`
	PropertyValuePairs []NameValuePair `json:"property_value_pairs"`
}

// Filter Ling Properties

type FilterLingPropertiesRequest struct {
	Group          int   `json:"group"`
	LingProperties []int `json:"ling_properties"`
	Lings          []int `json:"lings"`
	LingsInclusive bool  `json:"lings_inclusive"`
}

type FilterLingPropertiesResponse struct {
	Type       string                                 `json:"type"`
	On         string                                 `json:"on"`
	Properties []FilterLingPropertiesResponseProperty `json:"properties"`
}

type FilterLingPropertiesResponseProperty struct {
	Id             string          `json:"id"`
	Name           string          `json:"name"`
	LingValuePairs []NameValuePair `json:"ling_value_pairs"`
}

// Filter Linglets

type FilterLingletsRequest struct {
	Group                      int   `json:"group"`
	Linglets                   []int `json:"linglets"`
	LingletProperties          []int `json:"linglet_properties"`
	LingletPropertiesInclusive bool  `json:"linglet_properties_inclusive"`
}

type FilterLingletsResponse struct {
	Type  string                    `json:"type"`
	On    string                    `json:"on"`
	Lings []FilterLingsResponseLing `json:"lings"`
}

type FilterLingletsResponseLing struct {
	Id       string                          `json:"id"`
	Name     string                          `json:"name"`
	Linglets []FilterLingletsResponseLinglet `json:"linglets,omitempty"`
}

type FilterLingletsResponseLinglet struct {
	Id                 string          `json:"id"`
	Name               string          `json:"name"`
	PropertyValuePairs []NameValuePair `json:"property_value_pairs"`
}

// Filter Linglet Properties

type FilterLingletPropertiesRequest struct {
	Group             int   `json:"group"`
	LingletProperties []int `json:"linglet_properties"`
	Linglets          []int `json:"linglets"`
	LingletsInclusive bool  `json:"linglets_inclusive"`
}

type FilterLingletPropertiesResponse struct {
	Type       string                                 `json:"type"`
	On         string                                 `json:"on"`
	Properties []FilterLingPropertiesResponseProperty `json:"properties"`
}

type FilterLingletPropertiesResponseProperty struct {
	Id                string          `json:"id"`
	Name              string          `json:"name"`
	LingletValuePairs []NameValuePair `json:"linglet_value_pairs"`
}

/******** Database Models ********/

type Ling struct {
	Id   string
	Name string
}

type Linglet struct {
	Id   string
	Name string
}

type Property struct {
	Id    string
	Name  string
	Value string
}
