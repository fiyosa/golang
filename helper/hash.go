package helper

import "github.com/speps/go-hashids/v2"

const secretHashId = "secret"

func EncodeId(data int) string {
	hd := hashids.NewData()
	hd.Salt = secretHashId
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)
	value, err := h.Encode([]int{data})
	if err != nil {
		return ""
	}
	return value
}

func DecodeId(data string) int {
	hd := hashids.NewData()
	hd.Salt = secretHashId
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)
	value, err := h.DecodeWithError(data)
	if err != nil {
		return -1
	}
	return value[0]
}
