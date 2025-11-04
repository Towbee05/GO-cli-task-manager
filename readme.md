# TASK CLI WITH GO
A simple and efficient task manager cli with Go programming language.
This idea was gotten from the delevoper [roadmap](https://roadmap.sh/). Click [here](https://roadmap.sh/projects/task-tracker) to view project


---

## 📌 Features
 - Add new task.
 - List all tasks.
 - filter task based on "todo", "in-progress", "done".
 - Mark task as "in-progress" and "done".
 - Update task
 - Delete task

---

## 🛠️ Installation
Ensure you have **Go** installed (version 1.25.2 preferrably).

1. Clone this project repo 
```bash
git clone git@github.com:Towbee05/GO-cli-task-manager.git
cd GO-cli-task-manager
```

2. Build and run:
```bash
go build -o task-cli
.\task-cli <command> 
```
## 🚀 Usage
1. Add a task 
```bash
task-cli add <task name> <description(optional)>
```

2. List all tasks and filter
```bash
task-cli list 
task-cli list todo
task-cli list done
task-cli list in-progress
```

3. Mark task as **in-progress** or **done**
```bash
task-cli mark-in-progress <id>
task-cli mark-done <id>
```

4. Update task
```bash
task-cli update <task_id> <new_name>
```

5. Delete task 
```bash
task-cli delete <task_id>
```

All tasks are saved in a JSON file
