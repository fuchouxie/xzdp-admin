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

func newTbUserInfo(db *gorm.DB, opts ...gen.DOOption) tbUserInfo {
	_tbUserInfo := tbUserInfo{}

	_tbUserInfo.tbUserInfoDo.UseDB(db, opts...)
	_tbUserInfo.tbUserInfoDo.UseModel(&entity.TbUserInfo{})

	tableName := _tbUserInfo.tbUserInfoDo.TableName()
	_tbUserInfo.ALL = field.NewAsterisk(tableName)
	_tbUserInfo.UserID = field.NewInt64(tableName, "user_id")
	_tbUserInfo.City = field.NewString(tableName, "city")
	_tbUserInfo.Introduce = field.NewString(tableName, "introduce")
	_tbUserInfo.Fans = field.NewInt32(tableName, "fans")
	_tbUserInfo.Followee = field.NewInt32(tableName, "followee")
	_tbUserInfo.Gender = field.NewBool(tableName, "gender")
	_tbUserInfo.Birthday = field.NewTime(tableName, "birthday")
	_tbUserInfo.Credits = field.NewInt32(tableName, "credits")
	_tbUserInfo.Level = field.NewBool(tableName, "level")
	_tbUserInfo.CreateTime = field.NewTime(tableName, "create_time")
	_tbUserInfo.UpdateTime = field.NewTime(tableName, "update_time")

	_tbUserInfo.fillFieldMap()

	return _tbUserInfo
}

type tbUserInfo struct {
	tbUserInfoDo tbUserInfoDo

	ALL        field.Asterisk
	UserID     field.Int64  // 主键，用户id
	City       field.String // 城市名称
	Introduce  field.String // 个人介绍，不要超过128个字符
	Fans       field.Int32  // 粉丝数量
	Followee   field.Int32  // 关注的人的数量
	Gender     field.Bool   // 性别，0：男，1：女
	Birthday   field.Time   // 生日
	Credits    field.Int32  // 积分
	Level      field.Bool   // 会员级别，0~9级,0代表未开通会员
	CreateTime field.Time   // 创建时间
	UpdateTime field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (t tbUserInfo) Table(newTableName string) *tbUserInfo {
	t.tbUserInfoDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tbUserInfo) As(alias string) *tbUserInfo {
	t.tbUserInfoDo.DO = *(t.tbUserInfoDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tbUserInfo) updateTableName(table string) *tbUserInfo {
	t.ALL = field.NewAsterisk(table)
	t.UserID = field.NewInt64(table, "user_id")
	t.City = field.NewString(table, "city")
	t.Introduce = field.NewString(table, "introduce")
	t.Fans = field.NewInt32(table, "fans")
	t.Followee = field.NewInt32(table, "followee")
	t.Gender = field.NewBool(table, "gender")
	t.Birthday = field.NewTime(table, "birthday")
	t.Credits = field.NewInt32(table, "credits")
	t.Level = field.NewBool(table, "level")
	t.CreateTime = field.NewTime(table, "create_time")
	t.UpdateTime = field.NewTime(table, "update_time")

	t.fillFieldMap()

	return t
}

func (t *tbUserInfo) WithContext(ctx context.Context) *tbUserInfoDo {
	return t.tbUserInfoDo.WithContext(ctx)
}

func (t tbUserInfo) TableName() string { return t.tbUserInfoDo.TableName() }

func (t tbUserInfo) Alias() string { return t.tbUserInfoDo.Alias() }

func (t *tbUserInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tbUserInfo) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 11)
	t.fieldMap["user_id"] = t.UserID
	t.fieldMap["city"] = t.City
	t.fieldMap["introduce"] = t.Introduce
	t.fieldMap["fans"] = t.Fans
	t.fieldMap["followee"] = t.Followee
	t.fieldMap["gender"] = t.Gender
	t.fieldMap["birthday"] = t.Birthday
	t.fieldMap["credits"] = t.Credits
	t.fieldMap["level"] = t.Level
	t.fieldMap["create_time"] = t.CreateTime
	t.fieldMap["update_time"] = t.UpdateTime
}

