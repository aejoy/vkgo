package responses

import "github.com/aejoy/vkgo/update"

type AlbumsReply struct {
	ErrorInterface
	Response Albums `json:"response"`
}

type Albums struct {
	Albums []Album `json:"items"`
}

type Album struct {
	ID               int                `json:"id"`
	CoverID          int                `json:"thumb_id"`
	CoverURL         int                `json:"thumb_src"`
	OwnerID          int                `json:"owner_id"`
	Title            string             `json:"title"`
	Description      string             `json:"description"`
	CreatedAt        int                `json:"created,omitempty"`
	UpdatedAt        int                `json:"updated,omitempty"`
	Count            int                `json:"size"`
	Sizes            []update.PhotoSize `json:"sizes"`
	CanUpload        int                `json:"can_upload"`
	CanDelete        bool               `json:"can_delete,omitempty"`
	FeedDisabled     int                `json:"feed_disabled,omitempty"`
	FeedHasPinned    int                `json:"feed_has_pinned,omitempty"`
	CanIncludeToFeed bool               `json:"can_include_to_feed,omitempty"`
	CoverIsLast      int                `json:"thumb_is_last,omitempty"`
}
