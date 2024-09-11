### Full Slice Expression Syntax

```go
newSlice := originalSlice[start:end:capacity]
```

- `start` is the index where the new slice begins (inclusive).
- `end` is the index where the new slice ends (exclusive).
- `capacity` is the maximum capacity of the new slice, which cannot exceed the capacity of the original slice minus the `start` index.
