package resource

type SetDiscordAvatarToCommunityLogo struct{}

func (*SetDiscordAvatarToCommunityLogo) Type() string {
	return "image"
}

func (*SetDiscordAvatarToCommunityLogo) Title() string {
	return "Add Project logo to Discord profile picture"
}

func (*SetDiscordAvatarToCommunityLogo) Description() string {
	return "MISSION 🎯\n\nAdd Project Logo to your Discord Profile Picture.\n\nYou can find Project Logo here: Link\n\nSUBMISSION 📝\n\nScreenshoot your profile picture please"
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
