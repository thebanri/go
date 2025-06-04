# Task CLI - Simple TODO Manager

A command-line task management tool written in Go. Easily add, update, delete, and track the status of your tasks directly from your terminal.

## Features

- Create new tasks with descriptions
- List all tasks with their details
- Update task descriptions
- Mark tasks as "in-progress" or "done"
- Delete tasks
- Simple JSON-based storage

## Installation

### Prerequisites

- Go 1.24.3 or higher
- Sudo privileges (for installation)

### Install from source

1. Clone this repository:

   ```bash
   git clone https://github.com/thebanri/go/tree/main/TODO
   cd task
   ```

2. Run the installation script:
   ```bash
   ./install.sh
   ```

This script will:

- Build the Go executable
- Create necessary directories
- Set up the tasks.json file
- Install the binary to /usr/local/bin for system-wide access

## Usage

### Available Commands

```
task list                           - List all tasks
task add "<description>"            - Add a new task
task update <id> "<new description>" - Update a task's description
task mark-in-progress <id>          - Mark a task as 'in-progress'
task mark-done <id>                 - Mark a task as 'done'
task delete <id>                    - Delete a task
task help                           - Show this help message
```

### Examples

Add a new task:

```bash
task add "Complete the README file"
```

List all tasks:

```bash
task list
```

Mark a task as in-progress:

```bash
task mark-in-progress 1
```

Mark a task as done:

```bash
task mark-done 1
```

Update a task description:

```bash
task update 1 "Complete the improved README file"
```

Delete a task:

```bash
task delete 1
```

## Data Storage

Task data is stored in `/etc/taskcli/tasks.json`. Each task contains:

- ID
- Description
- Status (todo, in-progress, done)
- Creation date
- Last update date

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
