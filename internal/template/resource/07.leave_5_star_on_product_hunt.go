package resource

type LeaveStarOnProductHunt struct{}

func (*LeaveStarOnProductHunt) Type() string {
	return "image"
}

func (*LeaveStarOnProductHunt) Title() string {
	return "Leave 5 Star on Product Hunt"
}

func (*LeaveStarOnProductHunt) Description() string {
	return "MISSION 🎯\n\nShare your feedback about our product on Product Hunt\n\nSUBMISSION 📝\n\nUpload a screenshot of your upvote"
}

func (*LeaveStarOnProductHunt) Recurrence() string {
	return "once"
}

func (*LeaveStarOnProductHunt) Category() string {
	return SocialNetwork
}

func (*LeaveStarOnProductHunt) Awards() []map[string]any {
	return []map[string]any{
		{"type": "points", "data": map[string]any{"points": 200}},
	}
}

func (*LeaveStarOnProductHunt) ValidationData() map[string]any {
	return nil
}
