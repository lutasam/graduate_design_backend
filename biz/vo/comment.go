package vo

type CommentVO struct {
	ID        string `json:"id"`
	UserName  string `json:"user_name"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
