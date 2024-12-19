# Roadmap

## CLI
- [ ] Add commands 
    - [x] list all tasks
    - [x] list task by id
    - [x] create task
    - [x] update task
    - [ ] list task with filters
    - [ ] delete task

## TUI
- [ ] Kanban
- [x] Table
- [ ] Progress bar for projects?
- [ ] lipgloss list for tasks grouped by project


# Brainstorming

## CLI commands

List all tasks (table format)

```
./tazam
```

Show kanban board

```
./tazam kanban
(shortcut) ./tazam k
```

Create new task
```
./tazam this is a new task title
```

View a specific task (by ID)

```
./tazam show 2
(shortcut) ./tazam 2
```


Modify task (status, project, tag, etc.)

```
./tazam doing 2
./tazam done 2
./tazam modify 2 -p=project_name -t=newTag -d="this is a detailed description of the task"
```

Delete or archive task

```
./tazam delete 2
./tazam archive 2
```

## Deeper brainstorming for commands

add, list, list with filter, modify, archive/delete

```
tazam add Hello world first task
tazam list
tazam list 3
tazam list -p=project_name
tazam list -t=tag_name
tazam modify 3 -p=project_name -t=newTag -d="this is a detailed description of the task"
tazam archive 3
```

### status updates

provide no flags to the modify command to automatically transition task to next status
- will require a task status state-machine
- well defined status AND order of status (similar to orgmode and Obsidian's Task plugin)
```
tazam modify 3 
```

alternative, use the `-s/--status` flag explicitly

```
tazam modify 3 -s=doing
tazam modify 3 -s=done
```

**best of both worlds**
- allow use of flag
- infer status change if no flag provided
- also allow use of number for state transition (eg., todo=0, doing=1, done=2)
