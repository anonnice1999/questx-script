package resource

type ParticipateInCommunityEvent struct{}

func (*ParticipateInCommunityEvent) Type() string {
	return "text"
}

func (*ParticipateInCommunityEvent) Title() string {
	return "Weekly participate in Community event"
}

func (*ParticipateInCommunityEvent) Description() string {
	return "MISSION 🎯\n\nJoin the community Weekly event\n\nGUIDE 📚\n\nJoin in the community weekly event\n\nSUBMISSION 📝\n\nType yes if you participate in the event in the text area"
}

func (*ParticipateInCommunityEvent) Recurrence() string {
	return "weekly"
}

func (*ParticipateInCommunityEvent) Category() string {
	return Community
}

func (*ParticipateInCommunityEvent) Points() int {
	return 40
}

func (*ParticipateInCommunityEvent) Rewards() []map[string]any {
	return nil
}

func (*ParticipateInCommunityEvent) ValidationData() map[string]any {
	return map[string]any{"auto_validate": true, "answer": "yes"}
}