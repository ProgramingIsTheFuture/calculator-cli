package args

type SumInt struct {
	Num1 int32 `default:"0"`
	Num2 int32 `default:"0"`
}

type SumFloat struct {
	Num1 float32 `default:"0"`
	Num2 float32 `default:"0"`
}

type Args struct {
	SumInt   *SumInt   `arg:"subcommand:intsum"`
	SumFloat *SumFloat `arg:"subcommand:floatsum"`
}
