Recursive parser of kotlin function declarations

Example input: `fun printSum(a: Int, b: Int): Unit`

Grammar:
```
S -> fun N
N -> s(A)
N -> s(A): t
A -> V,
A -> V
V -> v: t
```