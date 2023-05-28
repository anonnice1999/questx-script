package resource

type DailyKnowFrens struct{}

func (*DailyKnowFrens) Type() string {
	return "image"
}

func (*DailyKnowFrens) Title() string {
	return "Daily Know Frens"
}

func (*DailyKnowFrens) Description() string {
	return "MISSION 🎯\n\nOnly doing the daily tasks is boring. Buidlers should start building connections in this community.\n\nLet's start a conversation in 💬︱general Channel or international channel with others and meet some friends in our channel!\n\nGuide 📝\n\nStart a conversation in 💬︱general Link\n\nOr international rooms/channels with at least 3 sentences in your own languages"
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
