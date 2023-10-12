package main

import (
	"fmt"
	"net/http"
	"time"

	"math/rand"

	"github.com/gofiber/fiber/v2"
)

func onlyReturndOods(odds chan int) {
	fmt.Println("GOROUTINE RUNNING AND SLEEPING")
	// time.Sleep(3 * time.Second)
	fmt.Println("GOROUTINE RUNNING AND STOPED SLEEPING")
	number := rand.Intn(100)

	if (number % 2) == 0 {
		odds <- number
		return // Se tirar esse return da panic pois, a função pai fecha o canal antes de termianr a execução da rotina filha
	}
	odds <- 0
}

func doSomething() {
	oddsChan := make(chan int)
	fmt.Println("BEFORE CALL GOROUTINE")
	go onlyReturndOods(oddsChan)
	fmt.Println("BEFORE ACCESS VALUE FROM CHANNEL")
	result, ok := <-oddsChan
	fmt.Println("AFTER ACCESS VALUE FROM CHANNEL")
	if ok {
		fmt.Println("channel open", ok, result)
	}
	close(oddsChan)
	result, ok = <-oddsChan
	chanisNil := oddsChan == nil
	fmt.Println("channel open", ok, result, "chan is nil?", chanisNil)
}

func countToTen() chan bool {
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println(i)
		}
		done <- true
	}()
	return done
}

func coutToTenFather() {
	done := countToTen()
	fmt.Println("countToTen() exited")
	<-done // block until countToTen()'s goroutine is done

	// note that we need to block the main thread until countToTen()'s goroutine completes.
	// If we don't do this, the main thread will exit and all other goroutines will
	// be stopped even if they haven't completed their task yet.
}

func main() {
	app := fiber.New()

	app.Get("salve", func(c *fiber.Ctx) error {
		doSomething()
		return c.Status(http.StatusOK).SendString("finish calculating odds")
	})

	app.Listen(":8080")
}
