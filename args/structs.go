package args

type Ints struct {
	Num1 int32 `default:"0"`
	Num2 int32 `default:"0"`
}

type Floats struct {
	Num1 float32 `default:"0"`
	Num2 float32 `default:"0"`
}

type Args struct {
	SumInt   *Ints   `arg:"subcommand:intsum"`
	SubInt   *Ints   `arg:"subcommand:intsub"`
	SumFloat *Floats `arg:"subcommand:floatsum"`
	SubFloat *Floats `arg:"subcommand:floatsub"`
}
