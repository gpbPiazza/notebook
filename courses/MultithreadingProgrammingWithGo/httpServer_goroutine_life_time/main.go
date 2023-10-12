package main

import (
	"fmt"
	"net/http"
	"time"

	"math/rand"

	"github.com/gofiber/fiber/v2"
)

func onlyReturndOods(odds chan *int) {
	zeroValue := 0
	var number *int
	number = &zeroValue
	fmt.Println("GOROUTINE RUNNING")
	// time.Sleep(3 * time.Second)
	// fmt.Println("GOROUTINE RUNNING AND STOPED SLEEPING")
	randInt := rand.Intn(100)

	if (randInt % 2) == 0 {
		number = &randInt
		odds <- number
		return // Se tirar esse return da panic pois, a função pai fecha o canal antes de termianr a execução da rotina filha
	}
	odds <- number
}

func doSomething() {
	type guardian struct {
		Number *int
	}
	zapper := guardian{}
	oddsChan := make(chan *int)
	fmt.Println("BEFORE CALL GOROUTINE", "CHANNEL IS NEW VALUE?", oddsChan)
	go onlyReturndOods(oddsChan)
	fmt.Println("BEFORE ACCESS VALUE FROM CHANNEL")
	result, ok := <-oddsChan
	fmt.Println("AFTER ACCESS VALUE FROM CHANNEL")
	if ok {
		zapper.Number = result
		fmt.Println("channel open", ok, *result)
	}
	close(oddsChan)

	// DEPOIS DE FECHAR UM CANAL QUE O VALOR DO MESMO É UM PONTEIRO O CANAL "FECHA O PONTEIRO TAMBÉM"
	// TORNANDO O VALOR DO CANAL NIL! Pois o canal está fechado
	// TUDO VAI PRO CARALHO DEPOIS DE FECHAR O CANAL E TENTAR ACESSAR ELE DENOVO. Valores e ponteiro antes disso
	// Se mantem OK
	result, ok = <-oddsChan
	chanisNil := oddsChan == nil
	fmt.Println("guardian number is nil after close the channel?", *zapper.Number)
	fmt.Println("channel IS CLOSED ->", !ok, "AND RESULT IS NIL ->", result, "chan is nil?", chanisNil)
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
