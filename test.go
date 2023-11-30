package main

func main() {
	n := 1
	{
		n := 2
		print(n)
	}
	print(n)
}
