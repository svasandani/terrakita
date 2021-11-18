/******** API Models ********/

struct FilterRequest {
  Lings []int `json:"lings"`
  LingProperties []int `json:"ling_properties"`
  Linglets []int `json:"linglets"`
  LingletProperties []int `json:"linglet_properties"`
}