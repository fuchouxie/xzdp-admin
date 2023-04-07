// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"xzdp-admin/app/gen/dal/entity"
)

func newAdminUser(db *gorm.DB, opts ...gen.DOOption) adminUser {
	_adminUser := adminUser{}

	_adminUser.adminUserDo.UseDB(db, opts...)
	_adminUser.adminUserDo.UseModel(&entity.AdminUser{})

	tableName := _adminUser.adminUserDo.TableName()
	_adminUser.ALL = field.NewAsterisk(tableName)
	_adminUser.ID = field.NewInt32(tableName, "id")
	_adminUser.Username = field.NewString(tableName, "username")
	_adminUser.Password = field.NewString(tableName, "password")
	_adminUser.Phone = field.NewString(tableName, "phone")
	_adminUser.Name = field.NewString(tableName, "name")
	_adminUser.Remark = field.NewString(tableName, "remark")
	_adminUser.RoleID = field.NewInt32(tableName, "role_id")
	_adminUser.CreatedAt = field.NewTime(tableName, "created_at")
	_adminUser.UpdatedAt = field.NewTime(tableName, "updated_at")
	_adminUser.DeletedAt = field.NewField(tableName, "deleted_at")
	_adminUser.Code = field.NewString(tableName, "code")

	_adminUser.fillFieldMap()

	return _adminUser
}

type adminUser struct {
	adminUserDo adminUserDo

	ALL       field.Asterisk
	ID        field.Int32
	Username  field.String
	Password  field.String
	Phone     field.String
	Name      field.String
	Remark    field.String
	RoleID    field.Int32
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Code      field.String

	fieldMap map[string]field.Expr
}

func (a adminUser) Table(newTableName string) *adminUser {
	a.adminUserDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminUser) As(alias string) *adminUser {
	a.adminUserDo.DO = *(a.adminUserDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminUser) updateTableName(table string) *adminUser {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "id")
	a.Username = field.NewString(table, "username")
	a.Password = field.NewString(table, "password")
	a.Phone = field.NewString(table, "phone")
	a.Name = field.NewString(table, "name")
	a.Remark = field.NewString(table, "remark")
	a.RoleID = field.NewInt32(table, "role_id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Code = field.NewString(table, "code")

	a.fillFieldMap()

	return a
}

func (a *adminUser) WithContext(ctx context.Context) *adminUserDo {
	return a.adminUserDo.WithContext(ctx)
}

func (a adminUser) TableName() string { return a.adminUserDo.TableName() }

func (a adminUser) Alias() string { return a.adminUserDo.Alias() }

func (a *adminUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminUser) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 11)
	a.fieldMap["id"] = a.ID
	a.fieldMap["username"] = a.Username
	a.fieldMap["password"] = a.Password
	a.fieldMap["phone"] = a.Phone
	a.fieldMap["name"] = a.Name
	a.fieldMap["remark"] = a.Remark
	a.fieldMap["role_id"] = a.RoleID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["code"] = a.Code
}

func (a adminUser) clone(db *gorm.DB) adminUser {
	a.adminUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminUser) replaceDB(db *gorm.DB) adminUser {
	a.adminUserDo.ReplaceDB(db)
	return a
}

type adminUserDo struct{ gen.DO }

func (a adminUserDo) Debug() *adminUserDo {
	return a.withDO(a.DO.Debug())
}

func (a adminUserDo) WithContext(ctx context.Context) *adminUserDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminUserDo) ReadDB() *adminUserDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminUserDo) WriteDB() *adminUserDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminUserDo) Session(config *gorm.Session) *adminUserDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminUserDo) Clauses(conds ...clause.Expression) *adminUserDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminUserDo) Returning(value interface{}, columns ...string) *adminUserDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminUserDo) Not(conds ...gen.Condition) *adminUserDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminUserDo) Or(conds ...gen.Condition) *adminUserDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminUserDo) Select(conds ...field.Expr) *adminUserDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminUserDo) Where(conds ...gen.Condition) *adminUserDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminUserDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *adminUserDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a adminUserDo) Order(conds ...field.Expr) *adminUserDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminUserDo) Distinct(cols ...field.Expr) *adminUserDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminUserDo) Omit(cols ...field.Expr) *adminUserDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminUserDo) Join(table schema.Tabler, on ...field.Expr) *adminUserDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) *adminUserDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminUserDo) RightJoin(table schema.Tabler, on ...field.Expr) *adminUserDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminUserDo) Group(cols ...field.Expr) *adminUserDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminUserDo) Having(conds ...gen.Condition) *adminUserDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminUserDo) Limit(limit int) *adminUserDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminUserDo) Offset(offset int) *adminUserDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *adminUserDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminUserDo) Unscoped() *adminUserDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminUserDo) Create(values ...*entity.AdminUser) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminUserDo) CreateInBatches(values []*entity.AdminUser, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminUserDo) Save(values ...*entity.AdminUser) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminUserDo) First() (*entity.AdminUser, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AdminUser), nil
	}
}

func (a adminUserDo) Take() (*entity.AdminUser, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AdminUser), nil
	}
}

func (a adminUserDo) Last() (*entity.AdminUser, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AdminUser), nil
	}
}

func (a adminUserDo) Find() ([]*entity.AdminUser, error) {
	result, err := a.DO.Find()
	return result.([]*entity.AdminUser), err
}

func (a adminUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.AdminUser, err error) {
	buf := make([]*entity.AdminUser, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminUserDo) FindInBatches(result *[]*entity.AdminUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminUserDo) Attrs(attrs ...field.AssignExpr) *adminUserDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminUserDo) Assign(attrs ...field.AssignExpr) *adminUserDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminUserDo) Joins(fields ...field.RelationField) *adminUserDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminUserDo) Preload(fields ...field.RelationField) *adminUserDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminUserDo) FirstOrInit() (*entity.AdminUser, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AdminUser), nil
	}
}

func (a adminUserDo) FirstOrCreate() (*entity.AdminUser, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AdminUser), nil
	}
}

func (a adminUserDo) FindByPage(offset int, limit int) (result []*entity.AdminUser, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminUserDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminUserDo) Delete(models ...*entity.AdminUser) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminUserDo) withDO(do gen.Dao) *adminUserDo {
	a.DO = *do.(*gen.DO)
	return a
}
