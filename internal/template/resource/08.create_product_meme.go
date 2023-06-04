package resource

type CreateProductMeme struct{}

func (*CreateProductMeme) Type() string {
	return "image"
}

func (*CreateProductMeme) Title() string {
	return "Create product meme"
}

func (*CreateProductMeme) Description() string {
	return `<p><strong>MISSION</strong> 🎯</p><p><br></p><p>Create a funny meme about the project and post in 🤣︱memes Channel Link</p><p><br></p><p>We value creativity, good luck!</p><p><br></p><p><strong>SUBMISSION</strong> 📝</p><p><br></p><p>Upload the Printscreen of you meme message 🤣︱memes Channel&nbsp;Link</p><p><br></p><p>Note: Please do not copy</p>`
}

func (*CreateProductMeme) Recurrence() string {
	return "weekly"
}

func (*CreateProductMeme) Category() string {
	return SocialNetwork
}

func (*CreateProductMeme) Points() int {
	return 200
}

func (*CreateProductMeme) Rewards() []map[string]any {
	return nil
}

func (*CreateProductMeme) ValidationData() map[string]any {
	return nil
}
