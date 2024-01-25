package role

import (
	"time"

	"github.com/mrexmelle/connect-org/internal/config"
	"gorm.io/gorm"
)

type Repository interface {
	Create(req *Entity) (*Entity, error)
	FindById(id string) (*Entity, error)
	UpdateById(fields map[string]interface{}, ehid string) error
	DeleteById(id string) error
	CountById(id string) int64
}

type RepositoryImpl struct {
	ConfigService *config.Service
	TableName     string
	Query         Query
}

func NewRepository(cfg *config.Service) Repository {
	return &RepositoryImpl{
		ConfigService: cfg,
		TableName:     "roles",
		Query:         NewQuery(cfg.ReadDb, "roles"),
	}
}

func (r *RepositoryImpl) Create(req *Entity) (*Entity, error) {
	result := r.ConfigService.WriteDb.Exec(
		"INSERT INTO "+r.TableName+"(id, name, rank, max_count, "+
			"created_at, updated_at) "+
			"VALUES(?, ?, ?, ?, NOW(), NOW())",
		req.Id,
		req.Name,
		req.Rank,
		req.MaxCount,
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
				"deleted_at": now,
				"updated_at": now,
			},
		)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryImpl) CountById(id string) int64 {
	var count int64
	result := r.Query.SelectById(FieldsMaxCount, id).Scan(&count)
	if result.Error != nil {
		return 0
	}
	return count
}
