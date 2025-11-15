package dbdriver

import (
	"fmt"
	"time"
)

type User struct {
	UserID    int64     `db:"user_id"`
	UserName  string    `db:"user_name"`
	Password  string    `db:"password"`
	Deleted   bool      `db:"deleted"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// コンストラクタ
func NewUser(userId int64, userName string, password string, deleted bool) *User {
	return &User{
		UserID:   userId,
		UserName: userName,
		Password: password,
		Deleted:  deleted,
	}
}

// 全ユーザ取得
func SelectUsers() ([]User, error) {
	var users []User
	db, err := NewConnectDB()
	if err != nil {
		return nil, fmt.Errorf("エラー：接続エラー（%v）", err)
	}
	// 最後に接続を閉じる
	defer db.Close()

	rows, err := db.Query("select user_id, user_name, password, deleted, created_at, updated_at from users")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.UserName, &user.Password, &user.Deleted, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("エラー：Cast users")
		}
		users = append(users, user)
	}

	return users, nil
}

// ユーザ名・パスワード指定でユーザ取得
func SelectUsersWhereUsernamePassword(user_name string, password string) ([]User, error) {
	var users []User
	db, err := NewConnectDB()
	if err != nil {
		return nil, fmt.Errorf("エラー：接続エラー（%v）", err)
	}
	// 最後に接続を閉じる
	defer db.Close()

	rows, err := db.Query("select user_id, user_name, password, deleted, created_at, updated_at from users where user_name = $1 and password = $2 and deleted = 0", user_name, password)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.UserName, &user.Password, &user.Deleted, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("エラー：Cast users")
		}
		users = append(users, user)
	}

	return users, nil
}
