Recursive parser of kotlin function declarations

Example input: `fun printSum(a: Int, b: Int): Unit`

Grammar:
```
Start -> fun Declaration
Declaration -> FunctionName(Arguments)
Declaraion -> FunctionName(Arguments):Type
Arguments -> Variable,Arguments | eps
Arguments -> VariableAndType
VariableAndType -> Variableariable:Type
Variable -> variable
Type -> TypeName | TypeNameTypeMark
TypeMark = ? | ! | eps
```

Removing right branching:

```
Start -> fun Declaration
Declaration -> FunctionName(Arguments)Ending
FunctionName -> name
Ending -> :Type | eps
Arguments -> VariableAndTypeVariableAndType-continuation | eps
VariableAndType-continuation -> ,VariableAndTypeVariableAndType-continuation | eps
VariableAndType -> Variable:Type
Variable -> variable
Type -> TypeName | TypeNameTypeMark
TypeMark = ? | ! | eps
```

First:
```
Start: {fun}
Declaration: {name}
FunctionName: {name}
Ending: {:, eps}
Arguments: {variable, eps}
VariableAndType-continuation: {',', eps}
VariableAndType: {variable}
Variable: {variable}
Type: {type}
TypeName: {type}
TypeMark: {!, ?}
```

Follow:
```
Start: {$}
Declaration: {$}
FunctionName: {(}
Ending: {$}
Arguments: {)}
VariableAndType-continuation: {)}
VariableAndType: {',', )}
Variable: {:}
Type: {$, )}
TypeName: {!, ?}
TypeMark: {$, )}
```