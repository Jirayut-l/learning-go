#### Basic Concept of Slicing

The syntax for slicing is `slice[start:end]`, where:

- `start` is the index where the slice begins (รวมอยู่ด้วยเช่นเริ่มที่ 1 ) .
- `end` is the index where the slice ends (ไม่รวมอยู่ด้วยเช่น 1:3 จะได้จริง ๆแค่ index ที่ 1,2).
  *If `start` is omitted, it defaults to 0. If `end` is omitted, it defaults to the length of the slice.*

#### Syntax Slicing

```go
a[low : high] 
or
a[start:end]
```
