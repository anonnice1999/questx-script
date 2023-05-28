package resource

type TryProductDemo struct{}

func (*TryProductDemo) Type() string {
	return "visit_link"
}

func (*TryProductDemo) Title() string {
	return "Try product demo"
}

func (*TryProductDemo) Description() string {
	return "MISSION 🎯\n\nRead this page\n\nSUBMISSION 📝\n\nIt will auto-validate if you click on the link."
}

func (*TryProductDemo) Recurrence() string {
	return "once"
}

func (*TryProductDemo) Category() string {
	return SocialNetwork
}

func (*TryProductDemo) Points() int {
	return 100
}

func (*TryProductDemo) Rewards() []map[string]any {
	return nil
}

func (*TryProductDemo) ValidationData() map[string]any {
	return map[string]any{"link": "https://example.com/demo/product"}
}
