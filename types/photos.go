package types

type Albums struct {
	OwnerID    int   `json:"owner_id"`
	AlbumsID   []int `json:"albums_id"`
	Offset     int   `json:"offset"`
	Count      int   `json:"count"`
	NeedSystem bool  `json:"need_system"`
	NeedCovers bool  `json:"need_covers"`
	Sizes      bool  `json:"sizes"`
}

type AlbumPhotos struct {
	OwnerID  int    `json:"owner_id"`
	AlbumID  any    `json:"album_id"`
	Type     string `json:"type"`
	PhotoIDs []int  `json:"photo_ids"`
	Rev      bool   `json:"rev"`
	Extended bool   `json:"extended"`
	FeedType string `json:"feed_type"`
	Feed     int    `json:"feed"`
	Sizes    bool   `json:"sizes"`
	Offset   int    `json:"offset"`
	Count    int    `json:"count"`
}
