package helper

var version = "1.0.0"
var Application = "golang"

func SayHello(name string) string {
	return "Hello" + name
}

// could not access from outside package
func sayGoodBye() string {
	return "goodbye"
}
