package dbdriver

import (
	"fmt"
	"time"
)

type User struct {
	UserID    int64
	UserName  string
	Password  string
	Deleted   bool
	CreatedAt time.Time
	UpdateAt  time.Time
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

/* Getter */
func (u *User) GetUserID() int64 {
	return u.UserID
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetDeleted() bool {
	return u.Deleted
}

/* Setter */
func (u *User) SetUserID(userId int64) {
	u.UserID = userId
}

func (u *User) SetUserName(userName string) {
	u.UserName = userName
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetDeleted(deleted bool) {
	u.Deleted = deleted
}

func Select() ([]User, error) {
	var users []User
	db, err := NewConnectDB()
	if err != nil {
		return nil, fmt.Errorf("エラー：接続エラー（%v）", err)
	}
	rows, err := db.Query("select user_id, user_name, password, deleted, created_at, updated_at from users")
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.UserName, &user.Password, &user.Deleted, &user.CreatedAt, &user.UpdateAt); err != nil {
			return nil, fmt.Errorf("エラー：Cast users")
		}
		fmt.Printf("UserID：%v,UserName：%v,Password：%v,Deleted：%v\n", user.UserID, user.UserName, user.Password, user.Deleted)
		users = append(users, user)

	}

	return users, nil
}
