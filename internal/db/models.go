package db

/******** Generic Database Models ********/

type DatabaseConnection struct {
	Username string
	Password string
	Host string
	Port string
	Database string
}

/******** API Models ********/

/*=== Generic ===*/

type ErrorResponse struct {
	Message string `json:"message"`
	StatusCode int `json:"status_code"`
}

type PropertyValuePair struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}

/*=== Filter ===*/

type FilterRequest struct {
	Lings             []int `json:"lings"`
	LingProperties    []int `json:"ling_properties"`
	Linglets          []int `json:"linglets"`
	LingletProperties []int `json:"linglet_properties"`
}

type FilterResponse struct {
	Lings []FilterResponseLing `json:"lings"`
}

type FilterResponseLing struct {
	PropertyValuePairs []PropertyValuePair     `json:"property_value_pairs"`
	Linglets           []FilterResponseLinglet `json:"linglets,omitempty"`
}

type FilterResponseLinglet struct {
	PropertyValuePairs []PropertyValuePair `json:"property_value_pairs"`
}
