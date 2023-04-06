package types

type Plant struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	WaterFrequency int    `json:"water_frequency"`
	ImageUrl       string `json:"image_url"`
	LastWatered    int    `json:"last_watered"`
}
