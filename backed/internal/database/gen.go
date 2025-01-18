// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package database

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q         = new(Query)
	File      *file
	FruitKind *fruitKind
	GoodsCard *goodsCard
	UserInfo  *userInfo
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	File = &Q.File
	FruitKind = &Q.FruitKind
	GoodsCard = &Q.GoodsCard
	UserInfo = &Q.UserInfo
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:        db,
		File:      newFile(db, opts...),
		FruitKind: newFruitKind(db, opts...),
		GoodsCard: newGoodsCard(db, opts...),
		UserInfo:  newUserInfo(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	File      file
	FruitKind fruitKind
	GoodsCard goodsCard
	UserInfo  userInfo
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:        db,
		File:      q.File.clone(db),
		FruitKind: q.FruitKind.clone(db),
		GoodsCard: q.GoodsCard.clone(db),
		UserInfo:  q.UserInfo.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:        db,
		File:      q.File.replaceDB(db),
		FruitKind: q.FruitKind.replaceDB(db),
		GoodsCard: q.GoodsCard.replaceDB(db),
		UserInfo:  q.UserInfo.replaceDB(db),
	}
}

type queryCtx struct {
	File      IFileDo
	FruitKind IFruitKindDo
	GoodsCard IGoodsCardDo
	UserInfo  IUserInfoDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		File:      q.File.WithContext(ctx),
		FruitKind: q.FruitKind.WithContext(ctx),
		GoodsCard: q.GoodsCard.WithContext(ctx),
		UserInfo:  q.UserInfo.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
