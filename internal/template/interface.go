package template

type Template interface {
	Type() string
	Title() string
	Description() string
	Recurrence() string
	Points() int
	Rewards() []map[string]any
	Category() string
	ValidationData() map[string]any
}
