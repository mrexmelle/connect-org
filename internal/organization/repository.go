package organization

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/mrexmelle/connect-orgs/internal/config"
	"gorm.io/gorm"
)

type Repository interface {
	Create(req *Entity) (*Entity, error)
	FindById(id string) (*Entity, error)
	DeleteById(id string) error
	FindSiblingsAndAncestralSiblingsByHierarchy(hierarchy string) ([]Entity, error)
	FindChildrenByHierarchy(hierarchy string) ([]Entity, error)
	FindChildrenById(id string) ([]Entity, error)
	FindLineageByHierarchy(hierarchy string) ([]Entity, error)
}

type RepositoryImpl struct {
	ConfigService *config.Service
	TableName     string
}

func NewRepository(cfg *config.Service) Repository {
	return &RepositoryImpl{
		ConfigService: cfg,
		TableName:     "organizations",
	}
}

func (r *RepositoryImpl) Create(req *Entity) (*Entity, error) {
	result := r.ConfigService.WriteDb.Exec(
		"INSERT INTO "+r.TableName+"(id, hierarchy, name, email_address, "+
			"private_slack_channel, public_slack_channel, "+
			"created_at, updated_at) "+
			"VALUES(?, ?, ?, ?, ?, ?, NOW(), NOW())",
		req.Id,
		req.Hierarchy,
		req.Name,
		req.EmailAddress,
		req.PrivateSlackChannel,
		req.PublicSlackChannel,
	)
	if result.Error != nil {
		return nil, result.Error
	}

	return req, nil
}

func (r *RepositoryImpl) FindById(id string) (*Entity, error) {
	entity := &Entity{
		Id: id,
	}
	err := r.ConfigService.ReadDb.
		Select("hierarchy, name, email_address, private_slack_channel, public_slack_channel").
		Table(r.TableName).
		Where("id = ?", id).
		Row().
		Scan(
			&entity.Hierarchy,
			&entity.Name,
			&entity.EmailAddress,
			&entity.PrivateSlackChannel,
			&entity.PublicSlackChannel,
		)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *RepositoryImpl) DeleteById(id string) error {
	now := time.Now()
	result := r.ConfigService.WriteDb.
		Table(r.TableName).
		Where("id = ? AND deleted_at IS NULL", id).
		Updates(
			map[string]interface{}{
				"hierarchy":             id,
				"email_address":         "",
				"private_slack_channel": "",
				"public_slack_channel":  "",
				"deleted_at":            now,
				"updated_at":            now,
			},
		)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *RepositoryImpl) FindSiblingsAndAncestralSiblingsByHierarchy(hierarchy string) ([]Entity, error) {
	lineage := strings.Split(hierarchy, ".")
	if len(lineage) == 0 {
		return []Entity{}, errors.New("no hierarchy found")
	} else if len(lineage) == 1 {
		o, err := r.FindById(lineage[0])
		if err != nil {
			return []Entity{}, err
		}
		return []Entity{*o}, nil
	}

	for i := 1; i < len(lineage); i++ {
		lineage[i] = fmt.Sprintf("%s.%s", lineage[i-1], lineage[i])
	}

	whereClause := "hierarchy SIMILAR TO '[A-Z0-9]*' "
	for i := 0; i < len(lineage)-1; i++ {
		whereClause += fmt.Sprintf("OR hierarchy SIMILAR TO '%s.[A-Z0-9]*' ", lineage[i])
	}
	result, err := r.composeQuery(whereClause).Order("hierarchy ASC").Rows()
	if err != nil {
		return []Entity{}, err
	}
	defer result.Close()

	orgs := make([]Entity, 0)
	for result.Next() {
		org := Entity{}
		result.Scan(&org.Id, &org.Hierarchy, &org.Name, &org.EmailAddress,
			&org.PrivateSlackChannel, &org.PublicSlackChannel,
		)
		orgs = append(orgs, org)
	}
	return orgs, nil
}

func (r *RepositoryImpl) FindChildrenById(id string) ([]Entity, error) {
	subQuery := fmt.Sprintf("SELECT hierarchy FROM %s WHERE id = '%s'", r.TableName, id)
	whereClause := fmt.Sprintf("hierarchy SIMILAR TO CONCAT((%s), '.[A-Z0-9]*')", subQuery)
	result, err := r.ConfigService.ReadDb.
		Select("id, hierarchy, name, email_address, private_slack_channel, public_slack_channel").
		Table(r.TableName).
		Where(whereClause).
		Where("deleted_at IS NULL").Order("hierarchy ASC").Rows()
	if err != nil {
		return []Entity{}, err
	}
	defer result.Close()

	orgs := make([]Entity, 0)
	for result.Next() {
		org := Entity{}
		result.Scan(&org.Id, &org.Hierarchy, &org.Name, &org.EmailAddress,
			&org.PrivateSlackChannel, &org.PublicSlackChannel,
		)
		orgs = append(orgs, org)
	}
	return orgs, nil
}

func (r *RepositoryImpl) FindChildrenByHierarchy(hierarchy string) ([]Entity, error) {
	whereClause := fmt.Sprintf("hierarchy SIMILAR TO '%s.[A-Z0-9]*'", hierarchy)
	result, err := r.composeQuery(whereClause).Order("hierarchy ASC").Rows()
	if err != nil {
		return []Entity{}, err
	}
	defer result.Close()

	orgs := make([]Entity, 0)
	for result.Next() {
		org := Entity{}
		result.Scan(&org.Id, &org.Hierarchy, &org.Name, &org.EmailAddress,
			&org.PrivateSlackChannel, &org.PublicSlackChannel,
		)
		orgs = append(orgs, org)
	}
	return orgs, nil
}

func (r *RepositoryImpl) FindLineageByHierarchy(hierarchy string) ([]Entity, error) {
	lineage := strings.Split(hierarchy, ".")
	if len(lineage) == 0 {
		return []Entity{}, errors.New("no hierarchy found")
	} else if len(lineage) == 1 {
		o, err := r.FindById(lineage[0])
		if err != nil {
			return []Entity{}, err
		}
		return []Entity{*o}, nil
	}

	for i := 1; i < len(lineage); i++ {
		lineage[i] = fmt.Sprintf("%s.%s", lineage[i-1], lineage[i])
	}

	whereClause := fmt.Sprintf("hierarchy = '%s' ", lineage[0])
	for i := 1; i < len(lineage); i++ {
		whereClause += fmt.Sprintf("OR hierarchy = '%s' ", lineage[i])
	}
	result, err := r.composeQuery(whereClause).Order("hierarchy ASC").Rows()
	if err != nil {
		return []Entity{}, err
	}
	defer result.Close()

	orgs := make([]Entity, 0)
	for result.Next() {
		org := Entity{}
		result.Scan(&org.Id, &org.Hierarchy, &org.Name, &org.EmailAddress,
			&org.PrivateSlackChannel, &org.PublicSlackChannel,
		)
		orgs = append(orgs, org)
	}
	return orgs, nil
}

func (r *RepositoryImpl) composeQuery(whereClause string) *gorm.DB {
	return r.ConfigService.ReadDb.
		Select("id, hierarchy, name, email_address, private_slack_channel, public_slack_channel").
		Table(r.TableName).
		Where(whereClause).
		Where("deleted_at IS NULL")
}
