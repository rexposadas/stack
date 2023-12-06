package main

func whatType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

func main() {

	// Answers, how can we know what time is being passed in?

	list := []interface{}{1, "hello", true, `raw`}

	for _, v := range list {
		println(whatType(v))
	}

}
