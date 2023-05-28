package resource

type FollowTwitter struct{}

func (*FollowTwitter) Type() string {
	return "twitter_follow"
}

func (*FollowTwitter) Title() string {
	return "Follow Twitter"
}

func (*FollowTwitter) Description() string {
	return "MISSION 🎯\n\nFollow us on twitter\n\nhttps://twitter.com/exemple\n\nGUIDE 📚\n\nFollow us and this quest will automatically claim\n\nSUBMISSION 📝\n\nThis is automatically validate quest"
}

func (*FollowTwitter) Recurrence() string {
	return "once"
}

func (*FollowTwitter) Category() string {
	return SocialNetwork
}

func (*FollowTwitter) Points() int {
	return 10
}

func (*FollowTwitter) Rewards() []map[string]any {
	return nil
}

func (*FollowTwitter) ValidationData() map[string]any {
	return map[string]any{"twitter_handle": "https://twitter.com/exemple"}
}
