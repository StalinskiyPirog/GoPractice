package fibonacci

import "fmt"

// fibo - function that takes integer value - n and as a result return Fibonacci number with index n
func fibo(n int) {
	if n*10%10 != 0{
	panic("Input argument is not an integer. Please make sure to input integer value")
}else if n < 0{
	panic("Input argument less than zero. Please make sure to input integer value")
}else if n == 0 {
	return 0
} else if n == 1{
	return 1
} else {
	return fibo(n-1) + fibo(n-2)
}
}

func main(){
	fmt.Println(fibo(10))
}

