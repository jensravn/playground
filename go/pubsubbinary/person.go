package main

type Person struct {
	ID     int
	Name   string
	Email  string
	Phones []*PhoneNumber
}
type PhoneNumber struct {
	Number string
	Type   PhoneType
}

type PhoneType string

const (
	PhoneTypeMOBILE PhoneType = "MOBILE"
	PhoneTypeHOME   PhoneType = "HOME"
	PhoneTypeWORK   PhoneType = "WORK"
)
