package update

type AttachmentPhoto Photo

type Photo struct {
	AlbumID       int         `json:"album_id,omitempty"`
	Date          int         `json:"date,omitempty"`
	ID            int         `json:"id,omitempty"`
	OwnerID       int         `json:"owner_id,omitempty"`
	UserID        int         `json:"user_id,omitempty"`
	CanComment    int         `json:"can_comment,omitempty"`
	PostID        int         `json:"post_id,omitempty"`
	Sizes         []PhotoSize `json:"sizes,omitempty"`
	Text          string      `json:"text,omitempty"`
	WebViewToken  string      `json:"web_view_token,omitempty"`
	Likes         Likes       `json:"likes,omitempty"`
	Comments      Comments    `json:"comments,omitempty"`
	Reposts       Reposts     `json:"reposts,omitempty"`
	Tags          Tags        `json:"tags,omitempty"`
	OriginalPhoto PhotoSize   `json:"orig_photo,omitempty"`
	CanRepost     int         `json:"can_repost"`
	Width         int         `json:"width,omitempty"`
	Height        int         `json:"height,omitempty"`
	SquareCrop    string      `json:"square_crop,omitempty"`
	HasTags       bool        `json:"has_tags,omitempty"`
}

type PhotoSize struct {
	Type      string `json:"type,omitempty"`
	URL       string `json:"url,omitempty"`
	SourceURL string `json:"src,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
}

type PhotoComment struct {
	PostComment

	PhotoID      int `json:"photo_id,omitempty"`
	PhotoOwnerID int `json:"photo_owner_id,omitempty"`
}
