# Методы

## Account

| Метод API                 | Метод библиотеки  | Состояние |
|---------------------------|-------------------|-----------|
| account.ban               | Ban               | ✅         |
| account.changePassword    | ChangePassword    | ✅         |
| account.getActiveOffers   | GetActiveOffers   | ✅         |
| account.getAppPermissions | GetAppPermissions | ✅         |
| account.getBanned         | GetBans           | ✅         |
| account.getCounters       | GetCounters       | ✅         |
| account.getInfo           | GetInfo           | ✅         |
| account.getProfileInfo    | GetProfileInfo    | ✅         |
| account.getPushSettings   | GetPushSettings   | ❌         |
| account.lookupContacts    |                   | 🚾        |
| account.registerDevice    |                   | ❌         |
| account.saveProfileInfo   | EditProfileInfo   | ✅         |
| account.setInfo           | EditInfo          | ✅         |
| account.setNameInMenu     |                   | 🚾        |
| account.setOffline        | SetOffline        | ✅         |
| account.setOnline         | SetOnline         | ✅         |
| account.setPushSettings   |                   | ❌         |
| account.setSilenceMode    |                   | ❌         |
| account.unban             | Unban             | ✅         |
| account.unregisterDevice  |                   | ❌         |

## Ads, Apps, Asr, Auth, Bugtracker

The current iteration skips this section of the methods

## Friends

| Метод API                 | Метод библиотеки         | Состояние |
|---------------------------|--------------------------|-----------|
| friends.add               | AddFriend                | ✅         |
| friends.addList           | AddFriendsList           | ✅         |
| friends.areFriends        | AreFriends               | ✅         |
| friends.delete            | DeleteFriend             | ✅         |
| friends.deleteAllRequests | DeleteAllFriendsRequests | ✅         |
| friends.deleteList        | DeleteFriendsList        | ✅         |
| friends.edit              | EditFriends              | ✅         |
| friends.editList          | EditFriendsList          | ✅         |
| friends.get               | GetFriends               | ✅         |
| friends.getAppUsers       | GetAppUsersFriends       | ❌         |
| friends.getLists          | GetFriendsLists          | ✅         |
| friends.getMutual         | GetMutualFriends         | ✅         |
| friends.getOnline         | GetOnlineFriends         | ✅         |
| friends.getRecent         | GetRecentFriends         | ✅         |
| friends.getRequests       | GetFriendsRequests       | ✅         |
| friends.getSuggestions    | GetSuggestionsFriends    | ✅         |
| friends.search            | SearchFriends            | ✅         |

## Groups

