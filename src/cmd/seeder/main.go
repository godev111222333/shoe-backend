package main

import "github.com/godev111222333/shoe-backend/src/seeder"

func main() {
	seed := seeder.NewDataSeeder("http://0.0.0.0:9003")
	seed.SeedProducts()
}
