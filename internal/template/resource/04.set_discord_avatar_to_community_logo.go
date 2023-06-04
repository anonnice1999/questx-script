package resource

type SetDiscordAvatarToCommunityLogo struct{}

func (*SetDiscordAvatarToCommunityLogo) Type() string {
	return "image"
}

func (*SetDiscordAvatarToCommunityLogo) Title() string {
	return "Add Project logo to Discord profile picture"
}

func (*SetDiscordAvatarToCommunityLogo) Description() string {
	return `<p><strong>MISSION</strong> 🎯</p><p><br></p><p>Add Project Logo to your Discord Profile Picture.</p><p>You can find Project Logo here:&nbsp;Link</p><p><br></p><p><strong>SUBMISSION</strong> 📝</p><p><br></p><p>Screenshoot your profile picture please</p>`
}

func (*SetDiscordAvatarToCommunityLogo) Recurrence() string {
	return "once"
}

func (*SetDiscordAvatarToCommunityLogo) Category() string {
	return SocialNetwork
}

func (*SetDiscordAvatarToCommunityLogo) Points() int {
	return 100
}

func (*SetDiscordAvatarToCommunityLogo) Rewards() []map[string]any {
	return nil
}

func (*SetDiscordAvatarToCommunityLogo) ValidationData() map[string]any {
	return nil
}
