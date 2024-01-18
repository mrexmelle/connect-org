package placement

import (
	"gorm.io/gorm"
)

var (
	FieldsAll = []string{
		"id",
		"organization_id",
		"role_id",
		"ehid",
	}

	FieldsAllExceptId = []string{
		"organization_id",
		"role_id",
		"ehid",
	}

	FieldsAllExceptOrganizationid = []string{
		"id",
		"role_id",
		"ehid",
	}

	FieldsPatchable = []string{
		"organization_id",
		"role_id",
		"ehid",
	}
)

type Query interface {
	SelectById(fields []string, id string) *gorm.DB
	SelectByOrganizationId(fields []string, organization_id string) *gorm.DB
	SelectByOrganizationIdAndRoleId(
		fields []string,
		organizationId string,
		roleId string,
	) *gorm.DB
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

func (q *QueryImpl) SelectByOrganizationId(
	fields []string,
	organization_id string,
) *gorm.DB {
	return q.performSelect(fields).
		Where("organization_id = ?", organization_id)
}

func (q *QueryImpl) SelectByOrganizationIdAndRoleId(
	fields []string,
	organizationId string,
	roleId string,
) *gorm.DB {
	return q.performSelect(fields).
		Where("organization_id = ?", organizationId).
		Where("role_id = ?", roleId)
}
