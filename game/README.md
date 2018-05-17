This is a non-optimized code that generates all the winning possibilities for a player

```shell
$ go run main.go | tail
[8 1 2 3 6 4 5 0 7]
[8 1 2 3 6 0 5 4 7]
[8 1 2 4 5 3 6 0 7]
[8 1 2 4 5 0 6 3 7]
[8 1 2 4 6 3 5 0 7]
[8 1 2 4 6 0 5 3 7]
[8 1 2 0 5 3 6 4 7]
[8 1 2 0 5 4 6 3 7]
[8 1 2 0 6 3 5 4 7]
[8 1 2 0 6 4 5 3 7]


$ go run main.go | wc -l
  243936
```
