package entity

type User struct {
	Username              string `db:"username"`
	ChatID                int64  `db:"chat_id"`
	Tel                   string `db:"tel"`
	NotificationIsEnabled bool   `db:"notification_is_enabled"`
}
