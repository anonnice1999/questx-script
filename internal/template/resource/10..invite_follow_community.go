package resource

type InviteFollowCommunity struct{}

func (*InviteFollowCommunity) Type() string {
	return "invite"
}

func (*InviteFollowCommunity) Title() string {
	return "Invite XQuest Frens"
}

func (*InviteFollowCommunity) Description() string {
	return "MISSION 🎯\n\n\n\nGUIDE 📚\n\nGo in your profile, and get your referral by clicking on the button \"invite frens\".\n\nSend your referral links to your friend so they can join our community and claim their first quest.\n\nYou can check your number of invites in your profile.\n\nSUBMISSION 📝\n\nThis quest will auto-validate if you have 1 valid invite."
}

func (*InviteFollowCommunity) Recurrence() string {
	return "once"
}

func (*InviteFollowCommunity) Category() string {
	return Community
}

func (*InviteFollowCommunity) Points() int {
	return 100
}

func (*InviteFollowCommunity) Rewards() []map[string]any {
	return nil
}

func (*InviteFollowCommunity) ValidationData() map[string]any {
	return nil
}
