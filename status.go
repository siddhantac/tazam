package main

type status int

const (
	todo status = iota
	inProgress
	done

	unknown = -1
)

var statusStateMachine = []status{
	todo,
	inProgress,
	done,
}

func (s status) String() string {
	if s == unknown {
		return "unknown"
	}
	return [...]string{"todo", "in-progress", "done"}[s]
}

func StatusFromString(s string) status {
	switch s {
	case "todo":
		return todo
	case "in progress":
		return inProgress
	case "done":
		return done
	}
	return unknown
}
