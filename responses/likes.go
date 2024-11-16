package responses

// AddLike
// DeleteLike

type LikeActionReply struct {
	ErrorInterface
	Response LikeAction `json:"response,omitempty"`
}

type LikeAction struct {
	Likes     int       `json:"likes,omitempty"`
	Reactions Reactions `json:"reactions,omitempty"`
}

type Reactions struct {
	Items        []ReactionItem `json:"items,omitempty"`
	UserReaction int            `json:"user_reaction,omitempty"`
}

type ReactionItem struct {
	ID    int `json:"id,omitempty"`
	Count int `json:"count,omitempty"`
}

// GetLikes

type LikesReply struct {
	ErrorInterface
	Response Likes `json:"response,omitempty"`
}

type Likes struct {
	Users []User `json:"items,omitempty"`
}

type IsLikedReply struct {
	ErrorInterface
	Response IsLiked `json:"response,omitempty"`
}

type IsLiked struct {
	Liked      int `json:"liked,omitempty"`
	Copied     int `json:"copied,omitempty"`
	ReactionID int `json:"reaction_id,omitempty"`
}
