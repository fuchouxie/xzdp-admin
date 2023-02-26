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

func newTbUser(db *gorm.DB, opts ...gen.DOOption) tbUser {
	_tbUser := tbUser{}

	_tbUser.tbUserDo.UseDB(db, opts...)
	_tbUser.tbUserDo.UseModel(&entity.TbUser{})

	tableName := _tbUser.tbUserDo.TableName()
	_tbUser.ALL = field.NewAsterisk(tableName)
	_tbUser.ID = field.NewInt64(tableName, "id")
	_tbUser.Phone = field.NewString(tableName, "phone")
	_tbUser.Password = field.NewString(tableName, "password")
	_tbUser.NickName = field.NewString(tableName, "nick_name")
	_tbUser.Icon = field.NewString(tableName, "icon")
	_tbUser.CreatedAt = field.NewTime(tableName, "created_at")
	_tbUser.UpdatedAt = field.NewTime(tableName, "updated_at")
	_tbUser.DeletedAt = field.NewField(tableName, "deleted_at")
	_tbUser.UserInfo = tbUserHasOneUserInfo{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("UserInfo", "entity.TbUserInfo"),
	}

	_tbUser.fillFieldMap()

	return _tbUser
}

type tbUser struct {
	tbUserDo tbUserDo

	ALL       field.Asterisk
	ID        field.Int64  // 主键
	Phone     field.String // 手机号码
	Password  field.String // 密码，加密存储
	NickName  field.String // 昵称，默认是用户id
	Icon      field.String // 人物头像
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间
	DeletedAt field.Field
	UserInfo  tbUserHasOneUserInfo

	fieldMap map[string]field.Expr
}

func (t tbUser) Table(newTableName string) *tbUser {
	t.tbUserDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tbUser) As(alias string) *tbUser {
	t.tbUserDo.DO = *(t.tbUserDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tbUser) updateTableName(table string) *tbUser {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.Phone = field.NewString(table, "phone")
	t.Password = field.NewString(table, "password")
	t.NickName = field.NewString(table, "nick_name")
	t.Icon = field.NewString(table, "icon")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *tbUser) WithContext(ctx context.Context) *tbUserDo { return t.tbUserDo.WithContext(ctx) }

func (t tbUser) TableName() string { return t.tbUserDo.TableName() }

func (t tbUser) Alias() string { return t.tbUserDo.Alias() }

func (t *tbUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tbUser) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 9)
	t.fieldMap["id"] = t.ID
	t.fieldMap["phone"] = t.Phone
	t.fieldMap["password"] = t.Password
	t.fieldMap["nick_name"] = t.NickName
	t.fieldMap["icon"] = t.Icon
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt

}

func (t tbUser) clone(db *gorm.DB) tbUser {
	t.tbUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tbUser) replaceDB(db *gorm.DB) tbUser {
	t.tbUserDo.ReplaceDB(db)
	return t
}

type tbUserHasOneUserInfo struct {
	db *gorm.DB

	field.RelationField
}

func (a tbUserHasOneUserInfo) Where(conds ...field.Expr) *tbUserHasOneUserInfo {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a tbUserHasOneUserInfo) WithContext(ctx context.Context) *tbUserHasOneUserInfo {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a tbUserHasOneUserInfo) Model(m *entity.TbUser) *tbUserHasOneUserInfoTx {
	return &tbUserHasOneUserInfoTx{a.db.Model(m).Association(a.Name())}
}

type tbUserHasOneUserInfoTx struct{ tx *gorm.Association }

func (a tbUserHasOneUserInfoTx) Find() (result *entity.TbUserInfo, err error) {
	return result, a.tx.Find(&result)
}

func (a tbUserHasOneUserInfoTx) Append(values ...*entity.TbUserInfo) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a tbUserHasOneUserInfoTx) Replace(values ...*entity.TbUserInfo) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a tbUserHasOneUserInfoTx) Delete(values ...*entity.TbUserInfo) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a tbUserHasOneUserInfoTx) Clear() error {
	return a.tx.Clear()
}

func (a tbUserHasOneUserInfoTx) Count() int64 {
	return a.tx.Count()
}

type tbUserDo struct{ gen.DO }

func (t tbUserDo) Debug() *tbUserDo {
	return t.withDO(t.DO.Debug())
}

func (t tbUserDo) WithContext(ctx context.Context) *tbUserDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tbUserDo) ReadDB() *tbUserDo {
	return t.Clauses(dbresolver.Read)
}

func (t tbUserDo) WriteDB() *tbUserDo {
	return t.Clauses(dbresolver.Write)
}

func (t tbUserDo) Session(config *gorm.Session) *tbUserDo {
	return t.withDO(t.DO.Session(config))
}

func (t tbUserDo) Clauses(conds ...clause.Expression) *tbUserDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tbUserDo) Returning(value interface{}, columns ...string) *tbUserDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tbUserDo) Not(conds ...gen.Condition) *tbUserDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tbUserDo) Or(conds ...gen.Condition) *tbUserDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tbUserDo) Select(conds ...field.Expr) *tbUserDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tbUserDo) Where(conds ...gen.Condition) *tbUserDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tbUserDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tbUserDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tbUserDo) Order(conds ...field.Expr) *tbUserDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tbUserDo) Distinct(cols ...field.Expr) *tbUserDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tbUserDo) Omit(cols ...field.Expr) *tbUserDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tbUserDo) Join(table schema.Tabler, on ...field.Expr) *tbUserDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tbUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tbUserDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tbUserDo) RightJoin(table schema.Tabler, on ...field.Expr) *tbUserDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tbUserDo) Group(cols ...field.Expr) *tbUserDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tbUserDo) Having(conds ...gen.Condition) *tbUserDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tbUserDo) Limit(limit int) *tbUserDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tbUserDo) Offset(offset int) *tbUserDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tbUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tbUserDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tbUserDo) Unscoped() *tbUserDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tbUserDo) Create(values ...*entity.TbUser) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tbUserDo) CreateInBatches(values []*entity.TbUser, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tbUserDo) Save(values ...*entity.TbUser) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tbUserDo) First() (*entity.TbUser, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbUser), nil
	}
}

func (t tbUserDo) Take() (*entity.TbUser, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbUser), nil
	}
}

func (t tbUserDo) Last() (*entity.TbUser, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbUser), nil
	}
}

func (t tbUserDo) Find() ([]*entity.TbUser, error) {
	result, err := t.DO.Find()
	return result.([]*entity.TbUser), err
}

func (t tbUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.TbUser, err error) {
	buf := make([]*entity.TbUser, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tbUserDo) FindInBatches(result *[]*entity.TbUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tbUserDo) Attrs(attrs ...field.AssignExpr) *tbUserDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tbUserDo) Assign(attrs ...field.AssignExpr) *tbUserDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tbUserDo) Joins(fields ...field.RelationField) *tbUserDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tbUserDo) Preload(fields ...field.RelationField) *tbUserDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tbUserDo) FirstOrInit() (*entity.TbUser, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbUser), nil
	}
}

func (t tbUserDo) FirstOrCreate() (*entity.TbUser, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbUser), nil
	}
}

func (t tbUserDo) FindByPage(offset int, limit int) (result []*entity.TbUser, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tbUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tbUserDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tbUserDo) Delete(models ...*entity.TbUser) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tbUserDo) withDO(do gen.Dao) *tbUserDo {
	t.DO = *do.(*gen.DO)
	return t
}
