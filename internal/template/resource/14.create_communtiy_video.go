package resource

type CreateCommunityVideo struct{}

func (*CreateCommunityVideo) Type() string {
	return "url"
}

func (*CreateCommunityVideo) Title() string {
	return "Create a Video about Project 📺"
}

func (*CreateCommunityVideo) Description() string {
	return "MISSION 🎯\nCreate a  short video media about the Project\n\nGUIDE 📚\nThe Youtube Video must be from 3 minutes long.\n\nThe video can speak about anything. Example of videos:\n- presentation\n- product review or demo\n- testimonials \n- tutorial\n- educational\n- analyses\n- cases studies\n- ....Make sure you use hashtags\n\nSUBMISSION 📝\nUpload the link of your video on the Youtube also our #🖥-youtube channel Link in discord"
}

func (*CreateCommunityVideo) Recurrence() string {
	return "weekly"
}

func (*CreateCommunityVideo) Category() string {
	return SocialNetwork
}

func (*CreateCommunityVideo) Awards() []map[string]any {
	return []map[string]any{
		{"type": "points", "data": map[string]any{"points": 700}},
	}
}

func (*CreateCommunityVideo) ValidationData() map[string]any {
	return nil
}
