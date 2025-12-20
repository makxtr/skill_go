package main

type User struct{}

func main() {
	m := make(map[User]int)

	u := User{}

	m[u] = 2

	//fmt.Println(&m[u])
	//fmt.Println(m[&u])
}
