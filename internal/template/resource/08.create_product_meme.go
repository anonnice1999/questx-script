package resource

type CreateProductMeme struct{}

func (*CreateProductMeme) Type() string {
	return "image"
}

func (*CreateProductMeme) Title() string {
	return "Create product meme"
}

func (*CreateProductMeme) Description() string {
	return "MISSION 🎯\n\nCreate a funny meme about the project and post in 🤣︱memes Channel Link\n\nWe value creativity, good luck!\n\nSUBMISSION 📝\n\nUpload the Printscreen of you meme message  🤣︱memes Channel Link\n\nNote: Please do not copy"
}

func (*CreateProductMeme) Recurrence() string {
	return "weekly"
}

func (*CreateProductMeme) Category() string {
	return SocialNetwork
}

func (*CreateProductMeme) Awards() []map[string]any {
	return []map[string]any{
		{"type": "points", "data": map[string]any{"points": 200}},
	}
}

func (*CreateProductMeme) ValidationData() map[string]any {
	return nil
}
