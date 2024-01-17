package organization

type Entity struct {
	Id                  string `json:"id"`
	Hierarchy           string `json:"hierarchy"`
	Name                string `json:"name"`
	EmailAddress        string `json:"email_address"`
	PrivateSlackChannel string `json:"private_slack_channel"`
	PublicSlackChannel  string `json:"public_slack_channel"`
}

func (e Entity) GetId() string {
	return e.Id
}

func (e Entity) SetId(id string) {
	(&e).Id = id
}
