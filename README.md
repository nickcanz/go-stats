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
count: 10
min: 1.00
avg: 5.50
max: 10.00
p50: 5.00
p95: 9.50
p99: 9.50

# Read from stdin

$ cat numbers | go-stats
count: 10
min: 1.00
avg: 5.50
max: 10.00
p50: 5.00
p95: 9.50
p99: 9.50

# Declare delimter and field to calculate on

$ cat data.csv
a,10000
b,20000
# ...

$ go-stats -d "," -f 2 data.csv
count: 98969
min: 55.00
avg: 135477.68
max: 745191934.00
p50: 5929.00
p95: 352395.00
p99: 3397098.00
```
