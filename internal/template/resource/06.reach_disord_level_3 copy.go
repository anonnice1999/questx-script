package resource

type ReachDiscordLevel struct{}

func (*ReachDiscordLevel) Type() string {
	return "image"
}

func (*ReachDiscordLevel) Title() string {
	return "Reach Discord level 3"
}

func (*ReachDiscordLevel) Description() string {
	return "MISSION 🎯\n\nLove to chat or do you think that you are an interesting storyteller? Now you have quests too! Reach level 3 in discord by chatting with our community to boost your XP!\n\nSUBMISSION 📝\n\nReach level 3 in discord by chatting with the community.\n\nJoin The Project Discord \n\nGo to the 🤖bot-channel in our discord and use the \"!rank\" command, screenshot it and post here as proof. (cheaters will be have XP removed!)"
}

func (*ReachDiscordLevel) Recurrence() string {
	return "once"
}

func (*ReachDiscordLevel) Category() string {
	return SocialNetwork
}

func (*ReachDiscordLevel) Awards() []map[string]any {
	return []map[string]any{
		{"type": "points", "data": map[string]any{"points": 200}},
	}
}

func (*ReachDiscordLevel) ValidationData() map[string]any {
	return nil
}
