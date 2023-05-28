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

func (*DailyConnectTemplate) Points() int {
	return 10
}

func (*DailyConnectTemplate) Rewards() []map[string]any {
	return nil
}

func (*DailyConnectTemplate) ValidationData() map[string]any {
	return nil
}
