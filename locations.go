package gocorona

// Coordinates hols coordinates of a location
type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Location holds data of a location
type Location struct {
	Coordinates Coordinates `json:"coordinates"`
	Country     string      `json:"country"`
	CountryCode string      `json:"country_code"`
	ID          int         `json:"id"`
	Latest      Latest      `json:"latest"`
	Province    string      `json:"province"`
}

// Locations holds response from endpoint /v2/locations
type Locations struct {
	Locations []Location `json:"locations"`
}
