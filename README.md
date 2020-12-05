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