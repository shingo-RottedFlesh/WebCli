package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type clip struct {
	clipID    int64     `db:"clip_id"`
	UserID    int64     `db:"user_id"`
	clipName  string    `db:"clip_name"`
	clipValue string    `db:"clip_value"`
	Deleted   bool      `db:"deleted"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// 内部にDBの接続プールを保持
type clipRepository struct {
	db *sql.DB
}

// NewclipRepositoryは、clipRepositoryのインスタンスを生成する
// 依存性として、初期化済みの*sql.DBを受け取る
func NewclipRepository(db *sql.DB) *clipRepository {
	return &clipRepository{db: db}
}

// コンストラクタ
func Newclip(clipName string) *clip {
	return &clip{
		clipName: clipName,
	}
}

// ユーザに紐づくクリップ取得
func (ur *clipRepository) Selectclips(userID int64) ([]clip, error) {
	var clips []clip

	rows, err := ur.db.Query("select clip_id, user_id, clip_name, clip_value, deleted, created_at, updated_at from clips WHERE user_id = $1 AND deleted = false", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query clips: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var clip clip
		if err := rows.Scan(&clip.clipID, &clip.UserID, &clip.clipName, &clip.clipValue, &clip.Deleted, &clip.CreatedAt, &clip.UpdatedAt); err != nil {
			return nil, fmt.Errorf("エラー：Cast clips")
		}
		clips = append(clips, clip)
	}

	return clips, nil
}
