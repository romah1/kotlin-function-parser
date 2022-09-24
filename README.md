Recursive parser of kotlin function declarations

Example input: `fun printSum(a: Int, b: Int): Unit`

Grammar:
```
Start -> fun Declaration
Declaration -> name(Arguments)
Declaraion -> name(Arguments):type
Arguments -> Variable,Arguments
Arguments -> VariableAndType
VariableAndType -> Variableariable:Type
Variable -> variable
Type -> type
```

Removing right branching:

```
Start -> fun Declaration
Declaration -> name(Arguments)Ending
Ending -> eps
Ending -> :type
Arguments -> VariableAndTypeVariableAndType-continuation
VariableAndType-continuation -> ,VariableAndType
VariableAndType-continuation -> eps
VariableAndType -> Variable:Type
Variable -> variable
Type -> type
```

First:
```
S: {fun}
N: {s}
A: {v}
V: {v}
```

Follow:
```
S: {$}
N: {$}
A: {)}
V: {,}
```