package mode

type Struct2 struct {
	// 不能为空，长度大于0，元素不能有空值
	Array []string `validate:"required,gt=0,dive,required"`
}
type Struct3 struct {
	// 不能为空，长度大于0，key值长度最大为5，value值不能为空并且最大长度为2
	Map map[string]string `validate:"required,gt=0,dive,keys,keymax,endkeys,required,max=2"`
}

type Struct4 struct {
	// 不能为空，内容不能存在空值
	Column1 []*Struct5 `validate:"required,dive,required"`
}

type Struct5 struct {
	// 不能为空
	Column2 string `validate:"required"`
}
