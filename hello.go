package main
import "fmt"

func Create(x int) func (int)int{

	x=10;
	return func(y int)int{
	x+=y;
	return x;
	}
}




func main(){
	
	xxx:=111;
	func(){
	fmt.Println(xxx);
	}();
	f:=Create(2);
	fmt.Println(f(2));
	fmt.Println(f(3));
	g:=Create(3);
	fmt.Println(g(2));
	fmt.Println(g(3));
}
