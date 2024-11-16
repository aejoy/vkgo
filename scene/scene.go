package scene

import (
	"github.com/aejoy/vkgo/api"
	"github.com/aejoy/vkgo/update"
)

type Func func(*api.API, update.Update, Scenes)

type Scenes struct {
	MessageFunc        MessageFunc
	ReplyMessageFunc   ReplyMessageFunc
	EditMessageFunc    EditMessageFunc
	EnableMessageFunc  EnableMessageFunc
	DisableMessageFunc DisableMessageFunc
	TypingMessageFunc  TypingMessageFunc
	CallbackFunc       CallbackFunc

	PhotoFunc               PhotoFunc
	PhotoCommentFunc        PhotoCommentFunc
	EditPhotoCommentFunc    EditPhotoCommentFunc
	RestorePhotoCommentFunc RestorePhotoCommentFunc
	DeletePhotoCommentFunc  DeletePhotoCommentFunc

	NewAudioFunc NewAudioFunc

	NewVideoFunc            NewVideoFunc
	VideoCommentFunc        VideoCommentFunc
	EditVideoCommentFunc    EditVideoCommentFunc
	RestoreVideoCommentFunc RestoreVideoCommentFunc
	DeleteVideoCommentFunc  DeleteVideoCommentFunc

	PostFunc               PostFunc
	RepostFunc             RepostFunc
	PostCommentFunc        PostCommentFunc
	EditPostCommentFunc    EditPostCommentFunc
	RestorePostCommentFunc RestorePostCommentFunc
	DeletePostCommentFunc  DeletePostCommentFunc

	LikeFunc   LikeFunc
	UnlikeFunc UnlikeFunc

	BoardCommentFunc        BoardCommentFunc
	EditBoardCommentFunc    EditBoardCommentFunc
	RestoreBoardCommentFunc RestoreBoardCommentFunc
	DeleteBoardCommentFunc  DeleteBoardCommentFunc

	MarketCommentFunc        MarketCommentFunc
	EditMarketCommentFunc    EditMarketCommentFunc
	RestoreMarketCommentFunc RestoreMarketCommentFunc
	DeleteMarketCommentFunc  DeleteMarketCommentFunc

	OrderMarketFunc     OrderMarketFunc
	EditOrderMarketFunc EditOrderMarketFunc

	JoinGroupFunc  JoinGroupFunc
	LeaveGroupFunc LeaveGroupFunc

	BlockUserFunc   BlockUserFunc
	UnblockUserFunc UnblockUserFunc

	VoteFunc             VoteFunc
	EditOwnersFunc       EditOwnersFunc
	ControlGroupFunc     ControlGroupFunc
	ChangeGroupPhotoFunc ChangeGroupPhotoFunc
	VKPayTransactionFunc VKPayTransactionFunc
	AppNotificationFunc  AppNotificationFunc

	DonateFunc             DonateFunc
	ProlongedDonateFunc    ProlongedDonateFunc
	ExpiredDonateFunc      ExpiredDonateFunc
	CancelDonateFunc       CancelDonateFunc
	ChangeDonatePriceFunc  ChangeDonatePriceFunc
	WithdrawMoneyFunc      WithdrawMoneyFunc
	WithdrawMoneyErrorFunc WithdrawMoneyErrorFunc
}

type MessageFunc func(*api.API, update.Message)
type ReplyMessageFunc func(*api.API, update.Message)
type EditMessageFunc func(*api.API, update.Message)
type EnableMessageFunc func(*api.API, update.EnableMessage)
type DisableMessageFunc func(*api.API, update.DisableMessage)
type TypingMessageFunc func(*api.API, update.Typing)
type CallbackFunc func(*api.API, update.Callback)

type PhotoFunc func(*api.API, update.Photo)
type PhotoCommentFunc func(*api.API, update.PhotoComment)
type EditPhotoCommentFunc PhotoCommentFunc
type RestorePhotoCommentFunc PhotoCommentFunc
type DeletePhotoCommentFunc PhotoCommentFunc

type NewAudioFunc func(*api.API, update.Audio)

type NewVideoFunc func(*api.API, update.Video)
type VideoCommentFunc func(*api.API, update.VideoComment)
type EditVideoCommentFunc VideoCommentFunc
type RestoreVideoCommentFunc VideoCommentFunc
type DeleteVideoCommentFunc func(*api.API, update.DeleteVideoComment)

type PostFunc func(*api.API, update.Post)
type RepostFunc PostFunc
type PostCommentFunc func(*api.API, update.PostComment)
type EditPostCommentFunc PostCommentFunc
type RestorePostCommentFunc PostCommentFunc
type DeletePostCommentFunc func(*api.API, update.DeletePostComment)

type LikeFunc func(*api.API, update.Like)
type UnlikeFunc func(*api.API, update.Unlike)

type BoardCommentFunc func(*api.API, update.BoardComment)
type EditBoardCommentFunc BoardCommentFunc
type RestoreBoardCommentFunc BoardCommentFunc
type DeleteBoardCommentFunc func(*api.API, update.DeleteBoardComment)

type MarketCommentFunc func(*api.API, update.MarketComment)
type EditMarketCommentFunc MarketCommentFunc
type RestoreMarketCommentFunc MarketCommentFunc
type DeleteMarketCommentFunc func(*api.API, update.DeleteMarketComment)

type OrderMarketFunc func(*api.API, update.OrderMarket)
type EditOrderMarketFunc OrderMarketFunc

type JoinGroupFunc func(*api.API, update.JoinGroup)
type LeaveGroupFunc func(*api.API, update.LeaveGroup)

type BlockUserFunc func(*api.API, update.BlockUser)
type UnblockUserFunc func(*api.API, update.UnblockUser)

type VoteFunc func(*api.API, update.Vote)
type EditOwnersFunc func(*api.API, update.EditOwners)
type ControlGroupFunc func(*api.API, update.ControlGroup)
type ChangeGroupPhotoFunc func(*api.API, update.ChangeGroupPhoto)
type VKPayTransactionFunc func(*api.API, update.VKPayTransaction)
type AppNotificationFunc func(*api.API, update.AppNotification)

type DonateFunc func(*api.API, update.ProlongedDonate)
type ProlongedDonateFunc func(*api.API, update.ProlongedDonate)
type ExpiredDonateFunc func(*api.API, update.ExpiredDonate)
type CancelDonateFunc func(*api.API, update.ExpiredDonate)
type ChangeDonatePriceFunc func(*api.API, update.ChangeDonatePrice)
type WithdrawMoneyFunc func(*api.API, update.WithdrawMoney)
type WithdrawMoneyErrorFunc func(*api.API, update.WithdrawMoneyError)

func NewScene() *Scenes {
	return new(Scenes)
}
