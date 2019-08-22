package telegram

/*
ChatMember contains information about one member of a chat.
*/
type ChatMember struct {
	User                  *User  `json:"user"`                                // Information about the user
	Status                string `json:"status"`                              // The member's status in the chat. Can be “creator”, “administrator”, “member”, “restricted”, “left” or “kicked”
	UntilDate             int    `json:"until_date,omitempty"`                // Optional. Restricted and kicked only. Date when restrictions will be lifted for this user; unix time
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`             // Optional. Administrators only. True, if the bot is allowed to edit administrator privileges of that user
	CanPostMessages       bool   `json:"can_post_messages,omitempty"`         // Optional. Administrators only. True, if the administrator can post in the channel; channels only
	CanEditMessages       bool   `json:"can_edit_messages,omitempty"`         // Optional. Administrators only. True, if the administrator can edit messages of other users and can pin messages; channels only
	CanDeleteMessages     bool   `json:"can_delete_messages,omitempty"`       // Optional. Administrators only. True, if the administrator can delete messages of other users
	CanRestrictMembers    bool   `json:"can_restrict_members,omitempty"`      // Optional. Administrators only. True, if the administrator can restrict, ban or unban chat members
	CanPromoteMembers     bool   `json:"can_promote_members,omitempty"`       // Optional. Administrators only. True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo         bool   `json:"can_change_info,omitempty"`           // Optional. Administrators and restricted only. True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers        bool   `json:"can_invite_users,omitempty"`          // Optional. Administrators and restricted only. True, if the user is allowed to invite new users to the chat
	CanPinMessages        bool   `json:"can_pin_messages,omitempty"`          // Optional. Administrators and restricted only. True, if the user is allowed to pin messages; groups and supergroups only
	IsMember              bool   `json:"is_member,omitempty"`                 // Optional. Restricted only. True, if the user is a member of the chat at the moment of the request
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`         // Optional. Restricted only. True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMediaMessages  bool   `json:"can_send_media_messages,omitempty"`   // Optional. Restricted only. True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`            // Optional. Restricted only. True, if the user is allowed to send polls
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`   // Optional. Restricted only. True, if the user is allowed to send animations, games, stickers and use inline bots
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"` // Optional. Restricted only. True, if the user is allowed to add web page previews to their messages
}
