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

func (*SetTwitterAvatarToCommunityLogo) Awards() []map[string]any {
	return []map[string]any{
		{"type": "points", "data": map[string]any{"points": 100}},
	}
}

func (*SetTwitterAvatarToCommunityLogo) ValidationData() map[string]any {
	return nil
}
