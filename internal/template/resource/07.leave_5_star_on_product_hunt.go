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

func (*LeaveStarOnProductHunt) Points() int {
	return 200
}

func (*LeaveStarOnProductHunt) Rewards() []map[string]any {
	return nil
}

func (*LeaveStarOnProductHunt) ValidationData() map[string]any {
	return nil
}
