# statemachine

![State machine diagram](sm.png)

```go
// Implement stateMachine.
type stateMachine struct{}

func main() {
    sm := newStateMachine()
    sm.send(1)          // "state A + 1 => state B"
    sm.send(0)          // "state B + 0 => state C"
    println(sm.state()) // "state C"
}
```
