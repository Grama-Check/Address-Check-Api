package models

type Person struct {
	NIC     string `json:"nic"`
	Address string `json:"address"`
	Name    string `json:"name"`
}
