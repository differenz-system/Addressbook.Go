package model

type Users struct {
	UserID     int    `json:"user_id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ExternalID string `json:"external_id"`
	CreateDate string `json:"create_date"`
}
type Address struct {
	AddressID     int
	Name          string
	Email         string
	ContactNumber string
	IsActive      int
	CreateDate    string
	UserID        int
	IsDeleted     int
}
