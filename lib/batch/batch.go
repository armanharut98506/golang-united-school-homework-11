package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func worker(id int64, jobs <-chan int64, results chan<- user) {
	for j := range jobs {
		results <- getOne(j)
	}
}

func getBatch(n int64, pool int64) (res []user) {
	jobs := make(chan int64, n)
	results := make(chan user, n)

	for i := int64(0); i < pool; i++ {
		go worker(i, jobs, results)
	}

	for j := int64(0); j < n; j++ {
		jobs <- j
	}
	close(jobs)

	users := make([]user, 0, n)
	for i := int64(0); i < n; i++ {
		users = append(users, <-results)
	}
	
	return users
}