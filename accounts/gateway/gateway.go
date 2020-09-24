package gateway

type AccountGateway struct {
	GetGateway
	CreateGateway
	UpdateGateway
}

func NewGateway(get GetGateway, create CreateGateway, update UpdateGateway) *AccountGateway {
	return &AccountGateway{
		GetGateway:    get,
		CreateGateway: create,
		UpdateGateway: update,
	}
}