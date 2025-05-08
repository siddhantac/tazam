package task

type status int

const (
	todo status = iota
	doing
	done

	unknown = -1
)

var statusStateMachine = []status{
	todo,
	doing,
	done,
}

func (s status) String() string {
	if s == unknown {
		return "unknown"
	}
	return [...]string{"todo", "doing", "done"}[s]
}

func (s status) Next() status {
	next := int(s) + 1
	if next >= len(statusStateMachine) {
		next = 0
	}
	return statusStateMachine[next]
}
