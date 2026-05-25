package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gerarID() string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%08d", rng.Intn(100000000))
}
