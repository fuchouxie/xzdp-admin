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

func newTbBlog(db *gorm.DB, opts ...gen.DOOption) tbBlog {
	_tbBlog := tbBlog{}

	_tbBlog.tbBlogDo.UseDB(db, opts...)
	_tbBlog.tbBlogDo.UseModel(&entity.TbBlog{})

	tableName := _tbBlog.tbBlogDo.TableName()
	_tbBlog.ALL = field.NewAsterisk(tableName)
	_tbBlog.ID = field.NewInt64(tableName, "id")
	_tbBlog.ShopID = field.NewInt64(tableName, "shop_id")
	_tbBlog.UserID = field.NewInt64(tableName, "user_id")
	_tbBlog.Title = field.NewString(tableName, "title")
	_tbBlog.Images = field.NewString(tableName, "images")
	_tbBlog.Content = field.NewString(tableName, "content")
	_tbBlog.Liked = field.NewInt32(tableName, "liked")
	_tbBlog.Comments = field.NewInt32(tableName, "comments")
	_tbBlog.CreateTime = field.NewTime(tableName, "create_time")
	_tbBlog.UpdateTime = field.NewTime(tableName, "update_time")

	_tbBlog.fillFieldMap()

	return _tbBlog
}

type tbBlog struct {
	tbBlogDo tbBlogDo

	ALL        field.Asterisk
	ID         field.Int64  // 主键
	ShopID     field.Int64  // 商户id
	UserID     field.Int64  // 用户id
	Title      field.String // 标题
	Images     field.String // 探店的照片，最多9张，多张以","隔开
	Content    field.String // 探店的文字描述
	Liked      field.Int32  // 点赞数量
	Comments   field.Int32  // 评论数量
	CreateTime field.Time   // 创建时间
	UpdateTime field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (t tbBlog) Table(newTableName string) *tbBlog {
	t.tbBlogDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tbBlog) As(alias string) *tbBlog {
	t.tbBlogDo.DO = *(t.tbBlogDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tbBlog) updateTableName(table string) *tbBlog {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.ShopID = field.NewInt64(table, "shop_id")
	t.UserID = field.NewInt64(table, "user_id")
	t.Title = field.NewString(table, "title")
	t.Images = field.NewString(table, "images")
	t.Content = field.NewString(table, "content")
	t.Liked = field.NewInt32(table, "liked")
	t.Comments = field.NewInt32(table, "comments")
	t.CreateTime = field.NewTime(table, "create_time")
	t.UpdateTime = field.NewTime(table, "update_time")

	t.fillFieldMap()

	return t
}

func (t *tbBlog) WithContext(ctx context.Context) *tbBlogDo { return t.tbBlogDo.WithContext(ctx) }

func (t tbBlog) TableName() string { return t.tbBlogDo.TableName() }

func (t tbBlog) Alias() string { return t.tbBlogDo.Alias() }

func (t *tbBlog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tbBlog) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 10)
	t.fieldMap["id"] = t.ID
	t.fieldMap["shop_id"] = t.ShopID
	t.fieldMap["user_id"] = t.UserID
	t.fieldMap["title"] = t.Title
	t.fieldMap["images"] = t.Images
	t.fieldMap["content"] = t.Content
	t.fieldMap["liked"] = t.Liked
	t.fieldMap["comments"] = t.Comments
	t.fieldMap["create_time"] = t.CreateTime
	t.fieldMap["update_time"] = t.UpdateTime
}

func (t tbBlog) clone(db *gorm.DB) tbBlog {
	t.tbBlogDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tbBlog) replaceDB(db *gorm.DB) tbBlog {
	t.tbBlogDo.ReplaceDB(db)
	return t
}

type tbBlogDo struct{ gen.DO }

func (t tbBlogDo) Debug() *tbBlogDo {
	return t.withDO(t.DO.Debug())
}

func (t tbBlogDo) WithContext(ctx context.Context) *tbBlogDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tbBlogDo) ReadDB() *tbBlogDo {
	return t.Clauses(dbresolver.Read)
}

func (t tbBlogDo) WriteDB() *tbBlogDo {
	return t.Clauses(dbresolver.Write)
}

func (t tbBlogDo) Session(config *gorm.Session) *tbBlogDo {
	return t.withDO(t.DO.Session(config))
}

func (t tbBlogDo) Clauses(conds ...clause.Expression) *tbBlogDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tbBlogDo) Returning(value interface{}, columns ...string) *tbBlogDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tbBlogDo) Not(conds ...gen.Condition) *tbBlogDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tbBlogDo) Or(conds ...gen.Condition) *tbBlogDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tbBlogDo) Select(conds ...field.Expr) *tbBlogDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tbBlogDo) Where(conds ...gen.Condition) *tbBlogDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tbBlogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tbBlogDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tbBlogDo) Order(conds ...field.Expr) *tbBlogDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tbBlogDo) Distinct(cols ...field.Expr) *tbBlogDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tbBlogDo) Omit(cols ...field.Expr) *tbBlogDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tbBlogDo) Join(table schema.Tabler, on ...field.Expr) *tbBlogDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tbBlogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tbBlogDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tbBlogDo) RightJoin(table schema.Tabler, on ...field.Expr) *tbBlogDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tbBlogDo) Group(cols ...field.Expr) *tbBlogDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tbBlogDo) Having(conds ...gen.Condition) *tbBlogDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tbBlogDo) Limit(limit int) *tbBlogDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tbBlogDo) Offset(offset int) *tbBlogDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tbBlogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tbBlogDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tbBlogDo) Unscoped() *tbBlogDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tbBlogDo) Create(values ...*entity.TbBlog) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tbBlogDo) CreateInBatches(values []*entity.TbBlog, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: myGrom.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tbBlogDo) Save(values ...*entity.TbBlog) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tbBlogDo) First() (*entity.TbBlog, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbBlog), nil
	}
}

func (t tbBlogDo) Take() (*entity.TbBlog, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbBlog), nil
	}
}

func (t tbBlogDo) Last() (*entity.TbBlog, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbBlog), nil
	}
}

func (t tbBlogDo) Find() ([]*entity.TbBlog, error) {
	result, err := t.DO.Find()
	return result.([]*entity.TbBlog), err
}

func (t tbBlogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.TbBlog, err error) {
	buf := make([]*entity.TbBlog, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tbBlogDo) FindInBatches(result *[]*entity.TbBlog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tbBlogDo) Attrs(attrs ...field.AssignExpr) *tbBlogDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tbBlogDo) Assign(attrs ...field.AssignExpr) *tbBlogDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tbBlogDo) Joins(fields ...field.RelationField) *tbBlogDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tbBlogDo) Preload(fields ...field.RelationField) *tbBlogDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tbBlogDo) FirstOrInit() (*entity.TbBlog, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbBlog), nil
	}
}

func (t tbBlogDo) FirstOrCreate() (*entity.TbBlog, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbBlog), nil
	}
}

func (t tbBlogDo) FindByPage(offset int, limit int) (result []*entity.TbBlog, count int64, err error) {
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

func (t tbBlogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tbBlogDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tbBlogDo) Delete(models ...*entity.TbBlog) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tbBlogDo) withDO(do gen.Dao) *tbBlogDo {
	t.DO = *do.(*gen.DO)
	return t
}
