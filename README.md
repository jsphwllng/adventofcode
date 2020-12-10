# Advent of Code 2020 🎅🏖️
Advent of code is an annual event orchestrated by [Eric Wastl](https://twitter.com/ericwastl). The challanges have become increasingly difficult and I have used this opportunity to practice my golang skills 🎅🐹. Sometimes I have used python to quickly get the correct answer then gone back and made it work in golang too. Go is _fast_ and I will be timing each function I pass - so far the longest is `26ms` on day 4 (I blame regex).

# Progress:

| Day | Completion | Feelings |   |   |
|-----|------------|----------|---|---|
| 1   |      👍      |   🐹       |   |   |
| 2   |    👍        |     🐍     |   |   |
| 4   |        👍    |      🥴    |   |   |
| 5   |          👍  |    😎      |   |   |
| 6   |👍            |  🤔        |   |   |
| 7   |  👍          |     🤷     |   |   |
| 8   |    👍        |    😫      |   |   |
| 9   |      👍      |    🐙      |   |   |
| 10  |     ...       |      ...    |   |   |


## My takeaways from the event:
1. *a cool quick function to check `if i in range` like Python!*
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
2. *printing takes a _long_ time!*
```bash
map[c:2 e:1 j:1 k:1 l:1 m:1 n:1 q:2 r:1 x:1 y:2]
map[a:4 c:4 d:4 e:4 f:4 g:4 h:4 i:4 k:4 l:4 m:4 n:4 o:1 p:4 q:1 s:4 t:4 z:4]
3260
time taken:  119.049761ms
➜  day6 git:(master) ✗ go run day6.go
3260
time taken:  5.547017ms
```
3. *it is _marginally_ faster not to assign multiple variables on the same line* (by literally .5ms)

```go
total, answer := a, append(answer, a) //slow"
total := a 
answer := append(answer, a) //"faster"
```

4. golang isn't really a scripting language and this would be much easier in python _however_ I am stubborn and determined to learn
