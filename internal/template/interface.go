package template

type Template interface {
	Type() string
	Title() string
	Description() string
	Recurrence() string
	Awards() []map[string]any
	Category() string
	ValidationData() map[string]any
}
