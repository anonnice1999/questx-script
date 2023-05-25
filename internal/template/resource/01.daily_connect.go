package resource

type DailyConnectTemplate struct{}

func (*DailyConnectTemplate) Type() string {
	return "empty"
}

func (*DailyConnectTemplate) Title() string {
	return "Daily Connect"
}

func (*DailyConnectTemplate) Description() string {
	return "Claim your free Gem daily"
}

func (*DailyConnectTemplate) Recurrence() string {
	return "daily"
}

func (*DailyConnectTemplate) Category() string {
	return Community
}

func (*DailyConnectTemplate) Awards() []map[string]any {
	return []map[string]any{
		{"type": "points", "data": map[string]any{"points": 10}},
	}
}

func (*DailyConnectTemplate) ValidationData() map[string]any {
	return nil
}
