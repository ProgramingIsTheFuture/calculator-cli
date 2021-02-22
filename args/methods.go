package args

func (a Args) SumIntNums() int32 {
	return a.SumInt.Num1 + a.SumInt.Num2
}

func (a Args) SumFloatNums() float32 {
	return a.SumFloat.Num1 + a.SumFloat.Num2
}
