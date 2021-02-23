package args

// This will sum the 2 integers numbers
func (a Args) SumIntNums() int32 {
	return a.SumInt.Num1 + a.SumInt.Num2
}

// This will subtract the 2 integers numbers
func (a Args) SubIntNums() int32 {
	return a.SubInt.Num1 - a.SubInt.Num2
}

// This will sum the 2 float numbers
func (a Args) SumFloatNums() float32 {
	return a.SumFloat.Num1 + a.SumFloat.Num2
}

// This will subtract the 2 float numbers
func (a Args) SubFloatNums() float32 {
	return a.SubFloat.Num1 - a.SubFloat.Num2
}
