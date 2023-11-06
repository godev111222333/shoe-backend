package main

import "github.com/godev111222333/shoe-backend/src/seeder"

func main() {
	seed := seeder.NewDataSeeder("http://35.187.224.84:9003")
	seed.SeedProducts()
}
