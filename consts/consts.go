package consts

const (
	DefaultAPIVersion = "5.238"
	CommunityAPILimit = 19
	UserAPILimit      = 2

	DefaultLongPollTimeout = 60
	DefaultLongPollMode    = 3

	CommunityBotType = 1
	UserBotType      = 2
)

const (
	UpdatesStatusOK = iota
	UpdatesStatusOutdated
	UpdatesStatusExpired
	UpdatesStatusLost
	UpdatesStatusInvalid
)
