package main
import "os"
import "fmt"
import "simplemath"
import "strconv"

var Usage = func() {
	fmt.Println("USAGE: calc command [arguments] ...")
	fmt.Println("\nThe commands are:\nadd\tAddition of two value.\nsqrt\tSquare root of non-negative value.")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 3 {
		Usage()
		return
	}

	switch args[1] {
		case "add":
			if  len(args) != 4 {
				fmt.Println("Usage: cale add <integer1> <integer2>")
				return
			}

			v1, err1 := strconv.Atoi(args[2])
			v2, err2 := strconv.Atoi(args[3])
			if err1 != nil || err2 != nil {
				fmt.Println("Usage: cale add <integer1> <integer2>")
				return
			}

			ret :=simplemath.Add(v1, v2)
			fmt.Println("Result: ", ret)
		case "sqrt":
			if len(args) != 3 {
				fmt.Println("Usage: calc sqrt <integer>")
				return
			}

			v, err := strconv.Atoi(args[2])
			if err != nil{
				fmt.Println("Usage: calc sqrt <integer>")
				return
			}

			ret := simplemath.Sqrt(v)
			fmt.Println("Result: ", ret)
		default:
			fmt.Println("unknown command ", args[1])
			Usage()
	}
}

