package resource

type DailyKnowFrens struct{}

func (*DailyKnowFrens) Type() string {
	return "image"
}

func (*DailyKnowFrens) Title() string {
	return "Daily Know Frens"
}

func (*DailyKnowFrens) Description() string {
	return `<p><strong>MISSION</strong> 🎯</p><p><br></p><p>Only doing the daily tasks is boring. Buidlers should start building connections in this community.&nbsp;</p><p><br></p><p>Let's start a conversation in 💬︱general Channel or international channel with others and meet some friends in our channel!&nbsp;</p><p><br></p><p><strong>Guide</strong> 📝</p><p><br></p><p>Start a conversation in 💬︱general Link</p><p><br></p><p>Or international rooms/channels with at least 3 sentences in your own languages</p>`
}

func (*DailyKnowFrens) Recurrence() string {
	return "daily"
}

func (*DailyKnowFrens) Category() string {
	return SocialNetwork
}

func (*DailyKnowFrens) Points() int {
	return 150
}

func (*DailyKnowFrens) Rewards() []map[string]any {
	return nil
}

func (*DailyKnowFrens) ValidationData() map[string]any {
	return nil
}
