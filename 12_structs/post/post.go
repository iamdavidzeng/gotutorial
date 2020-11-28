package post

type Post struct {
	Id      int    `json:id`
	Title   string `json:title`
	Content string `json:content`
}

type Repo struct {
	Posts []Post
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Add(post Post) {
	r.Posts = append(r.Posts, post)
}

func (r *Repo) GetAll() []Post {
	return r.Posts
}
