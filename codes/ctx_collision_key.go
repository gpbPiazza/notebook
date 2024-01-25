package main

import (
	"context"
	"fmt"
	"log"
)

type ctxKey struct{}

type MyConnectionPool struct {
	Number int
	DbName string
}

var gPool MyConnectionPool

func SetConnCtx(ctx context.Context, cPool *MyConnectionPool) context.Context {
	if cPool != nil {
		return context.WithValue(ctx, ctxKey{}, cPool)
	}

	return context.WithValue(ctx, ctxKey{}, gPool)
}

func Conn(ctx context.Context) *MyConnectionPool {
	cPool, ok := ctx.Value(ctxKey{}).(*MyConnectionPool)

	if !ok {
		log.Fatal("ctx has not connection pool with db")
	}

	return cPool
}

// Aprendizado, quando você sobreescreve uma chave no ctx ja existente
// quando você consultar essa chave você sempre irá receber o último valor setada, como uma pilha.
func main() {
	ctx := context.Background()

	for i := 0; i < 15; i++ {
		gremio := new(MyConnectionPool)
		gremio.DbName = fmt.Sprintf("DD%d", i)
		gremio.Number = i
		ctx = SetConnCtx(ctx, gremio)
		meuGremio := Conn(ctx)

		if meuGremio.DbName != gremio.DbName {
			fmt.Println("OPA NÃO FUNCINOU COMO DIZ NO COMENTÁRIO")
		}
		if meuGremio.Number != gremio.Number {
			fmt.Println("OPA NÃO FUNCINOU COMO DIZ NO COMENTÁRIO")
		}
	}

	fmt.Println("OPA FUNCINOU COMO DIZ NO COMENTÁRIO")
}
