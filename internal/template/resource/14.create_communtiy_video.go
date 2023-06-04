package resource

type CreateCommunityVideo struct{}

func (*CreateCommunityVideo) Type() string {
	return "url"
}

func (*CreateCommunityVideo) Title() string {
	return "Create a Video about Project 📺"
}

func (*CreateCommunityVideo) Description() string {
	return `<p><strong>MISSION</strong> 🎯</p><p>Create a short video media about the Project</p><p><br></p><p><strong>GUIDE</strong> 📚</p><p>The Youtube Video must be from 3 minutes long.</p><p><br></p><p>The video can speak about anything. Example of videos:</p><p><br></p><p>- presentation</p><p><br></p><p>- product review or demo</p><p><br></p><p>- testimonials</p><p><br></p><p>- tutorial</p><p><br></p><p>- educational</p><p><br></p><p>- analyses</p><p><br></p><p>- cases studies</p><p><br></p><p>- ....Make sure you use hashtags</p><p><br></p><p><strong>SUBMISSION</strong> 📝</p><p>Upload the link of your video on the Youtube also our #🖥-youtube channel Link in discord</p><p><br></p>`
}

func (*CreateCommunityVideo) Recurrence() string {
	return "weekly"
}

func (*CreateCommunityVideo) Category() string {
	return SocialNetwork
}

func (*CreateCommunityVideo) Points() int {
	return 700
}

func (*CreateCommunityVideo) Rewards() []map[string]any {
	return nil
}

func (*CreateCommunityVideo) ValidationData() map[string]any {
	return nil
}
