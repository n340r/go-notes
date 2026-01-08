# What i learned

1. Recursion training. Very common tree traversal task
2. Very important line

```go
var _ Directory = (*DefaultDirectory)(nil)
```

In-depth value/pointer receivers. Runtime sugar and compile-time interfaces.  
Difference between line above from

```go
var _ Directory = DefaultDirectory{}
```
