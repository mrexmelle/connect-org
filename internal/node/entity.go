package node

type Entity struct {
	Id           string `json:"id"`
	Hierarchy    string `json:"hierarchy"`
	Name         string `json:"name"`
	EmailAddress string `json:"email_address"`
}

func (e Entity) GetId() string {
	return e.Id
}

func (e Entity) SetId(id string) {
	(&e).Id = id
}
