package designation

import (
	"gorm.io/gorm"
)

var (
	FieldsAll = []string{
		"id",
		"node_id",
		"role_id",
		"ehid",
	}

	FieldsAllExceptId = []string{
		"node_id",
		"role_id",
		"ehid",
	}

	FieldsAllExceptNodeid = []string{
		"id",
		"role_id",
		"ehid",
	}

	FieldsPatchable = []string{
		"node_id",
		"role_id",
		"ehid",
	}
)

type Query interface {
	SelectById(fields []string, id string) *gorm.DB
	SelectByNodeId(fields []string, node_id string) *gorm.DB
	SelectByNodeIdAndRoleId(fields []string, nodeId string, roleId string) *gorm.DB
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
	return q.performSelect(fields).
		Where("id = ?", id)
}

func (q *QueryImpl) SelectByNodeId(fields []string, nodeId string) *gorm.DB {
	return q.performSelect(fields).
		Where("node_id = ?", nodeId)
}

func (q *QueryImpl) SelectByNodeIdAndRoleId(fields []string, nodeId string, roleId string) *gorm.DB {
	return q.performSelect(fields).
		Where("node_id = ?", nodeId).
		Where("role_id = ?", roleId)
}
