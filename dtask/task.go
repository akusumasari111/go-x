package dtask

type Task interface {
	Bind(data any) error
	TaskID() string
	Type() string
}

// 2
// Memperbarui README.md dengan petunjuk instalasi
