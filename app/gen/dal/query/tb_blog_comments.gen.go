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

func newTbBlogComment(db *gorm.DB, opts ...gen.DOOption) tbBlogComment {
	_tbBlogComment := tbBlogComment{}

	_tbBlogComment.tbBlogCommentDo.UseDB(db, opts...)
	_tbBlogComment.tbBlogCommentDo.UseModel(&entity.TbBlogComment{})

	tableName := _tbBlogComment.tbBlogCommentDo.TableName()
	_tbBlogComment.ALL = field.NewAsterisk(tableName)
	_tbBlogComment.ID = field.NewInt64(tableName, "id")
	_tbBlogComment.UserID = field.NewInt64(tableName, "user_id")
	_tbBlogComment.BlogID = field.NewInt64(tableName, "blog_id")
	_tbBlogComment.ParentID = field.NewInt64(tableName, "parent_id")
	_tbBlogComment.AnswerID = field.NewInt64(tableName, "answer_id")
	_tbBlogComment.Content = field.NewString(tableName, "content")
	_tbBlogComment.Liked = field.NewInt32(tableName, "liked")
	_tbBlogComment.Status = field.NewBool(tableName, "status")
	_tbBlogComment.CreateTime = field.NewTime(tableName, "create_time")
	_tbBlogComment.UpdateTime = field.NewTime(tableName, "update_time")

	_tbBlogComment.fillFieldMap()

	return _tbBlogComment
}

type tbBlogComment struct {
	tbBlogCommentDo tbBlogCommentDo

	ALL        field.Asterisk
	ID         field.Int64  // 主键
	UserID     field.Int64  // 用户id
	BlogID     field.Int64  // 探店id
	ParentID   field.Int64  // 关联的1级评论id，如果是一级评论，则值为0
	AnswerID   field.Int64  // 回复的评论id
	Content    field.String // 回复的内容
	Liked      field.Int32  // 点赞数
	Status     field.Bool   // 状态，0：正常，1：被举报，2：禁止查看
	CreateTime field.Time   // 创建时间
	UpdateTime field.Time   // 更新时间

	fieldMap map[string]field.Expr
}

func (t tbBlogComment) Table(newTableName string) *tbBlogComment {
	t.tbBlogCommentDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tbBlogComment) As(alias string) *tbBlogComment {
	t.tbBlogCommentDo.DO = *(t.tbBlogCommentDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tbBlogComment) updateTableName(table string) *tbBlogComment {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.UserID = field.NewInt64(table, "user_id")
	t.BlogID = field.NewInt64(table, "blog_id")
	t.ParentID = field.NewInt64(table, "parent_id")
	t.AnswerID = field.NewInt64(table, "answer_id")
	t.Content = field.NewString(table, "content")
	t.Liked = field.NewInt32(table, "liked")
	t.Status = field.NewBool(table, "status")
	t.CreateTime = field.NewTime(table, "create_time")
	t.UpdateTime = field.NewTime(table, "update_time")

	t.fillFieldMap()

	return t
}

func (t *tbBlogComment) WithContext(ctx context.Context) *tbBlogCommentDo {
	return t.tbBlogCommentDo.WithContext(ctx)
}

func (t tbBlogComment) TableName() string { return t.tbBlogCommentDo.TableName() }

func (t tbBlogComment) Alias() string { return t.tbBlogCommentDo.Alias() }

func (t *tbBlogComment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tbBlogComment) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 10)
	t.fieldMap["id"] = t.ID
	t.fieldMap["user_id"] = t.UserID
	t.fieldMap["blog_id"] = t.BlogID
	t.fieldMap["parent_id"] = t.ParentID
	t.fieldMap["answer_id"] = t.AnswerID
	t.fieldMap["content"] = t.Content
	t.fieldMap["liked"] = t.Liked
	t.fieldMap["status"] = t.Status
	t.fieldMap["create_time"] = t.CreateTime
	t.fieldMap["update_time"] = t.UpdateTime
}

func (t tbBlogComment) clone(db *gorm.DB) tbBlogComment {
	t.tbBlogCommentDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t tbBlogComment) replaceDB(db *gorm.DB) tbBlogComment {
	t.tbBlogCommentDo.ReplaceDB(db)
	return t
}

type tbBlogCommentDo struct{ gen.DO }

