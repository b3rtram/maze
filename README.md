# maze_generation_go
This is a maze generation implementation for Golang. It implements the growing tree algorithm as described in this link:  [Buckblog](https://weblog.jamisbuck.org/2011/1/27/maze-generation-growing-tree-algorithm)

PrintConsole is only tested on Macbook terminal. It uses underlining so i could be that it does not work correctly in other terminals.

This is a result:

![Example](example.png)

The result is 2D array with ints:

[1 13 9 9 25 9 1 17 5 5]
[9 3 17 5 5 17 5 5 9 19]
[3 1 5 17 5 17 9 17 3 1]
[19 5 5 5 3 13 3 17 3 5]
[9 9 17 1 9 17 1 17 9 3]
[17 5 5 3 3 17 3 5 19 5]
[17 1 5 7 3 17 9 17 17 3]
[9 9 17 3 3 9 3 17 1 3]
[17 5 9 19 3 9 1 9 17 3]
[1 3 5 5 3 3 5 5 13 3]

You can use this flags to find out the directions of a cell with bit wise and operation

```
const (
	U = 2
	L = 4
	R = 8
	D = 16
)

if grid[y][x] & U == U {
    //here you know that the way to the upper cell is open
}

if grid[y][x] & R == R {
    //here you know that the way to the right cell is open
}

//and so on

```
