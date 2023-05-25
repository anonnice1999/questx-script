package resource

type WriteMedium struct{}

func (*WriteMedium) Type() string {
	return "image"
}

func (*WriteMedium) Title() string {
	return "Add Project logo to Twitter profile picture"
}

func (*WriteMedium) Description() string {
	return "MISSION 🎯\n\nAdd Logo to your Twitter Profile Picture.\nYou can find Logo here: Link\n\nSUBMISSION 📝\n\nSubmit your Twitter URL"
}

func (*WriteMedium) Recurrence() string {
	return "once"
}

func (*WriteMedium) Category() string {
	return SocialNetwork
}

func (*WriteMedium) Awards() []map[string]any {
	return []map[string]any{
		{"type": "points", "data": map[string]any{"points": 100}},
	}
}

func (*WriteMedium) ValidationData() map[string]any {
	return nil
}
