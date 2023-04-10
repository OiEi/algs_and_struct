package main

import (
	"algs-and-struct/graph"
	"context"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	res := graph.FindPathBetweenVert(ctx, "MOW", "KZN", 10, 5)
	log.Print(res)
}
