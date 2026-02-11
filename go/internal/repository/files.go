package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type File struct {
	FileID    int64     `db:"file_id"`
	UserID    int64     `db:"user_id"`
	FileName  string    `db:"file_name"`
	FileValue string    `db:"file_value"`
	Deleted   bool      `db:"deleted"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// 内部にDBの接続プールを保持
type FileRepository struct {
	db *sql.DB
}

// NewFileRepositoryは、FileRepositoryのインスタンスを生成する
// 依存性として、初期化済みの*sql.DBを受け取る
func NewFileRepository(db *sql.DB) *FileRepository {
	return &FileRepository{db: db}
}

// コンストラクタ
func NewFile(FileName string) *File {
	return &File{
		FileName: FileName,
	}
}

// ユーザに紐づくファイル取得
func (ur *FileRepository) SelectFiles(userID int64) ([]File, error) {
	var Files []File

	rows, err := ur.db.Query("select file_id, user_id, file_name, file_value, deleted, created_at, updated_at from files WHERE user_id = $1 AND deleted = false", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query Files: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var File File
		if err := rows.Scan(&File.FileID, &File.UserID, &File.FileName, &File.FileValue, &File.Deleted, &File.CreatedAt, &File.UpdatedAt); err != nil {
			return nil, fmt.Errorf("エラー：Cast files")
		}
		Files = append(Files, File)
	}

	return Files, nil
}
