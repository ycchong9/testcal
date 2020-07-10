package testcal

//Sum FUnction
func Sum(n1 int, n2 int) (int, error) {
	if n1 > 100 {
		return 0, &CusError{
			Msg: "first number cannot greater than 100",
		}
	}
	if n2 > 200 {
		return 0, &CusError{
			Msg: "second number cannot greater than 200",
		}
	}
	sum := n1 + n2
	return sum, nil
}

//CusError Type
type CusError struct {
	Msg string
}

func (m *CusError) Error() string {
	return m.Msg
}
