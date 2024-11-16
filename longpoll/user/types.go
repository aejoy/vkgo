package user

type Messages struct {
	TS         any     `json:"ts,omitempty"`
	Updates    [][]any `json:"updates,omitempty"`
	Failed     int     `json:"failed,omitempty"`
	MaxVersion int     `json:"max_version"`
	MinVersion int     `json:"min_version"`
}
