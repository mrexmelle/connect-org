package membership

import (
	"database/sql"
	"time"
)

type Entity struct {
	Id        int
	Ehid      string
	StartDate time.Time
	EndDate   sql.NullTime
	NodeId    string
}

type ViewEntity struct {
	Id        int    `json:"id"`
	Ehid      string `json:"ehid"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	NodeId    string `json:"node_id"`
}

func toViewEntity(e *Entity) *ViewEntity {
	ed := ""
	if e.EndDate.Valid {
		ed = e.EndDate.Time.Format("2006-01-02")
	}
	return &ViewEntity{
		Id:        e.Id,
		Ehid:      e.Ehid,
		StartDate: e.StartDate.Format("2006-01-02"),
		EndDate:   ed,
		NodeId:    e.NodeId,
	}
}

func toViewEntitySlice(s []Entity) []ViewEntity {
	viewEntities := []ViewEntity{}
	for _, e := range s {
		viewEntities = append(viewEntities, *toViewEntity(&e))
	}
	return viewEntities
}
