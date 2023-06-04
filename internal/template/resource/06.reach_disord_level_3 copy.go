package resource

type ReachDiscordLevel struct{}

func (*ReachDiscordLevel) Type() string {
	return "image"
}

func (*ReachDiscordLevel) Title() string {
	return "Reach Discord level 3"
}

func (*ReachDiscordLevel) Description() string {
	return `<p><strong>MISSION</strong> 🎯</p><p><br></p><p>Love to chat or do you think that you are an interesting storyteller? Now you have quests too! Reach level 3 in discord by chatting with our community to boost your XP!</p><p><br></p><p><strong>SUBMISSION</strong> 📝</p><p><br></p><p>Reach level 3 in discord by chatting with the community.</p><p><br></p><p>Join The Project Discord</p><p><br></p><p>Go to the 🤖bot-channel in our discord and use the "!rank" command, screenshot it and post here as proof. (cheaters will be have XP removed!)</p>`
}

func (*ReachDiscordLevel) Recurrence() string {
	return "once"
}

func (*ReachDiscordLevel) Category() string {
	return SocialNetwork
}

func (*ReachDiscordLevel) Points() int {
	return 200
}

func (*ReachDiscordLevel) Rewards() []map[string]any {
	return nil
}

func (*ReachDiscordLevel) ValidationData() map[string]any {
	return nil
}
