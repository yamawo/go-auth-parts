package gateway

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"auth-parts/entity"
	"auth-parts/usecase/port"
)

type UserRepository struct {
	conn *sql.DB
}

// NewUserRepository はUserRepositoryを返す．
func NewUserRepository(conn *sql.DB) port.UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

// GetUserByID はDBからデータを取得する．
func (u *UserRepository) GetUserByID(ctx context.Context, userID string) (*entity.User, error) {
	conn := u.GetDBConn()
	row := conn.QueryRowContext(ctx, "SELECT * FROM `user` WHERE id=?", userID)
	user := entity.User{}
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User Not Found. UserID = %s", userID)
		}
		log.Println(err)
		return nil, errors.New("Internal Server Error. adapter/gateway/GetUserByID")
	}
	return &user, nil
}

// GetDBConn はconnectionを取得する．
func (u *UserRepository) GetDBConn() *sql.DB {
	return u.conn
}
