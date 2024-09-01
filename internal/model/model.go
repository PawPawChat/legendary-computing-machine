package model

type Profile struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`

	Description string `json:"description"`
	Online      bool   `json:"onilne"`
	NumFriends  int32  `json:"num_friends"`

	LastSeen  string `json:"last_seen"`
	CreatedAt string `json:"created_at"`

	Biography Biography `json:"biography"`
}

type Biography struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Birthday   string `json:"birthday"`
}

type Avatar struct {
	ID      int64  `json:"id"`
	OrigURL string `json:"orig_url"`
	AddedAt string `json:"added_at"`
}

type Chat struct {
	ID         int64  `json:"id"`
	Title      string `json:"title"`
	NumMembers int32  `json:"num_members"`
	CreatedAt  string `json:"created_at"`
}

type Message struct {
	ChatID         int64  `json:"chat_id"`
	SenderID       int64  `json:"sender_id"`
	SenderUsername string `json:"sender_username"`
	Body           string `json:"body"`
	SentAt         string `json:"sent_at"`
}
