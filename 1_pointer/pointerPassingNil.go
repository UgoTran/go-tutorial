import "fmt"

func failedUpdate(g *int) {
	// Next, we change g in failedUpdate to point to x
	x := 10
	g = &x
	// This does not change the f in main, and when we exit failedUpdate and return to main, f is still nil
	fmt.Println("g: ", g)
}
/* Example passing pointer nil to Function */

func main() {
	var f *int // f is nil
	//when you pass a nil pointer to a function, you can't make the value non-nil
	// When we call failedUpdate, we copy the value of f, which is nil, into the parameter named g
	failedUpdate(f) //passing nil to f
	fmt.Println(f) // prints nil
}
