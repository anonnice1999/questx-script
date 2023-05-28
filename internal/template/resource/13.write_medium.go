package resource

type WriteMedium struct{}

func (*WriteMedium) Type() string {
	return "url"
}

func (*WriteMedium) Title() string {
	return "Write a Medium article about Project"
}

func (*WriteMedium) Description() string {
	return "MISSION 🎯\n\nWrite an article about Project at least 300 words in any languages.\n\nGUIDE 📚\n\nBe creative, be positive.\nType of article:\n- information\n- tutorial\n- analyses\n\n- cases studies\n\nSUBMISSION 📝\n\nUpload your piece of art on Medium (make it public!)\n\nPost your work on #✍-medium Channel on Discord Link"
}

func (*WriteMedium) Recurrence() string {
	return "weekly"
}

func (*WriteMedium) Category() string {
	return SocialNetwork
}

func (*WriteMedium) Points() int {
	return 200
}

func (*WriteMedium) Rewards() []map[string]any {
	return nil
}

func (*WriteMedium) ValidationData() map[string]any {
	return nil
}
