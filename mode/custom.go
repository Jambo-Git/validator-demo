package mode

type Struct7 struct {
	ULen         string `validate:"ulen=2"`
	UGreaterThan string `validate:"ugt=2"`
	ULessThan    string `validate:"ult=2"`
}
