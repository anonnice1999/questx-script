package resource

type SetTwitterAvatarToCommunityLogo struct{}

func (*SetTwitterAvatarToCommunityLogo) Type() string {
	return "image"
}

func (*SetTwitterAvatarToCommunityLogo) Title() string {
	return "Add Project logo to Twitter profile picture"
}

func (*SetTwitterAvatarToCommunityLogo) Description() string {
	return "MISSION 🎯\n\nAdd Logo to your Twitter Profile Picture.\nYou can find Logo here: Link\n\nSUBMISSION 📝\n\nSubmit your Twitter URL"
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
