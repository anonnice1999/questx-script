package resource

type InviteFriendToDiscord struct{}

func (*InviteFriendToDiscord) Type() string {
	return "invite_discord"
}

func (*InviteFriendToDiscord) Title() string {
	return "Invite friend to Discord"
}

func (*InviteFriendToDiscord) Description() string {
	return "MISSION 🎯\n\nInvite 3 persons in the Discord. Track your invites with invites tracker in Discord\n\nGuide 📝\n\nUpload a screenshot of your invites in the discord\n\nCreate your own invite link\n\nUse the command /invites in 🤖bot-channel"
}

func (*InviteFriendToDiscord) Recurrence() string {
	return "once"
}

func (*InviteFriendToDiscord) Category() string {
	return SocialNetwork
}

func (*InviteFriendToDiscord) Awards() []map[string]any {
	return []map[string]any{
		{"type": "points", "data": map[string]any{"points": 200}},
	}
}

func (*InviteFriendToDiscord) ValidationData() map[string]any {
	return map[string]any{"number": 3}
}
