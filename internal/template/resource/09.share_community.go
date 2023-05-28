package resource

type ShareCommunity struct{}

func (*ShareCommunity) Type() string {
	return "url"
}

func (*ShareCommunity) Title() string {
	return "Share Project to your Local Community and Friends😻"
}

func (*ShareCommunity) Description() string {
	return "MISSION 🎯\n\nShare Information to one of your Local Communities. \n\nIt could be a Local Online forum, Facebook page, Whatsapp Group or even your high school group!\n\nGUIDE 📚\n\nShare to someone you know or one of your local community\n\nSUBMISSION 📝\n\nLink to your shared post"
}

func (*ShareCommunity) Recurrence() string {
	return "daily"
}

func (*ShareCommunity) Category() string {
	return SocialNetwork
}

func (*ShareCommunity) Points() int {
	return 150
}

func (*ShareCommunity) Rewards() []map[string]any {
	return nil
}

func (*ShareCommunity) ValidationData() map[string]any {
	return nil
}