| Метод API                          | Метод библиотеки                 | Состояние |
|------------------------------------|----------------------------------|-----------|
| groups.addAddress                  | AddGroupAddress                  | ✅         |
| groups.addCallbackServer           | AddGroupCallbackServer           | ✅         |
| groups.addLink                     | AddGroupLink                     | ✅         |
| groups.approveRequest              | ApproveRequestGroup              | ✅         |
| groups.ban                         | BanUser                          | ✅         |
| groups.create                      | CreateGroup                      | ✅         |
| groups.deleteAddress               | DeleteGroupAddress               | ✅         |
| groups.deleteCallbackServer        | DeleteGroupCallbackServer        | ✅         |
| groups.deleteLink                  | DeleteGroupLink                  | ✅         |
| groups.disableOnline               | DeleteGroupCallbackServer        | ✅         |
| groups.disableOnline               | SetGroupOffline                  | ✅         |
| groups.edit                        | EditGroup                        | ✅         |
| groups.editAddress                 | EditAddressGroup                 | ❌         |
| groups.editCallbackServer          | EditGroupCallbackServer          | ✅         |
| groups.editLink                    | EditGroupLink                    | ✅         |
| groups.editManager                 | EditGroupManager                 | ✅         |
| groups.enableOnline                | SetGroupOnline                   | ✅         |
| groups.get                         | GetUserGroups                    | ✅         |
| groups.getById                     | GetGroup                         | ✅         |
| groups.getCallbackConfirmationCode | GetGroupCallbackConfirmationCode | ✅         |
| groups.getCallbackServers          | GetGroupCallbackServers          | ✅         |
| groups.getCallbackSettings         | GetGroupCallbackSettings         | ✅         |
| groups.getCatalogInfo              | GetGroupCatalogInfo              | ✅         |
| groups.getInvitedUsers             | GetGroupInvitedUsers             | ✅         |
| groups.getInvitedUsers             | GetGroupInvitedUsers             | ✅         |
| groups.getInvites                  | GetGroupInvites                  | ✅         |
| groups.getLongPollServer           | GetGroupLongPollServer           | ✅         |
| groups.getLongPollSettings         | GetGroupLongPollSettings         | ✅         |
| groups.getMembers                  | GetGroupMembers                  | ✅         |
| groups.getOnlineStatus             | GetGroupOnlineStatus             | ✅         |
| groups.getRequests                 | GetGroupRequests                 | ✅         |
| groups.getSettings                 | GetGroupSettings                 | ✅         |
| groups.getTagList                  | GetGroupTagList                  | ❌         |
| groups.getTokenPermissions         | GetGroupTokenPermissions         | ✅         |
| groups.invite                      | InviteGroup                      | ✅         |
| groups.isMember                    | IsGroupMember                    | ✅         |
| groups.join                        | JoinGroup                        | ✅         |
| groups.leave                       | LeaveGroup                       | ✅         |
| groups.removeUser                  | RemoveGroupUser                  | ✅         |
| groups.reorderLink                 | ReorderGroupLink                 | ✅         |
| groups.search                      | SearchGroup                      | ✅         |
| groups.setCallbackSettings         | SetGroupCallbackSettings         | ❌         |
| groups.setLongPollSettings         | SetGroupLongPollSettings         | ✅         |
| groups.setSettings                 | SetGroupSettings                 | ❌         |
| groups.setUserNote                 | SetGroupUserNote                 | ❌         |
| groups.tagAdd                      | TagGroupAdd                      | ❌         |
| groups.tagBind                     | TagGroupBind                     | ❌         |
| groups.tagDelete                   | TagGroupDelete                   | ❌         |
| groups.tagUpdate                   | TagGroupUpdate                   | ❌         |
| groups.toggleMarket                | ToggleGroupMarket                | ❌         |
| groups.unban                       | UnbanUser                        | ✅         |

## Likes

| Метод API     | Метод библиотеки | Состояние |
|---------------|------------------|-----------|
| likes.add     | AddLike          | ✅         |
| likes.delete  | DeleteLike       | ✅         |
| likes.getList | GetLikes         | ✅         |
| likes.isLiked | IsLiked          | ✅         |

## Messages

