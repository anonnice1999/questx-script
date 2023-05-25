package resource

type TwitterReaction struct{}

func (*TwitterReaction) Type() string {
	return "twitter_reaction"
}

func (*TwitterReaction) Title() string {
	return "Like, Reply and Retweet"
}

func (*TwitterReaction) Description() string {
	return "Just like reply and retweet the post, this quest will auto validate when you finish"
}

func (*TwitterReaction) Recurrence() string {
	return "once"
}

func (*TwitterReaction) Category() string {
	return SocialNetwork
}

func (*TwitterReaction) Awards() []map[string]any {
	return []map[string]any{
		{"type": "points", "data": map[string]any{"points": 100}},
	}
}

func (*TwitterReaction) ValidationData() map[string]any {
	return map[string]any{
		"like": true, "reply": true, "retweet": true,
		"tweet_url": "https://twitter.com/elonmusk/status/1652849795336159233",
	}
}
