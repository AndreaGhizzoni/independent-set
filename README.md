## Golang solution for "Independent Set of Intervals"

```
Find the independent set of intervals such that its cardinality is maximum.
Input: N ranges like I1=(a1,b1[ , ... , In=(an,bn[.
Output: a set of maximum cardinality such that every interval not overlap with 
        any other.
```

For example:
Let the intervals set be: R = {(2, 5) (1, 3) (2, 7) (4, 6) (4, 8) (2, 10)}

The independent set with maximum cardinality is: In = {(1, 3) (4, 6)}

## Try yourself
Change R as you prefer in the file and run:
```bash
go run independent-set.go
```
