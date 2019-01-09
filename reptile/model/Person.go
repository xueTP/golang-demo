package model

type BaseData struct {
	Id string
	Url string
}

type Person struct {
	BaseData
	UserName      string
	IsMarred      bool
	Age           int
	Constellation string
	Height        int
	Weight        int
	WorkAddress   string
	Income        string
	Job           string
	Education     string
	Content       string
	OtherDetail   string
}
