# Task Tracker CLI

A commend-line interface application built in Go for managing a personal todo list.
The application persists in a JSON file and allows for creating, updating, deleting and
filtering tasks.

## Instalation

1. Clone the repository:

```sh
git clone git@github.com:trobukan/task-tracker.git
```

2. Install dependencies

```sh
go mod tidy
```

3. Build the executable:

```sh
go build -o task-tracker
```

After building, you can run the application using `./task-tracker`.

## Usage

#### Add a task

```sh
./task-tracker add "Buy groceries"
```

#### List tasks

List all tasks

```sh
./task-tracker list
```

filter by the status

```sh
./task-tracker list done
./task-tracker list in-progress
./task-tracker list todo
```

#### Update a task

Update the description of a task by its list index:

```sh
./task-tracker update 1 "New Description"
```

#### Modify task status

```sh
./task-tracker mark-done 1
./task-tracker mask-in-progress 1
./task-tracker mask-todo 1
```

#### Delete a task

```sh
./task-tracker delete 1
```
