package construct

type Post struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserId    int64  `json:"userId"`
	CreatedAt string `json:"createdat"`
}

type PostService interface {
	All() (*[]Post, error)
	Get(id int64) (*Post, error)
	Create(p *Post) error
	Delete(id int64) error
	UserPosts(userId int64) (*[]Post, error)
}
