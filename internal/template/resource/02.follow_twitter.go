package resource

type FollowTwitter struct{}

func (*FollowTwitter) Type() string {
	return "twitter_follow"
}

func (*FollowTwitter) Title() string {
	return "Follow Twitter"
}

func (*FollowTwitter) Description() string {
	return `<p><strong>MISSION</strong> 🎯</p><p><br></p><p>Follow us on twitter</p><p><br></p><p><strong>GUIDE</strong> 📚</p><p><br></p><p>Follow us and this quest will automatically claim</p><p><strong>SUBMISSION</strong> 📝</p><p><br></p><p>This is automatically validate quest</p>`
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
