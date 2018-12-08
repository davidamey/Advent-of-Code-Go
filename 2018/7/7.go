/** Not the most concise **/

package main

import (
	"advent/util"
	"fmt"
)

func main() {
	file, _ := util.OpenInput()
	defer file.Close()
	lines, _ := util.ReadLines(file)

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	steps := make(StepMap)
	for _, l := range lines {
		var p, s rune
		fmt.Sscanf(l, "Step %c must be finished before step %c can begin.", &p, &s)

		steps.AddPre(s, p)
		steps.CreateStep(p)
	}

	fmt.Println("== part1 ==")
	for {
		next := '`'
		for s, p := range steps {
			if len(p) == 0 && s < next {
				next = s
			}
		}

		if next == '`' {
			break
		}

		fmt.Printf("%c", next)
		steps.CompleteStep(next)
	}
	fmt.Println()
}

func part2(lines []string) {
	steps := make(StepMap)
	for _, l := range lines {
		var p, s rune
		fmt.Sscanf(l, "Step %c must be finished before step %c can begin.", &p, &s)

		steps.AddPre(s, p)
		steps.CreateStep(p)
	}

	fmt.Println("== part2 ==")
	wp := newWorkerPool(5, steps)
	for len(steps) > 0 || wp.Working() {
		for wp.WorkerFree() {
			next := '`'
			for s, p := range steps {
				if len(p) == 0 && s < next {
					next = s
				}
			}

			if next == '`' {
				break
			}

			if wp.StartJob(next) {
				steps.StartStep(next)
			}
		}

		wp.Tick()
	}
}

type StepMap map[rune]map[rune]bool

func (sm StepMap) CreateStep(step rune) {
	if sm[step] == nil {
		sm[step] = make(map[rune]bool)
	}
}

func (sm StepMap) AddPre(step, pre rune) {
	sm.CreateStep(step)
	sm[step][pre] = true
}

func (sm StepMap) CompleteStep(step rune) {
	// Remove step from all pre's
	for s := range sm {
		delete(sm[s], step)
	}

	// Remove the step itself
	delete(sm, step)
}

func (sm StepMap) StartStep(step rune) {
	delete(sm, step)
}

type worker struct {
	Usage int
	Task  rune
	Total int
}

type workerPool struct {
	workers []worker
	steps   StepMap
}

func newWorkerPool(num int, steps StepMap) *workerPool {
	return &workerPool{
		make([]worker, num),
		steps,
	}
}

func (wp *workerPool) Working() bool {
	for _, w := range wp.workers {
		if w.Usage > 0 {
			return true
		}
	}
	return false
}

func (wp *workerPool) WorkerFree() bool {
	for _, w := range wp.workers {
		if w.Usage == 0 {
			return true
		}
	}
	return false
}

func (wp *workerPool) Elapsed() int {
	t := 0
	for _, w := range wp.workers {
		if w.Total > t {
			t = w.Total
		}
	}
	return t
}

func (wp *workerPool) StartJob(task rune) bool {
	for i, w := range wp.workers {
		if w.Usage == 0 {
			d := 60 + int(task-64)
			// fmt.Printf("worker %d starting task %c (%ds)\n", i, task, d)
			wp.workers[i].Total += d
			wp.workers[i].Usage += d
			wp.workers[i].Task = task
			return true
		}
	}
	return false
}

var second int = 0
var completed []rune = make([]rune, 26)

func (wp *workerPool) Tick() {
	fmt.Printf("%3d", second)
	second++
	for i, w := range wp.workers {
		if w.Usage > 0 {
			wp.workers[i].Usage--
			fmt.Printf("%3c", w.Task)
			if wp.workers[i].Usage == 0 {
				// fmt.Printf("worker %d finished task %c\n", i, w.Task)
				wp.steps.CompleteStep(w.Task)
				wp.workers[i].Task = '`'
				completed = append(completed, w.Task)
			}
		} else {
			fmt.Printf("%3c", '.')
		}
	}
	fmt.Print("  ")
	for _, c := range completed {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
