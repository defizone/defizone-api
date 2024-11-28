package main

func main() {
	server := NewAPIServer(":8000")
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
