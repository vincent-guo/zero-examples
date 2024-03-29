package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	userFieldNames          = builderx.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserIdPrefix       = "cache:user:id:"
	cacheUserUsernamePrefix = "cache:user:username:"
)

type (
	UserModel interface {
		Insert(data User) (sql.Result, error)
		FindOne(id int64) (*User, error)
		FindOneByUsername(username string) (*User, error)
		Update(data User) error
		Delete(id int64) error

		FindByName(name string) (*User, error)
		FindAll() ([]*User, error)
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id       int64  `db:"id"`       // id
		Username string `db:"username"` // username
		Password string `db:"password"` // password
	}
)

func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) FindAll() ([]*User, error) {
	var resp []*User
	err := m.QueryRowsNoCache(&resp, "select * from user")
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindByName(name string) (*User, error) {
	var resp User
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, name)
	err := m.QueryRow(&resp, userIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where username = ? limit 1", userRows, m.table)
		return conn.QueryRow(v, query, name)
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

func (m *defaultUserModel) Insert(data User) (sql.Result, error) {
	userUsernameKey := fmt.Sprintf("%s%v", cacheUserUsernamePrefix, data.Username)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userRowsExpectAutoSet)
		return conn.Exec(query, data.Username, data.Password)
	}, userUsernameKey)
	return ret, err
}

func (m *defaultUserModel) FindOne(id int64) (*User, error) {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	var resp User
	err := m.QueryRow(&resp, userIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
		return conn.QueryRow(v, query, id)
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

func (m *defaultUserModel) FindOneByUsername(username string) (*User, error) {
	userUsernameKey := fmt.Sprintf("%s%v", cacheUserUsernamePrefix, username)
	var resp User
	err := m.QueryRowIndex(&resp, userUsernameKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userRows, m.table)
		if err := conn.QueryRow(&resp, query, username); err != nil {
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

func (m *defaultUserModel) Update(data User) error {
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, data.Id)
	userUsernameKey := fmt.Sprintf("%s%v", cacheUserUsernamePrefix, data.Username)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.Exec(query, data.Username, data.Password, data.Id)
	}, userIdKey, userUsernameKey)
	return err
}

func (m *defaultUserModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	userUsernameKey := fmt.Sprintf("%s%v", cacheUserUsernamePrefix, data.Username)
	userIdKey := fmt.Sprintf("%s%v", cacheUserIdPrefix, id)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, userIdKey, userUsernameKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRow(v, query, primary)
}
