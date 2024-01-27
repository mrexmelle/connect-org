package membership

import (
	"gorm.io/gorm"
)

var (
	FieldsAll = []string{
		"id",
		"ehid",
		"start_date",
		"end_date",
		"node_id",
	}

	FieldsAllExceptId = []string{
		"ehid",
		"start_date",
		"end_date",
		"node_id",
	}

	FieldsAllExceptIdAndEndDate = []string{
		"ehid",
		"start_date",
		"node_id",
	}

	FieldsPatchable = []string{
		"end_date",
	}
)

type Query interface {
	SelectById(fields []string, id int) *gorm.DB
	SelectByEhid(fields []string, ehid string) *gorm.DB
	SelectActiveByNodeId(fields []string, nodeId string) *gorm.DB
	SelectActiveByEhid(fields []string, ehid string) *gorm.DB
}

type QueryImpl struct {
	Db        *gorm.DB
	TableName string
}

func NewQuery(db *gorm.DB, tableName string) Query {
	return &QueryImpl{
		Db:        db,
		TableName: tableName,
	}
}

func (q *QueryImpl) performSelect(fields []string) *gorm.DB {
	return q.Db.
		Select(fields).
		Table(q.TableName)
}

func (q *QueryImpl) SelectById(fields []string, id int) *gorm.DB {
	return q.performSelect(fields).
		Where("id = ?", id)
}

func (q *QueryImpl) SelectByEhid(fields []string, ehid string) *gorm.DB {
	return q.performSelect(fields).
		Where("ehid = ?", ehid)
}

func (q *QueryImpl) SelectActiveByNodeId(fields []string, nodeId string) *gorm.DB {
	return q.performSelect(fields).
		Where("node_id = ?", nodeId).
		Where("start_date < NOW()").
		Where("end_date IS NULL OR end_date > NOW()")
}

func (q *QueryImpl) SelectActiveByEhid(fields []string, ehid string) *gorm.DB {
	return q.performSelect(fields).
		Where("ehid = ?", ehid).
		Where("start_date < NOW()").
		Where("end_date IS NULL OR end_date > NOW()")
}
