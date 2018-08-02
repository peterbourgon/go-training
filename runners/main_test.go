package main

import "testing"

func TestBaby(t *testing.T) {
	for distance, want := range map[int]int{
		1: 30 * 1,
		2: 30 * 2,
		3: 30 * 3,
	} {
		have := baby{}.run(distance)
		if want != have {
			t.Errorf("run(%d): want %d, have %d", distance, want, have)
		}
	}
}

func TestTriathlete(t *testing.T) {
	for distance, want := range map[int]int{
		1: 1,
		2: 1 + 2,
		3: 1 + 2 + 3,
	} {
		have := triathlete{}.run(distance)
		if want != have {
			t.Errorf("run(%d): want %d, have %d", distance, want, have)
		}
	}
}

func TestRobot(t *testing.T) {
	for distance, want := range map[int]int{
		1: 5 * 1,
		2: 5 * 2,
		3: 5 * 3,
	} {
		have := robot{}.run(distance)
		if want != have {
			t.Errorf("run(%d): want %d, have %d", distance, want, have)
		}
	}
}

func TestRace(t *testing.T) {
	for i, testcase := range []struct {
		distance int
		runners  []runner
		want     string
	}{
		{40, []runner{baby{}}, "a baby"},
		{50, []runner{baby{}, triathlete{}}, "a triathlete"},
		{60, []runner{baby{}, triathlete{}}, "a baby"}, // lol
	} {
		if want, have := testcase.want, race(testcase.distance, testcase.runners...); want != have {
			t.Errorf("race %d: want %s, have %s", i+1, want, have)
		}
	}
}

func BenchmarkBaby1M(b *testing.B) {
	benchmark(b, baby{}, 1)
}

func BenchmarkBaby10M(b *testing.B) {
	benchmark(b, baby{}, 10)
}

func BenchmarkBaby100M(b *testing.B) {
	benchmark(b, baby{}, 100)
}

func BenchmarkTriathlete1M(b *testing.B) {
	benchmark(b, triathlete{}, 1)
}

func BenchmarkTriathlete10M(b *testing.B) {
	benchmark(b, triathlete{}, 10)
}

func BenchmarkTriathlete100M(b *testing.B) {
	benchmark(b, triathlete{}, 100)
}

func BenchmarkRobot1M(b *testing.B) {
	benchmark(b, robot{}, 1)
}

func BenchmarkRobot10M(b *testing.B) {
	benchmark(b, robot{}, 10)
}

func BenchmarkRobot100M(b *testing.B) {
	benchmark(b, robot{}, 100)
}

func benchmark(b *testing.B, r runner, distance int) {
	for i := 0; i < b.N; i++ {
		r.run(distance)
	}
}
