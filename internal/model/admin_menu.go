package model

type AdminMenu struct {
	Id uint32
	Pid uint32
	Title string
	Icon string
	UrlType string
	UrlValue string
	UrlTarget string
	OnlineHide uint8
	Sort uint16
	SystemMenu uint8
	Status uint8
	Params string
}
