package role

type Entity struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Rank     int    `json:"rank"`
	MaxCount int    `json:"max_count"`
}
