package resource

type ShareCommunity struct{}

func (*ShareCommunity) Type() string {
	return "url"
}

func (*ShareCommunity) Title() string {
	return "Share Project to your Local Community and Friends😻"
}

func (*ShareCommunity) Description() string {
	return `<p><strong>MISSION</strong> 🎯</p><p><br></p><p>Share Information to one of your Local Communities.&nbsp;</p><p><br></p><p>It could be a Local Online forum, Facebook page, Whatsapp Group or even your high school group!</p><p><br></p><p><strong>GUIDE</strong> 📚</p><p><br></p><p>Share to someone you know or one of your local community</p><p><br></p><p><strong>SUBMISSION</strong> 📝</p><p><br></p><p>Link to your shared post</p>`
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
