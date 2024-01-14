package organization

import (
	"fmt"
	"strings"

	"github.com/mrexmelle/connect-orgs/internal/localerror"
	"gorm.io/gorm"
)

var (
	FieldsAll = []string{
		"id",
		"hierarchy",
		"name",
		"email_address",
		"private_slack_channel",
		"public_slack_channel",
	}

	FieldsAllExceptId = []string{
		"hierarchy",
		"name",
		"email_address",
		"private_slack_channel",
		"public_slack_channel",
	}

	FieldsHierarchy = []string{
		"hierarchy",
	}
)

type Query interface {
	SelectById(fields []string, id string) *gorm.DB
	SelectChildrenByHierarchy(fields []string, hierarchy string) *gorm.DB
	SelectLineageByHierarchy(fields []string, hierarchy string) (*gorm.DB, error)
	SelectSiblingsAndAncestralSiblingsByHierarchy(fields []string, hierarchy string) (*gorm.DB, error)
	SelectChildrenById(fields []string, id string) *gorm.DB
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
	return q.performSelect(fields).Where("id = ?", id)
}

func (q *QueryImpl) SelectChildrenByHierarchy(fields []string, hierarchy string) *gorm.DB {
	return q.performSelect(fields).
		Where("hierarchy SIMILAR TO ?", hierarchy+".[A-Z0-9]*").
		Order("hierarchy ASC")
}

func (q *QueryImpl) SelectLineageByHierarchy(fields []string, hierarchy string) (*gorm.DB, error) {
	lineage := strings.Split(hierarchy, ".")
	if len(lineage) == 0 {
		return nil, localerror.ErrBadHierarchy
	}

	query := q.performSelect(fields).Where("hierarchy = ?", lineage[0])
	for i := 1; i < len(lineage); i++ {
		lineage[i] = fmt.Sprintf("%s.%s", lineage[i-1], lineage[i])
		query = query.Or("hierarchy = ?", lineage[i])
	}
	query = query.Order("hierarchy ASC")
	return query, nil
}

func (q *QueryImpl) SelectSiblingsAndAncestralSiblingsByHierarchy(
	fields []string,
	hierarchy string,
) (*gorm.DB, error) {
	lineage := strings.Split(hierarchy, ".")
	if len(lineage) == 0 {
		return nil, localerror.ErrBadHierarchy
	}

	query := q.performSelect(fields).Where("hierarchy = ?", lineage[0])
	for i := 1; i < len(lineage); i++ {
		lineage[i] = fmt.Sprintf("%s.%s", lineage[i-1], lineage[i])
		query = query.Or("hierarchy SIMILAR TO ?", lineage[i-1]+".[A-Z0-9]*")
	}
	query = query.Order("hierarchy ASC")

	return query, nil
}

func (q *QueryImpl) SelectChildrenById(fields []string, id string) *gorm.DB {
	subQuery := q.SelectById(FieldsHierarchy, id)
	query := q.performSelect(fields).
		Where("hierarchy SIMILAR TO CONCAT((?), '.[A-Z0-9]*')", subQuery).
		Order("hierarchy ASC")
	return query
}
