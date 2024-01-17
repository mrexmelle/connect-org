package placement

import (
	"time"

	"github.com/mrexmelle/connect-orgs/internal/config"
	"gorm.io/gorm"
)

type Repository interface {
	Create(req *Entity) (*Entity, error)
	FindById(id string) (*Entity, error)
	FindByOrganizationId(organizationId string) ([]Entity, error)
	UpdateById(fields map[string]interface{}, ehid string) error
	DeleteById(id string) error
}

type RepositoryImpl struct {
	ConfigService *config.Service
	TableName     string
	Query         Query
}

func NewRepository(cfg *config.Service) Repository {
	return &RepositoryImpl{
		ConfigService: cfg,
		TableName:     "placements",
		Query:         NewQuery(cfg.ReadDb, "placements"),
	}
}

func (r *RepositoryImpl) Create(req *Entity) (*Entity, error) {
	result := r.ConfigService.WriteDb.Exec(
		"INSERT INTO "+r.TableName+"(organization_id, role_id, ehid, "+
			"created_at, updated_at) "+
			"VALUES(?, ?, ?, NOW(), NOW())",
		req.OrganizationId,
		req.RoleId,
		req.Ehid,
	)
	if result.Error != nil {
		return nil, result.Error
	}

	return req, nil
}

func (r *RepositoryImpl) FindById(id string) (*Entity, error) {
	org := Entity{
		Id: id,
	}
	result := r.Query.SelectById(FieldsAllExceptId, id).First(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return &org, nil
}

func (r *RepositoryImpl) FindByOrganizationId(organizationId string) ([]Entity, error) {
	orgroles := []Entity{}
	result := r.Query.SelectByOrganizationId(FieldsAll, organizationId).Find(&orgroles)
	if result.Error != nil {
		return nil, result.Error
	}
	return orgroles, nil
}

func (r *RepositoryImpl) UpdateById(
	fields map[string]interface{},
	id string,
) error {
	dbFields := map[string]interface{}{}

	for i := range FieldsPatchable {
		introspectedKey := FieldsPatchable[i]
		value, ok := fields[introspectedKey]
		if ok {
			dbFields[introspectedKey] = value
		}
	}

	if len(dbFields) > 0 {
		dbFields["updated_at"] = time.Now()
		result := r.ConfigService.WriteDb.
			Table(r.TableName).
			Where("id = ?", id).
			Updates(dbFields)

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
	}

	return nil
}

func (r *RepositoryImpl) DeleteById(id string) error {
	now := time.Now()
	result := r.ConfigService.WriteDb.
		Table(r.TableName).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(
			map[string]interface{}{
				"ehid":       "",
				"deleted_at": now,
				"updated_at": now,
			},
		)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
