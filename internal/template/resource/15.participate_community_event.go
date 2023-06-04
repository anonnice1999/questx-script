package resource

type ParticipateInCommunityEvent struct{}

func (*ParticipateInCommunityEvent) Type() string {
	return "text"
}

func (*ParticipateInCommunityEvent) Title() string {
	return "Weekly participate in Community event"
}

func (*ParticipateInCommunityEvent) Description() string {
	return `<p><strong>MISSION</strong> 🎯</p><p><br></p><p>Join the community Weekly event</p><p><br></p><p><strong>GUIDE</strong> 📚</p><p><br></p><p>Join in the community weekly event</p><p><br></p><p><strong>SUBMISSION</strong> 📝</p><p><br></p><p>Type yes if you participate in the event in the text area</p>`
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
