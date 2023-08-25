// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userAuthFieldNames          = builder.RawFieldNames(&UserAuth{})
	userAuthRows                = strings.Join(userAuthFieldNames, ",")
	userAuthRowsExpectAutoSet   = strings.Join(stringx.Remove(userAuthFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userAuthRowsWithPlaceHolder = strings.Join(stringx.Remove(userAuthFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserAuthIdPrefix              = "cache:userAuth:id:"
	cacheUserAuthAuthTypeAuthKeyPrefix = "cache:userAuth:authType:authKey:"
	cacheUserAuthUserIdAuthTypePrefix  = "cache:userAuth:userId:authType:"
)

type (
	userAuthModel interface {
		Insert(ctx context.Context, data *UserAuth) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserAuth, error)
		FindOneByAuthTypeAuthKey(ctx context.Context, authType string, authKey string) (*UserAuth, error)
		FindOneByUserIdAuthType(ctx context.Context, userId int64, authType string) (*UserAuth, error)
		Update(ctx context.Context, data *UserAuth) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserAuthModel struct {
		sqlc.CachedConn
		table string
	}

	UserAuth struct {
		Id         int64     `db:"id"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		DeleteTime time.Time `db:"delete_time"`
		DelState   int64     `db:"del_state"`
		Version    int64     `db:"version"` // 版本号
		UserId     int64     `db:"user_id"`
		AuthKey    string    `db:"auth_key"`  // 平台唯一id
		AuthType   string    `db:"auth_type"` // 平台类型
	}
)

func newUserAuthModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserAuthModel {
	return &defaultUserAuthModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_auth`",
	}
}

func (m *defaultUserAuthModel) withSession(session sqlx.Session) *defaultUserAuthModel {
	return &defaultUserAuthModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`user_auth`",
	}
}

func (m *defaultUserAuthModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	userAuthIdKey := fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, id)
	userAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userAuthAuthTypeAuthKeyKey, userAuthIdKey, userAuthUserIdAuthTypeKey)
	return err
}

func (m *defaultUserAuthModel) FindOne(ctx context.Context, id int64) (*UserAuth, error) {
	userAuthIdKey := fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, id)
	var resp UserAuth
	err := m.QueryRowCtx(ctx, &resp, userAuthIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userAuthRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) FindOneByAuthTypeAuthKey(ctx context.Context, authType string, authKey string) (*UserAuth, error) {
	userAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheUserAuthAuthTypeAuthKeyPrefix, authType, authKey)
	var resp UserAuth
	err := m.QueryRowIndexCtx(ctx, &resp, userAuthAuthTypeAuthKeyKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `auth_type` = ? and `auth_key` = ? limit 1", userAuthRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, authType, authKey); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) FindOneByUserIdAuthType(ctx context.Context, userId int64, authType string) (*UserAuth, error) {
	userAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheUserAuthUserIdAuthTypePrefix, userId, authType)
	var resp UserAuth
	err := m.QueryRowIndexCtx(ctx, &resp, userAuthUserIdAuthTypeKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `auth_type` = ? limit 1", userAuthRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId, authType); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserAuthModel) Insert(ctx context.Context, data *UserAuth) (sql.Result, error) {
	userAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	userAuthIdKey := fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, data.Id)
	userAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, userAuthRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, time.Now(), data.DelState, data.Version, data.UserId, data.AuthKey, data.AuthType)
	}, userAuthAuthTypeAuthKeyKey, userAuthIdKey, userAuthUserIdAuthTypeKey)
	return ret, err
}

func (m *defaultUserAuthModel) Update(ctx context.Context, newData *UserAuth) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	userAuthAuthTypeAuthKeyKey := fmt.Sprintf("%s%v:%v", cacheUserAuthAuthTypeAuthKeyPrefix, data.AuthType, data.AuthKey)
	userAuthIdKey := fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, data.Id)
	userAuthUserIdAuthTypeKey := fmt.Sprintf("%s%v:%v", cacheUserAuthUserIdAuthTypePrefix, data.UserId, data.AuthType)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userAuthRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.Version, newData.UserId, newData.AuthKey, newData.AuthType, newData.Id)
	}, userAuthAuthTypeAuthKeyKey, userAuthIdKey, userAuthUserIdAuthTypeKey)
	return err
}

func (m *defaultUserAuthModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserAuthIdPrefix, primary)
}

func (m *defaultUserAuthModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userAuthRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserAuthModel) tableName() string {
	return m.table
}
