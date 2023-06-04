package resource

type InviteFriendToDiscord struct{}

func (*InviteFriendToDiscord) Type() string {
	return "invite_discord"
}

func (*InviteFriendToDiscord) Title() string {
	return "Invite friens to Discord"
}

func (*InviteFriendToDiscord) Description() string {
	return "<p><strong>MISSION 🎯</strong></p><p><br></p><p>Invite 3 persons in the Discord. Track your invites with invites tracker in Discord</p><p><br></p><p><strong>Guide</strong> 📝</p><p><br></p><p>Upload a screenshot of your invites in the discord</p><p><br></p><p>Create your own invite link</p><p><br></p><p>Use the command /invites in 🤖bot-channel</p><p><br></p><p>Take the picture and submit here</p>"
}

func (*InviteFriendToDiscord) Recurrence() string {
	return "once"
}

func (*InviteFriendToDiscord) Category() string {
	return SocialNetwork
}

func (*InviteFriendToDiscord) Points() int {
	return 200
}

func (*InviteFriendToDiscord) Rewards() []map[string]any {
	return nil
}

func (*InviteFriendToDiscord) ValidationData() map[string]any {
	return map[string]any{"number": 3}
}
