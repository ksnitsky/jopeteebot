package telegram_m

type Update struct {
	UpdateId int `json:"update_id"`
	Message *Message `json:"message,omitempty"`
}

func (u *Update) SentFrom() *User {
	switch {
	case u.Message != nil:
		return u.Message.From
	default:
		return nil
	}
}

type Message struct {
	MessageID int `json:"message_id"`
	// From is a sender, empty for messages sent to channels;
	//
	// optional
	From *User `json:"from,omitempty"`
	Text string `json:"text,omitempty"`
}

type User struct {
	// ID is a unique identifier for this user or bot
	ID int64 `json:"id"`
	// IsBot true, if this user is a bot
	//
	// optional
	IsBot bool `json:"is_bot,omitempty"`
	// IsPremium true, if user has Telegram Premium
	//
	// FirstName user's or bot's first name
	FirstName string `json:"first_name"`
	// LastName user's or bot's last name
	//
	// optional
	LastName string `json:"last_name,omitempty"`
	// UserName user's or bot's username
	//
	// optional
	UserName string `json:"username,omitempty"`
	// LanguageCode IETF language tag of the user's language
	// more info: https://en.wikipedia.org/wiki/IETF_language_tag
	//
	// optional
	LanguageCode string `json:"language_code,omitempty"`
	// CanJoinGroups is true, if the bot can be invited to groups.
	// Returned only in getMe.
	//
	// optional
	// CanJoinGroups bool `json:"can_join_groups,omitempty"`
	// CanReadAllGroupMessages is true, if privacy mode is disabled for the bot.
	// Returned only in getMe.
	//
	// optional
	// CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`
	// SupportsInlineQueries is true, if the bot supports inline queries.
	// Returned only in getMe.
	//
	// optional
	// SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}

type Chat struct {
	// ID is a unique identifier for this chat
	ID int64 `json:"id"`
	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`
	// Title for supergroups, channels and group chats
	//
	// optional
	Title string `json:"title,omitempty"`
	// UserName for private chats, supergroups and channels if available
	//
	// optional
	UserName string `json:"username,omitempty"`
	// FirstName of the other party in a private chat
	//
	// optional
	FirstName string `json:"first_name,omitempty"`
	// LastName of the other party in a private chat
	//
	// optional
	LastName string `json:"last_name,omitempty"`
	// Photo is a chat photo
	// Bio is the bio of the other party in a private chat. Returned only in
	// getChat
	//
	// optional
	// Bio string `json:"bio,omitempty"`
	// HasPrivateForwards is true if privacy settings of the other party in the
	// private chat allows to use tg://user?id=<user_id> links only in chats
	// with the user. Returned only in getChat.
	//
	// optional
	// HasPrivateForwards bool `json:"has_private_forwards,omitempty"`
	// Description for groups, supergroups and channel chats
	//
	// optional
	// Description string `json:"description,omitempty"`
	// InviteLink is a chat invite link, for groups, supergroups and channel chats.
	// Each administrator in a chat generates their own invite links,
	// so the bot must first generate the link using exportChatInviteLink
	//
	// optional
	// InviteLink string `json:"invite_link,omitempty"`
	// PinnedMessage is the pinned message, for groups, supergroups and channels
	//
	// optional
	// PinnedMessage *Message `json:"pinned_message,omitempty"`
	// Permissions are default chat member permissions, for groups and
	// supergroups. Returned only in getChat.
	//
	// optional
	// Permissions *ChatPermissions `json:"permissions,omitempty"`
	// SlowModeDelay is for supergroups, the minimum allowed delay between
	// consecutive messages sent by each unprivileged user. Returned only in
	// getChat.
	//
	// optional
	// SlowModeDelay int `json:"slow_mode_delay,omitempty"`
	// MessageAutoDeleteTime is the time after which all messages sent to the
	// chat will be automatically deleted; in seconds. Returned only in getChat.
	//
	// optional
	// MessageAutoDeleteTime int `json:"message_auto_delete_time,omitempty"`
	// HasProtectedContent is true if messages from the chat can't be forwarded
	// to other chats. Returned only in getChat.
	//
	// optional
	// HasProtectedContent bool `json:"has_protected_content,omitempty"`
	// StickerSetName is for supergroups, name of group sticker set.Returned
	// only in getChat.
	//
	// optional
	// StickerSetName string `json:"sticker_set_name,omitempty"`
	// CanSetStickerSet is true, if the bot can change the group sticker set.
	// Returned only in getChat.
	//
	// optional
	// CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
	// LinkedChatID is a unique identifier for the linked chat, i.e. the
	// discussion group identifier for a channel and vice versa; for supergroups
	// and channel chats.
	//
	// optional
	// LinkedChatID int64 `json:"linked_chat_id,omitempty"`
	// Location is for supergroups, the location to which the supergroup is
	// connected. Returned only in getChat.
	//
	// optional
	// Location *ChatLocation `json:"location,omitempty"`
}