func (t tbUserInfo) clone(db *gorm.DB) tbUserInfo {
	t.tbUserInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tbUserInfo) replaceDB(db *gorm.DB) tbUserInfo {
	t.tbUserInfoDo.ReplaceDB(db)
	return t
}

type tbUserInfoDo struct{ gen.DO }

func (t tbUserInfoDo) Debug() *tbUserInfoDo {
	return t.withDO(t.DO.Debug())
}

func (t tbUserInfoDo) WithContext(ctx context.Context) *tbUserInfoDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tbUserInfoDo) ReadDB() *tbUserInfoDo {
	return t.Clauses(dbresolver.Read)
}

func (t tbUserInfoDo) WriteDB() *tbUserInfoDo {
	return t.Clauses(dbresolver.Write)
}

func (t tbUserInfoDo) Session(config *gorm.Session) *tbUserInfoDo {
	return t.withDO(t.DO.Session(config))
}

func (t tbUserInfoDo) Clauses(conds ...clause.Expression) *tbUserInfoDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tbUserInfoDo) Returning(value interface{}, columns ...string) *tbUserInfoDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tbUserInfoDo) Not(conds ...gen.Condition) *tbUserInfoDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tbUserInfoDo) Or(conds ...gen.Condition) *tbUserInfoDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tbUserInfoDo) Select(conds ...field.Expr) *tbUserInfoDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tbUserInfoDo) Where(conds ...gen.Condition) *tbUserInfoDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tbUserInfoDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tbUserInfoDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tbUserInfoDo) Order(conds ...field.Expr) *tbUserInfoDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tbUserInfoDo) Distinct(cols ...field.Expr) *tbUserInfoDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tbUserInfoDo) Omit(cols ...field.Expr) *tbUserInfoDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tbUserInfoDo) Join(table schema.Tabler, on ...field.Expr) *tbUserInfoDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tbUserInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tbUserInfoDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tbUserInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) *tbUserInfoDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tbUserInfoDo) Group(cols ...field.Expr) *tbUserInfoDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tbUserInfoDo) Having(conds ...gen.Condition) *tbUserInfoDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tbUserInfoDo) Limit(limit int) *tbUserInfoDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tbUserInfoDo) Offset(offset int) *tbUserInfoDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tbUserInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tbUserInfoDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tbUserInfoDo) Unscoped() *tbUserInfoDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tbUserInfoDo) Create(values ...*entity.TbUserInfo) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tbUserInfoDo) CreateInBatches(values []*entity.TbUserInfo, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: myGrom.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tbUserInfoDo) Save(values ...*entity.TbUserInfo) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tbUserInfoDo) First() (*entity.TbUserInfo, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbUserInfo), nil
	}
}

func (t tbUserInfoDo) Take() (*entity.TbUserInfo, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbUserInfo), nil
	}
}

func (t tbUserInfoDo) Last() (*entity.TbUserInfo, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbUserInfo), nil
	}
}

func (t tbUserInfoDo) Find() ([]*entity.TbUserInfo, error) {
	result, err := t.DO.Find()
	return result.([]*entity.TbUserInfo), err
}

func (t tbUserInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.TbUserInfo, err error) {
	buf := make([]*entity.TbUserInfo, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tbUserInfoDo) FindInBatches(result *[]*entity.TbUserInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tbUserInfoDo) Attrs(attrs ...field.AssignExpr) *tbUserInfoDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tbUserInfoDo) Assign(attrs ...field.AssignExpr) *tbUserInfoDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tbUserInfoDo) Joins(fields ...field.RelationField) *tbUserInfoDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tbUserInfoDo) Preload(fields ...field.RelationField) *tbUserInfoDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tbUserInfoDo) FirstOrInit() (*entity.TbUserInfo, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbUserInfo), nil
	}
}

func (t tbUserInfoDo) FirstOrCreate() (*entity.TbUserInfo, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbUserInfo), nil
	}
}

func (t tbUserInfoDo) FindByPage(offset int, limit int) (result []*entity.TbUserInfo, count int64, err error) {
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

func (t tbUserInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tbUserInfoDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tbUserInfoDo) Delete(models ...*entity.TbUserInfo) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tbUserInfoDo) withDO(do gen.Dao) *tbUserInfoDo {
	t.DO = *do.(*gen.DO)
	return t
}