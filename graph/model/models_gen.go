// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Author  *User  `json:"author"`
	Post    *Post  `json:"post"`
}

type Meetup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	User        *User  `json:"user"`
}

type Mutation struct {
}

type NewComment struct {
	Content  string `json:"content"`
	AuthorID string `json:"authorId"`
	PostID   string `json:"postId"`
}

type NewMeetup struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewNotification struct {
	Message string `json:"message"`
	UserID  string `json:"userId"`
}

type NewPost struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID string `json:"authorId"`
}

type NewProfile struct {
	Bio    string `json:"bio"`
	UserID string `json:"userId"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Notification struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	User      *User  `json:"user"`
	CreatedAt string `json:"createdAt"`
}

type Post struct {
	ID       string     `json:"id"`
	Title    string     `json:"title"`
	Content  string     `json:"content"`
	Author   *User      `json:"author"`
	Comments []*Comment `json:"comments"`
}

type Profile struct {
	ID   string `json:"id"`
	Bio  string `json:"bio"`
	User *User  `json:"user"`
}

type Query struct {
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Meetups  []*Meetup `json:"meetups"`
}
