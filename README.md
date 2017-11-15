# go-stats

A command line tool to calculate some basic statistics.

```
# Example data
$ cat numbers
1
2
3
4
5
6
7
8
9
10

$ go-stats numbers
Count: 10
Min: 1
Max: 10
Average: 5.5

# Read from stdin

$ cat numbers | go-stats
Count: 10
Min: 1
Max: 10
Average: 5.5

# Declare delimter and field to calculate on

$ cat numbers.csv
a,1
b,2
c,3
d,4
e,5
f,6
g,7
h,8
i,9
j,10

$ go-stats -d "," -f 2 numbers.csv
Count: 10
Min: 1
Max: 10
Average: 5.5

```