| Метод API                                     | Метод библиотеки                    | Состояние |
|-----------------------------------------------|-------------------------------------|-----------|
| messages.addChatUser                          | AddChatUser                         | ✅         |
| messages.allowMessagesFromGroup               | AllowMessagesFromGroup              | ✅         |
| messages.createChat                           | CreateChat                          | ✅         |
| messages.delete                               | Spam, Spams                         | ✅         |
| messages.delete                               | DeleteMessage, DeleteMessages       | ✅         |
| messages.deleteChatPhoto                      | DeleteChatPhoto                     | ✅         |
| messages.deleteConversation                   | DeleteChat                          | ✅         |
| messages.deleteReaction                       | DeleteReaction                      | ✅         |
| messages.denyMessagesFromGroup                | DenyMessagesFromGroup               | ✅         |
| messages.edit                                 | EditMessage                         | ✅         |
| messages.editChat                             | EditChatTitle                       | ✅         |
| messages.getByConversationMessageId           | GetChatMessageID, GetChatMessageIDs | 🚮        |
| messages.getById                              | GetMessage, GetMessages             | ✅         |
| messages.getChat                              | GetChat                             | 🚮        |
| messages.getChatPreview                       | GetChatPreview                      | ❌         |
| messages.getConversationMembers               | GetChatMembers                      | ✅         |
| messages.getConversations                     | GetMyChats                          | ✅         |
| messages.getConversationsById                 | GetChat, GetChats                   | ✅         |
| messages.getHistory                           | GetChatHistory                      | ❌         |
| messages.getHistoryAttachments                | GetChatHistoryAttachments           | ❌         |
| messages.getImportantMessages                 | GetChatImportantMessages            | ❌         |
| messages.getImportantMessages                 | GetChatImportantMessages            | ❌         |
| messages.getIntentUsers                       | GetIntentUsers                      | ✅         |
| messages.getInviteLink                        | GetChatLink                         | ✅         |
| messages.getLastActivity                      | GetLastActivity                     | ✅         |
| messages.getLongPollHistory                   | GetLongPollHistory                  | ❌         |
| messages.getLongPollServer                    | GetUserLongPollServer               | ✅         |
| messages.getMessagesReactions                 | GetMessagesReactions                | ✅         |
| messages.getReactedPeers                      | GetReactionaryMessages              | ✅         |
| messages.getReactionsAssets                   | GetReactionsAssets                  | ✅         |
| messages.isMessagesFromGroupAllowed           | IsMessagesFromGroupAllowed          | ✅         |
| messages.joinChatByInviteLink                 | JoinChatByInviteLink                | ✅         |
| messages.markAsAnsweredConversation           | MarkChatAsAnswered                  | ✅         |
| messages.markAsImportant                      | MarkMessagesAsImportant             | ❌         |
| messages.markAsImportantConversation          | MarkChatAsAnswered                  | ❌         |
| messages.markAsRead                           | MarkMessagesAsRead                  | ❌         |
| messages.markReactionsAsRead                  | MarkMessagesReactionAsRead          | ❌         |
| messages.pin                                  | PinMessage                          | ✅         |
| messages.removeChatUser                       | KickUser                            | ✅         |
| messages.restore                              | RestoreMessage                      | ✅         |
| messages.search                               | SearchMessages                      | ❌         |
| messages.searchConversations                  | SearchChats                         | ❌         |
| messages.send                                 | SendMessage                         | ✅         |
| messages.sendMessageEventAnswer               | SendMessageEventAnswer              | ✅         |
| messages.sendReaction                         | SendReaction                        | ✅         |
| messages.setActivity                          | SetActivity                         | ✅         |
| messages.setChatPhoto                         | SetChatPhoto                        | ❌         |
| messages.unpin                                | UnpinMessage                        | ❌         |
| messages.changeConversationMemberRestrictions | MuteUser                            | ❌         |

## Status

| Метод API  | Метод библиотеки | Состояние |
|------------|------------------|-----------|
| status.get | GetStatus        | ✅         |
| status.set | SetStatus        | ✅         |

## Storage

| Метод API       | Метод библиотеки | Состояние |
|-----------------|------------------|-----------|
| storage.get     | GetStatus        | ✅         |
| storage.getKeys | GetStorageKeys   | ✅         |
| storage.set     | SetStatus        | ✅         |

## Store

| Метод API         | Метод библиотеки | Состояние |
|-------------------|------------------|-----------|
| store.getProducts | GetStickers      | ✅         |

## Users

| Метод API              | Метод библиотеки  | Состояние |
|------------------------|-------------------|-----------|
| users.get              | GetUsers, GetUser | ✅         |
| users.getFollowers     | GetFollowers      | ✅         |
| users.getSubscriptions | GetSubscriptions  | ✅         |
| users.report           | ReportUser        | ✅         |
| users.search           | SearchUser        | ✅         |