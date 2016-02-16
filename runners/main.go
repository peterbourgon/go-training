package main

type runner interface {
	name() string
	run(distance int) (seconds int)
}

func race(distance int, runners ...runner) (winner string) {
	var fastest int
	for i, runner := range runners {
		took := runner.run(distance)
		println(runner.name(), "ran", distance, "meters in", took, "seconds")
		if i == 0 || took < fastest {
			fastest = took
			winner = runner.name()
		}
	}
	return winner
}

func main() {
	var (
		r1 = baby{}
		r2 = triathlete{}
		r3 = robot{}
	)
	winner := race(100, r1, r2, r3)
	println(winner, "won the race")
}

type baby struct{}

func (baby) name() string {
	return "a baby"
}

func (baby) run(distance int) int {
	return distance * 30 // slow and steady
}

type triathlete struct{}

func (triathlete) name() string {
	return "a triathlete"
}

func (triathlete) run(distance int) int {
	var took int         // so far
	secondsPerMeter := 1 // initially fast
	for distance > 0 {
		took += secondsPerMeter
		secondsPerMeter++ // progressively slower
		distance--
	}
	return took
}

type robot struct{}

func (robot) name() string {
	return "RunBot 2000"
}

func (robot) run(distance int) int {
	return distance * 5 // pretty quick
}
