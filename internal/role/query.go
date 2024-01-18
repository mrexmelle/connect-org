package role

import (
	"gorm.io/gorm"
)

var (
	FieldsAll = []string{
		"id",
		"name",
		"rank",
		"max_count",
	}

	FieldsAllExceptId = []string{
		"name",
		"rank",
		"max_count",
	}

	FieldsMaxCount = []string{
		"max_count",
	}

	FieldsPatchable = []string{
		"name",
		"rank",
		"max_count",
	}
)

type Query interface {
	SelectById(fields []string, id string) *gorm.DB
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
		Table(q.TableName).
		Where("deleted_at IS NULL")
}

func (q *QueryImpl) SelectById(fields []string, id string) *gorm.DB {
	return q.performSelect(fields).Where("id = ?", id).
		Order("rank ASC")
}
