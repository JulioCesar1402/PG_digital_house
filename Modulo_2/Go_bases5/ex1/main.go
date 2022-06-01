package main

type MyError struct {
	Msg string
}

func (e *MyError) Error() string {
	return fmt
}

func verifySalary(salary int) (int, error){
	if salary < 15000 {
		return MyError
	}
}

func main()  {
	var salary int
}