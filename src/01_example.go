package examples

import (
	"fmt"
	"sync"
	"time"
)

func ExecuteExample01() {
	fmt.Println("Exec 01")
	now := time.Now()
	respch := make(chan string, 10)
	userId := 10
	wg := &sync.WaitGroup{}

	wg.Add(3)
	go fetchUserData(userId, respch, wg)
	go fetchUserRecommendation(userId, respch, wg)
	go fetchUserLikes(userId, respch, wg)
	wg.Wait()

	close(respch)

	for resp := range respch {
		fmt.Println(resp)
	}

	fmt.Println(time.Since(now))
}

func fetchUserData(userId int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(80 * time.Millisecond)
	respch <- "user data"
	fmt.Println("Data:", userId)
}

func fetchUserRecommendation(userId int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(120 * time.Millisecond)
	respch <- "user recommendation"
	fmt.Println("Recommendation:", userId)
}

func fetchUserLikes(userId int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(120 * time.Millisecond)
	respch <- "user likes"
	fmt.Println("Likes:", userId)
}
