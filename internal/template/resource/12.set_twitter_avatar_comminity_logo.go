package resource

type SetTwitterAvatarToCommunityLogo struct{}

func (*SetTwitterAvatarToCommunityLogo) Type() string {
	return "image"
}

func (*SetTwitterAvatarToCommunityLogo) Title() string {
	return "Add Project logo to Twitter profile picture"
}

func (*SetTwitterAvatarToCommunityLogo) Description() string {
	return `<p><strong>MISSION</strong> 🎯</p><p><br></p><p>Add Logo to your Twitter Profile Picture.</p><p>You can find Logo here:&nbsp;Link</p><p><br></p><p><strong>SUBMISSION</strong> 📝</p><p><br></p><p>Submit your Twitter URL</p>`
}

func (*SetTwitterAvatarToCommunityLogo) Recurrence() string {
	return "once"
}

func (*SetTwitterAvatarToCommunityLogo) Category() string {
	return SocialNetwork
}

func (*SetTwitterAvatarToCommunityLogo) Points() int {
	return 100
}

func (*SetTwitterAvatarToCommunityLogo) Rewards() []map[string]any {
	return nil
}

func (*SetTwitterAvatarToCommunityLogo) ValidationData() map[string]any {
	return nil
}
