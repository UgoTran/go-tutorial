import "fmt"

func failedUpdate(px *int) {
	// Next, we declare x2 in failedUpdate, set to 20
	x2 := 20
	// We then point px in failedUpdate to the address of x2
	px = &x2
}

func update(px *int) {
	//this time we change the value of what px in update points to, the variable x in main
	*px = 20
}

func main() {
	x := 10
	//When we call failedUpdate, we copy the address of x into the parameter px
	failedUpdate(&x)
	fmt.Println(x) // prints 10;  the value of x is unchanged

	//When we call update, we copy the address of x into px again
	update(&x)
	fmt.Println(x) // prints 20
}
