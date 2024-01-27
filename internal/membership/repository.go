package membership

import (
	"time"

	"github.com/mrexmelle/connect-org/internal/config"
	"gorm.io/gorm"
)

type Repository interface {
	Create(req *Entity) (*Entity, error)
	FindById(id int) (*Entity, error)
	UpdateById(fields map[string]interface{}, id int) error
	DeleteById(id int) error
	FindByEhid(ehid string) ([]Entity, error)
	FindCurrentByNodeId(nodeId string) ([]Entity, error)
	FindCurrentByEhid(ehid string) ([]Entity, error)
}

type RepositoryImpl struct {
	ConfigService *config.Service
	TableName     string
	Query         Query
}

func NewRepository(cfg *config.Service) Repository {
	return &RepositoryImpl{
		ConfigService: cfg,
		TableName:     "memberships",
		Query:         NewQuery(cfg.ReadDb, "memberships"),
	}
}

func (r *RepositoryImpl) Create(req *Entity) (*Entity, error) {
	var res *gorm.DB
	if req.EndDate.Valid {
		res = r.ConfigService.WriteDb.Raw(
			"INSERT INTO "+r.TableName+"(ehid, start_date, end_date, node_id, "+
				"created_at, updated_at) "+
				"VALUES(?, ?, ?, ?, NOW(), NOW()) RETURNING id",
			req.Ehid,
			req.StartDate,
			req.EndDate.Time,
			req.NodeId,
		).Scan(&req.Id)
	} else {
		res = r.ConfigService.WriteDb.Raw(
			"INSERT INTO "+r.TableName+"(ehid, start_date, node_id, "+
				"created_at, updated_at) "+
				"VALUES(?, ?, ?, NOW(), NOW()) RETURNING id",
			req.Ehid,
			req.StartDate,
			req.NodeId,
		).Scan(&req.Id)
	}

	if res.Error != nil {
		return nil, res.Error
	}

	return req, nil
}

func (r *RepositoryImpl) FindById(id int) (*Entity, error) {
	response := Entity{
		Id: id,
	}
	result := r.Query.SelectById(FieldsAllExceptId, id).First(&response)
	if result.Error != nil {
		return nil, result.Error
	}
	return &response, nil
}

func (r *RepositoryImpl) FindByEhid(ehid string) ([]Entity, error) {
	response := []Entity{}
	result := r.Query.SelectByEhid(FieldsAll, ehid).Find(&response)
	if result.Error != nil {
		return []Entity{}, result.Error
	}
	return response, nil
}

func (r *RepositoryImpl) FindCurrentByNodeId(nodeId string) ([]Entity, error) {
	response := []Entity{}
	result := r.Query.SelectActiveByNodeId(FieldsAll, nodeId).Find(&response)
	if result.Error != nil {
		return []Entity{}, result.Error
	}
	return response, nil
}

func (r *RepositoryImpl) FindCurrentByEhid(ehid string) ([]Entity, error) {
	response := []Entity{}
	result := r.Query.SelectActiveByEhid(FieldsAll, ehid).Find(&response)
	if result.Error != nil {
		return []Entity{}, result.Error
	}
	return response, nil
}

func (r *RepositoryImpl) UpdateById(fields map[string]interface{}, id int) error {
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

func (r *RepositoryImpl) DeleteById(id int) error {
	result := r.ConfigService.WriteDb.
		Table(r.TableName).
		Delete("id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
