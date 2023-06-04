package resource

type WriteMedium struct{}

func (*WriteMedium) Type() string {
	return "url"
}

func (*WriteMedium) Title() string {
	return "Write a Medium article about Project"
}

func (*WriteMedium) Description() string {
	return `<p><strong>MISSION</strong> 🎯</p><p><br></p><p>Write an article about Project at least 300 words in any languages.</p><p><br></p><p><strong>GUIDE</strong> 📚</p><p><br></p><p>Be creative, be positive.</p><p><br></p><p>Type of article:</p><p><br></p><p>- information</p><p><br></p><p>- tutorial</p><p><br></p><p>- analyses</p><p><br></p><p>- cases studies</p><p><br></p><p><br></p><p><strong>SUBMISSION</strong> 📝</p><p><br></p><p>Upload your piece of art on Medium (make it public!)</p><p><br></p><p>Post your work on #✍-medium Channel on Discord&nbsp;Link</p>`
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
