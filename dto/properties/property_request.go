package propertiesdto

import (
	"gorm.io/datatypes"
)

type PropertyRequest struct {
	Name_Property string  `json:"name_property"`
	City          string  `json:"city" `
	Address       string  `json:"address"`
	Price         float64 `json:"price"`
	TypeRent      string  `json:"type_rent" `
	// Amenities     string `json:"amenities"`
	Amenities   datatypes.JSON `json:"amenities" `
	Bedroom     int            `json:"bedroom"`
	Bathroom    int            `json:"bathroom"`
	Sqf         string         `json:"sqf"`
	Description string         `json:"description"`
	Image       string         `json:"image"  form:"image"`
	// User_Id       int    `json:"user_id"`
}
