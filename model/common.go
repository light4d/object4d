package model

type CommonResp struct {
	Error  interface{}
	Code   int
	Result interface{}
}

type Err struct {
	Detail string
	Data   interface{}
}

func (e Err) Error() string {
	return e.Detail
}

func NewErr(detail string) error {
	return &Err{Detail: detail}
}

func NewErrData(detail string, data interface{}) error {
	return &Err{Detail: detail, Data: data}
}

var FieldCannotupdate = "there fields can't updated"
var ErrLenBigThan1 = NewErr("len(*)>1")
var ErrLenNotEqual1 = NewErr("len(*)!=1")
var ErrLenEqual0 = NewErr("len(*)==0")
