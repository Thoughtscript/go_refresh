package main

import (
	"fmt"
	"sync"
)

type Host struct {
	a, b *Philosopher
}

type Chopstick struct {
	id int
	sync.Mutex
}

type Philosopher struct {
	id, eats int
	l, r     *Chopstick
}

func eat (wgi *sync.WaitGroup, p *Philosopher) {
	defer wgi.Done()
	if p.eats < 3 {
		fmt.Printf("chopstick %d picked up \n", p.l.id)
		fmt.Printf("chopstick %d picked up \n", p.r.id)
		p.l.Lock()
		p.r.Lock()
		fmt.Printf("philosopher %d already ate %d times \n", p.id, p.eats)
		fmt.Printf("starting to eat %d (philosopher %d) \n", p.id, p.id)
		p.eats += 1
		fmt.Printf("finishing eating %d (philosopher %d) \n", p.id, p.id)
		p.l.Unlock()
		p.r.Unlock()
		fmt.Printf("chopstick %d put down \n", p.l.id)
		fmt.Printf("chopstick %d put down \n", p.r.id)
	} else {
		fmt.Printf("philosopher %d has already finished their meal \n", p.id)
	}
}

func dine(wg *sync.WaitGroup, philosophers []*Philosopher) {
	defer wg.Done()
	fmt.Printf("Beginning meal!\n")
	a := 0
	b := 2
	count := 0
	host := Host{a: philosophers[a], b: philosophers[b]}

	for {
		pa := philosophers[a]
		pb := philosophers[b]
		fmt.Printf("Philosopher %d %d selected \n", pa.id, pb.id)
		host.a = pa
		host.b = pb

		wgi := new(sync.WaitGroup)
		wgi.Add(2)
		go eat(wgi, host.a)
		go eat(wgi, host.b)
		wgi.Wait()

		a = a + 1
		if a > 4 {
			a = 0
		}

		b = b + 1
		if b > 4 {
			b = 0
		}

		count += 1
		if count == 15 {
			fmt.Printf("Meal ended!\n\n\n\n")
			fmt.Printf("Philosopher %d ate %d times \n", philosophers[0].id, philosophers[0].eats)
			fmt.Printf("Philosopher %d ate %d times \n", philosophers[1].id, philosophers[1].eats)
			fmt.Printf("Philosopher %d ate %d times \n", philosophers[2].id, philosophers[2].eats)
			fmt.Printf("Philosopher %d ate %d times \n", philosophers[3].id, philosophers[3].eats)
			fmt.Printf("Philosopher %d ate %d times \n", philosophers[4].id, philosophers[4].eats)
			break
		}
	}
}

func main() {
	// Create necessary resources
	wg := new(sync.WaitGroup)
	num := 5
	chopsticks := make([]*Chopstick, 0, num)
	for i := 0; i < num; i++ {
		chopsticks = append(chopsticks, &Chopstick{id: i+1})
	}

	philosophers := make([]*Philosopher, 0, num)
	for i := 0; i < num; i++ {
		modulovar := chopsticks[0]
		if i+1 < num - 1 {
			modulovar = chopsticks[i+1]
		}
		philosophers = append(philosophers, &Philosopher{id: i+1, l: chopsticks[i], r: modulovar, eats: 0})
	}

	//Initialization only - randomly assigned
	wg.Add(1)
	go dine(wg, philosophers)
	wg.Wait()
}
