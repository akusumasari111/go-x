package dtask

type Task interface {
	Bind(data any) error
	TaskID() string
	Type() string
}

// 2
// 3
// 4
// 5
