package main

func main() {
	addr := ":3000"
	s := NewAPIServer(addr)
	s.Run()
}
