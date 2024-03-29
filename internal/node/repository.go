package node

import (
	"time"

	"github.com/mrexmelle/connect-org/internal/config"
	"gorm.io/gorm"
)

type Repository interface {
	Create(req *Entity) (*Entity, error)
	UpdateById(fields map[string]interface{}, id string) error
	FindById(id string) (*Entity, error)
	DeleteById(id string) error
	FindChildrenByHierarchy(hierarchy string) ([]Entity, error)
	FindLineageByHierarchy(hierarchy string) ([]Entity, error)
	FindLineageSiblingsByHierarchy(hierarchy string) ([]Entity, error)
	FindChildrenById(id string) ([]Entity, error)
	FindLineageById(hierarchy string) ([]Entity, error)
	FindLineageSiblingsById(hierarchy string) ([]Entity, error)
}

type RepositoryImpl struct {
	ConfigService *config.Service
	TableName     string
	Query         Query
}

func NewRepository(cfg *config.Service) Repository {
	return &RepositoryImpl{
		ConfigService: cfg,
		TableName:     "nodes",
		Query:         NewQuery(cfg.ReadDb, "nodes"),
	}
}

func (r *RepositoryImpl) Create(req *Entity) (*Entity, error) {
	result := r.ConfigService.WriteDb.Exec(
		"INSERT INTO "+r.TableName+"(id, hierarchy, name, email_address, "+
			"created_at, updated_at) "+
			"VALUES(?, ?, ?, ?, NOW(), NOW())",
		req.Id,
		req.Hierarchy,
		req.Name,
		req.EmailAddress,
	)
	if result.Error != nil {
		return nil, result.Error
	}

	return req, nil
}

func (r *RepositoryImpl) UpdateById(fields map[string]interface{}, id string) error {
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

func (r *RepositoryImpl) FindById(id string) (*Entity, error) {
	org := Entity{
		Id: id,
	}
	result := r.Query.SelectById(FieldsAll, id).First(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return &org, nil
}

func (r *RepositoryImpl) DeleteById(id string) error {
	now := time.Now()
	result := r.ConfigService.WriteDb.
		Table(r.TableName).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(
			map[string]interface{}{
				"hierarchy":     id,
				"email_address": "",
				"deleted_at":    now,
				"updated_at":    now,
			},
		)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryImpl) FindChildrenByHierarchy(hierarchy string) ([]Entity, error) {
	orgs := []Entity{}
	result := r.Query.SelectChildrenByHierarchy(FieldsAll, hierarchy).Find(&orgs)
	if result.Error != nil {
		return nil, result.Error
	}

	return orgs, nil
}

func (r *RepositoryImpl) FindLineageByHierarchy(hierarchy string) ([]Entity, error) {
	orgs := []Entity{}
	query, err := r.Query.SelectLineageByHierarchy(FieldsAll, hierarchy)
	if err != nil {
		return []Entity{}, err
	}

	result := query.Find(&orgs)
	if result.Error != nil {
		return []Entity{}, result.Error
	}

	return orgs, nil
}

func (r *RepositoryImpl) FindLineageSiblingsByHierarchy(
	hierarchy string,
) ([]Entity, error) {
	orgs := []Entity{}
	query, err := r.Query.SelectLineageSiblingsByHierarchy(
		FieldsAll, hierarchy,
	)
	if err != nil {
		return []Entity{}, err
	}

	result := query.Find(&orgs)
	if result.Error != nil {
		return []Entity{}, result.Error
	}

	return orgs, nil
}

func (r *RepositoryImpl) FindChildrenById(id string) ([]Entity, error) {
	orgs := []Entity{}
	result := r.Query.SelectChildrenById(FieldsAll, id).Find(&orgs)
	if result.Error != nil {
		return nil, result.Error
	}

	return orgs, nil
}

func (r *RepositoryImpl) FindLineageById(id string) ([]Entity, error) {
	org := Entity{}
	result := r.Query.SelectById(FieldsHierarchy, id).First(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.FindLineageByHierarchy(org.Hierarchy)
}

func (r *RepositoryImpl) FindLineageSiblingsById(id string) ([]Entity, error) {
	org := Entity{}
	result := r.Query.SelectById(FieldsHierarchy, id).First(&org)
	if result.Error != nil {
		return nil, result.Error
	}
	return r.FindLineageSiblingsByHierarchy(org.Hierarchy)
}
