package main

// var a = struct {
// 	a int
// 	b string
// }{
// 	a: 1,
// 	b: "2",
// }
// println(a.a)
// println(a.b)
		
type Person struct {
	name string
	age  int
	lastname string
	username string
}
	func main() {
		var p Person
		p.name = "Luis"
		p.age = 17
		p.lastname = "Alvarez"
		p.username = "mrLuisFer"

		println(p.name)
		println(p.age)
		println(p.lastname)
		println(p.username)
	}
