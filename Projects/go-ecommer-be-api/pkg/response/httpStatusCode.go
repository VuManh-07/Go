package response

const (
	CodeSuccess            = 20001
	CodeFailedInvalidParam = 20003
)

var msg = map[int]string{
	CodeSuccess:            "Success",
	CodeFailedInvalidParam: "Failed",
}
