package organization

type Entity struct {
	Id                  string `json:"id"`
	Hierarchy           string `json:"hierarchy"`
	Name                string `json:"name"`
	EmailAddress        string `json:"emailAddress"`
	PrivateSlackChannel string `json:"privateSlackChannel"`
	PublicSlackChannel  string `json:"publicSlackChannel"`
}
