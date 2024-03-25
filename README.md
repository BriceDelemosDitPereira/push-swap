# Push-Swap

Push-Swap is a project that implements a Non-Comparative Sorting Algorithm to sort a stack of integer values using a specific set of instructions.
This project was carried out as part of the training : [Zone01](https://zone01rouennormandie.org/).
The project consists of two programs:

1. **push-swap**: This program calculates and displays the smallest list of instructions required to sort the stack provided as input. Instructions are written in the push-swap instruction language.

2. **checker**: This program takes integer arguments and reads instructions from the standard input. After executing the instructions on the provided stack, it checks if the stack is sorted and displays "OK" if sorted, otherwise "KO".


## Instructions

The following instructions can be used to sort the stack:

- **pa**: Push the top element from stack b to stack a.
- **pb**: Push the top element from stack a to stack b.
- **sa**: Swap the first two elements of stack a.
- **sb**: Swap the first two elements of stack b.
- **ss**: Execute both sa and sb.
- **ra**: Rotate stack a (shift all elements up by 1, first element becomes last).
- **rb**: Rotate stack b.
- **rr**: Execute both ra and rb.
- **rra**: Reverse rotate stack a (shift all elements down by 1, last element becomes first).
- **rrb**: Reverse rotate stack b.
- **rrr**: Execute both rra and rrb.


## Example

```
Init a and b:
2
1
3
6
8
5
= =
a b
---------------------------------------
Exec sa:
1
2
3
6
8
5
= =
a b
---------------------------------------
Exec pb pb pb:
6 3
8 2
5 1
= =
a b
---------------------------------------
Exec rb:
6 2
8 1
5 3
= =
a b
---------------------------------------
Exec rra and rrb (equivalent to rrr):
5 3
6 2
8 1
= =
a b
---------------------------------------
Exec pa pa pa:
1
2
3
5
6
8
= =
a b
```

## push-swap

- This program receives the stack a as input, formatted as a list of integers (top element first).
- It displays the smallest list of instructions required to sort the stack a, with the smallest number at the top.
- Each instruction is separated by a newline character.
- In case of errors (e.g., non-integer arguments, duplicates), it displays "Error" followed by a newline on the standard error.
- If no arguments are provided, it does not display anything (0 instructions).


### Usage Example

```
$ ./push-swap "2 1 3 6 5 8"
pb
pb
ra
sa
rrr
pa
pa
$ ./push-swap "0 one 2 3"
Error
$ ./push-swap
$
```


## checker

- This program receives the stack a as input, formatted as a list of integers (top element first).
- It reads instructions from the standard input, each followed by a newline character, and executes them on the stack.
- After executing the instructions, if stack a is sorted and stack b is empty, it displays "OK" followed by a newline on the standard output; otherwise, it displays "KO".
- In case of errors (e.g., non-integer arguments, duplicates, invalid instructions), it displays "Error" followed by a newline on the standard error.
- If no arguments are provided, it does not display anything.


### Usage Example

```
$ ./checker "3 2 1 0"
sa
rra
pb
KO
$ echo -e "rra\npb\nsa\n" | ./checker "3 2 one 0"
Error
$ echo -e "rra\npb\nsa\nrra\npa"
rra
pb
sa
rra
pa
$ echo -e "rra\npb\nsa\nrra\npa" | ./checker "3 2 1 0"
OK
$ ./checker
$
```


## Contributing

Contributions to Push-Swap are welcome! If you have any ideas for improvements or bug fixes, feel free to open an issue or submit a pull request.


## Co-developers

- [Delestre Thomas](https://github.com/Thomas-Delestre)
- [Marchais Mickael](https://github.com/Jeancrock)



*********************************************************************
# Push-Swap

## Introduction

Push-Swap is a tool for sorting a series of numbers in ascending order using various commands. Each command is printed to indicate how many commands are needed to sort the numbers in ascending order and which command is used.

This code also includes a checker that verifies if the executed commands correctly sort the numbers and if the result is accurate. It also checks if a series of commands can sort an example series of numbers.

It is possible to run this code without creating an executable, but it increases the number of elements to input in the terminal:

```
Example:
with an executable -> ./push-swap "3 1 2"
without an executable -> go run cmd/push-swap/main.go "3 1 2"
```

## Creating an Executable

To create an executable (as they are not in the repository), simply enter the following command in the terminal:

```
Example:
for push-swap -> go build -o push-swap cmd/push-swap/main.go
for checker -> go build -o checker cmd/checker/main.go
```

In these examples, the name of our executable will be respectively push-swap and checker.

## Usage

To run the code, launch the main.go of push-swap (or the executable) followed by a series of numbers (in string format) separated by spaces.

```
Example:
go run cmd/push-swap/main.go "3 1 2"
./push-swap "3 1 2"
```

There are 2 possibilities to run the checker:

1. Provide it with commands and numbers:

```bash
echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2"
```

2. Run push-swap at the same time:

```bash
./push-swap "0 9 1 8 2" | ./checker "0 9 1 8 2"
```

You can set a variable in the terminal to avoid rewriting the numbers:

```bash
ARG=("4 67 3 87 23"); ./push-swap "$ARG" | ./checker "$ARG"
```

This will use the values stored in the `ARG` variable for both push-swap and checker commands.

Feel free to reach out if you have any further questions or need additional assistance!