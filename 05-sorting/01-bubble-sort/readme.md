# Bubble Sort

## Example
```
---- iteration 1 ----
[5 4 9 3]
 \ /        5 > 4 swap
[4 5 9 3]
   \ /      5 < 9 no swap
[4 5 9 3]
     \ /    9 > 3 swap

---- iteration 2 ----
[4 5 3 9]
 \ /        4 < 5 no swap
[4 5 3 9]
   \ /      5 > 3 swap

---- iteration 3 ----
[4 3 5 9]
 \ /        4 > 3 swap
[3 4 5 9]
   \ /      4 < 5 no swap

---- iteration 4 ----
[3 4 5 9]
 \ /        3 < 4 no swap
```

## Runtime
`O(n^2)`
