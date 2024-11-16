package responses

type StickersReply struct {
	ErrorInterface
	Response Stickers `json:"response"`
}

type Stickers struct {
	Stickers []Sticker `json:"items"`
}

type Sticker struct {
	ID        int    `json:"id"`
	Type      string `json:"stickers"`
	Purchased int    `json:"purchased"`
	Active    int    `json:"active"`
	Date      int    `json:"purchase_date"`
}
