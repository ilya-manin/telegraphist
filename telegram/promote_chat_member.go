package telegram

type PromoteChatMemberParams struct {
	ChatID             interface{} `json:"chat_id"`                        // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserID             int         `json:"user_id"`                        // Unique identifier of the target user
	CanChangeInfo      bool        `json:"can_change_info,omitempty"`      // Pass True, if the administrator can change chat title, photo and other settings
	CanPostMessages    bool        `json:"can_post_messages,omitempty"`    // Pass True, if the administrator can create channel posts, channels only
	CanEditMessages    bool        `json:"can_edit_messages,omitempty"`    // Pass True, if the administrator can edit messages of other users and can pin messages, channels only
	CanDeleteMessages  bool        `json:"can_delete_messages,omitempty"`  // Pass True, if the administrator can delete messages of other users
	CanInviteUsers     bool        `json:"can_invite_users,omitempty"`     // Pass True, if the administrator can invite new users to the chat
	CanRestrictMembers bool        `json:"can_restrict_members,omitempty"` // Pass True, if the administrator can restrict, ban or unban chat members
	CanPinMessages     bool        `json:"can_pin_messages,omitempty"`     // Pass True, if the administrator can pin messages, supergroups only
	CanPromoteMembers  bool        `json:"can_promote_members,omitempty"`  // Pass True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by him)
}

/*
Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Pass False for all boolean parameters to demote a user. Returns True on success.
*/
func (c *Client) PromoteChatMember(params *PromoteChatMemberParams) (bool, error) {
	var err error
	var result bool

	err = c.performRequest("promoteChatMember", params, nil, &result)

	return result, err
}
