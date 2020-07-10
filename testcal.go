package testcal

//Sum FUnction
func Sum(n1 int, n2 int) (int, error) {
	if n1 > 100 {
		return 0, &CusError{}
	}
	sum := n1 + n2
	return sum, nil
}

//CusError Type
type CusError struct{}

func (m *CusError) Error() string {
	return "first number cannot greater than 100"
}
