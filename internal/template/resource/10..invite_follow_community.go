package resource

type InviteFollowCommunity struct{}

func (*InviteFollowCommunity) Type() string {
	return "invite"
}

func (*InviteFollowCommunity) Title() string {
	return "Invite XQuest Frens"
}

func (*InviteFollowCommunity) Description() string {
	return `<p><strong>MISSION</strong> 🎯</p><p><br></p><p>Invite your frens to join our community</p><p><br></p><p><strong>GUIDE</strong> 📚</p><p><br></p><p>Go in your profile, and get your referral by clicking on the button "invite frens".&nbsp;</p><p><br></p><p>Send your referral links to your friend so they can join our community and claim their first quest.</p><p><br></p><p>You can check your number of invites in your profile.</p><p><br></p><p><strong>SUBMISSION</strong> 📝</p><p><br></p><p>This quest will auto-validate if you have 1 valid invite.</p>`
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
