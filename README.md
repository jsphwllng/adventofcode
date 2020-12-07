# Cool things I have learned doing AOC2020

*a cool quick function to check `if i in range` like Python!*
```go
func contains(arr []int, seat int) bool {
	for _, a := range arr {
		if a == seat {
			return true
		}
	}
	return false
}
```
*printing takes a _long_ time!*
```bash
map[c:2 e:1 j:1 k:1 l:1 m:1 n:1 q:2 r:1 x:1 y:2]
map[a:4 c:4 d:4 e:4 f:4 g:4 h:4 i:4 k:4 l:4 m:4 n:4 o:1 p:4 q:1 s:4 t:4 z:4]
3260
time taken:  119.049761ms
➜  day6 git:(master) ✗ go run day6.go
6457
time taken:  2.49411ms
3260
time taken:  5.547017ms
```

```go
          *
        .###.
       .#%##%.
      .%##%###.
     .##%###%##. 
    .#%###%##%##.
          # 
          #
	  
```
