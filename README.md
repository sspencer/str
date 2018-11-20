# str

Go string utils.

Starting off some chunking:

```go
    str.ChunkString("1234", 3)  -> ["123", "4"]
    str.Chunkstring("1234", -3) -> ["1", "234"]
    
    str.Comma(1234567) -> "1,234,567"
```