# codegen

Used to initialite a new day directory.

## Example

Before:
```
.
|_codegen/
| |_sources/
| | |_input.txt
| | |_problem.txt
| | |_solution_test.go
| | |_solution.go
| |_main.go
```

Execute cmd:
```
go run . -year 2025 -day 1
```

After:
```
.
|_2025/
| |_01/
| | |_input.txt
| | |_problem.txt
| | |_solution_test.go
| | |_solution.go
|_codegen/
| |_sources/
| | |_input.txt
| | |_problem.txt
| | |_solution_test.go
| | |_solution.go
| |_main.go
```