func (t tbBlogCommentDo) Debug() *tbBlogCommentDo {
	return t.withDO(t.DO.Debug())
}

func (t tbBlogCommentDo) WithContext(ctx context.Context) *tbBlogCommentDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tbBlogCommentDo) ReadDB() *tbBlogCommentDo {
	return t.Clauses(dbresolver.Read)
}

func (t tbBlogCommentDo) WriteDB() *tbBlogCommentDo {
	return t.Clauses(dbresolver.Write)
}

func (t tbBlogCommentDo) Session(config *gorm.Session) *tbBlogCommentDo {
	return t.withDO(t.DO.Session(config))
}

func (t tbBlogCommentDo) Clauses(conds ...clause.Expression) *tbBlogCommentDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tbBlogCommentDo) Returning(value interface{}, columns ...string) *tbBlogCommentDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tbBlogCommentDo) Not(conds ...gen.Condition) *tbBlogCommentDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tbBlogCommentDo) Or(conds ...gen.Condition) *tbBlogCommentDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tbBlogCommentDo) Select(conds ...field.Expr) *tbBlogCommentDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tbBlogCommentDo) Where(conds ...gen.Condition) *tbBlogCommentDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tbBlogCommentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *tbBlogCommentDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tbBlogCommentDo) Order(conds ...field.Expr) *tbBlogCommentDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tbBlogCommentDo) Distinct(cols ...field.Expr) *tbBlogCommentDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tbBlogCommentDo) Omit(cols ...field.Expr) *tbBlogCommentDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tbBlogCommentDo) Join(table schema.Tabler, on ...field.Expr) *tbBlogCommentDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tbBlogCommentDo) LeftJoin(table schema.Tabler, on ...field.Expr) *tbBlogCommentDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tbBlogCommentDo) RightJoin(table schema.Tabler, on ...field.Expr) *tbBlogCommentDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tbBlogCommentDo) Group(cols ...field.Expr) *tbBlogCommentDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tbBlogCommentDo) Having(conds ...gen.Condition) *tbBlogCommentDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tbBlogCommentDo) Limit(limit int) *tbBlogCommentDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tbBlogCommentDo) Offset(offset int) *tbBlogCommentDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tbBlogCommentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *tbBlogCommentDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tbBlogCommentDo) Unscoped() *tbBlogCommentDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tbBlogCommentDo) Create(values ...*entity.TbBlogComment) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tbBlogCommentDo) CreateInBatches(values []*entity.TbBlogComment, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: myGrom.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tbBlogCommentDo) Save(values ...*entity.TbBlogComment) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tbBlogCommentDo) First() (*entity.TbBlogComment, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbBlogComment), nil
	}
}

func (t tbBlogCommentDo) Take() (*entity.TbBlogComment, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbBlogComment), nil
	}
}

func (t tbBlogCommentDo) Last() (*entity.TbBlogComment, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbBlogComment), nil
	}
}

func (t tbBlogCommentDo) Find() ([]*entity.TbBlogComment, error) {
	result, err := t.DO.Find()
	return result.([]*entity.TbBlogComment), err
}

func (t tbBlogCommentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.TbBlogComment, err error) {
	buf := make([]*entity.TbBlogComment, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tbBlogCommentDo) FindInBatches(result *[]*entity.TbBlogComment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tbBlogCommentDo) Attrs(attrs ...field.AssignExpr) *tbBlogCommentDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tbBlogCommentDo) Assign(attrs ...field.AssignExpr) *tbBlogCommentDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tbBlogCommentDo) Joins(fields ...field.RelationField) *tbBlogCommentDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tbBlogCommentDo) Preload(fields ...field.RelationField) *tbBlogCommentDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tbBlogCommentDo) FirstOrInit() (*entity.TbBlogComment, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbBlogComment), nil
	}
}

func (t tbBlogCommentDo) FirstOrCreate() (*entity.TbBlogComment, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TbBlogComment), nil
	}
}

func (t tbBlogCommentDo) FindByPage(offset int, limit int) (result []*entity.TbBlogComment, count int64, err error) {
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

func (t tbBlogCommentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tbBlogCommentDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tbBlogCommentDo) Delete(models ...*entity.TbBlogComment) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tbBlogCommentDo) withDO(do gen.Dao) *tbBlogCommentDo {
	t.DO = *do.(*gen.DO)
	return t
}
