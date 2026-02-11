package dbdriver

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID    int64     `db:"user_id"`
	UserName  string    `db:"user_name"`
	Password  string    `db:"password"`
	Deleted   bool      `db:"deleted"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// 内部にDBの接続プールを保持
type UserRepository struct {
	db *sql.DB
}

// NewUserRepositoryは、UserRepositoryのインスタンスを生成する
// 依存性として、初期化済みの*sql.DBを受け取る
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// コンストラクタ
func NewUser(userName string, password string) *User {
	return &User{
		UserName: userName,
		Password: password,
	}
}

// 全ユーザ取得
func (ur *UserRepository) SelectUsers() ([]User, error) {
	var users []User

	rows, err := ur.db.Query("select user_id, user_name, password, deleted, created_at, updated_at from users")
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

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
func (ur *UserRepository) ValidatePassword(userName string, plainTextPassword string) (bool, error) {
	var storedHash string

	err := ur.db.QueryRow(`SELECT password FROM users where user_name = $1 and deleted = false`, userName).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			// ユーザーが存在しない場合、認証失敗 (エラーではない)
			return false, nil
		}
		return false, fmt.Errorf("failed to query users: %w", err)
	}

	// bcryptでハッシュと平文を比較する
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(plainTextPassword))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// パスワード不一致 (エラーではない)
			return false, nil
		}
		// それ以外のエラー (ハッシュ形式が不正など)
		return false, fmt.Errorf("password comparison error: %w", err)
	}

	return true, nil
}

// ユーザ名・パスワード指定でユーザ取得
func (ur *UserRepository) InsertUser(userName string, plainTextPassword string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(plainTextPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	_, err = ur.db.Exec(`INSERT INTO users (user_name, password) VALUES ($1, $2)`, userName, passwordHash)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}
