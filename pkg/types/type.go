package types

type Activities struct {
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	Participants  int32   `json:"participants"`
	Price         float32 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
	Accessibility float32 `json:"accessibility"`
}

type Response struct {
	Activity      string
	Category      string
	Participants  int32
	Price         float32
	Link          string
	Key           string
	Accessibility float32
}
