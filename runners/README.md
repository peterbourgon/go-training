# runners

Given the runner interface, implement func race, which takes a distance and
multiple runners, and calculates the winner based on the shortest time. Then,
implement several runners.

```go
type runner interface {
	name() string
	run(distance int) (seconds int)
}

func race( ??? ) (winner string) {
	// ???
}
```

Be creative. Some runners might run at a fixed speed; some might slow down with
the distance; some might cheat and teleport?

