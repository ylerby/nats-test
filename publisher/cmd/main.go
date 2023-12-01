package main

import "publisher"

func main() {
	pub := publisher.New("test-cluster", "client-2")
	pub.PublicMessage()
}
