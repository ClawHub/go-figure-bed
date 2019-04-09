package utils

import (
	"github.com/Unknwon/com"
	"github.com/sony/sonyflake"
)

//唯一ID
func GeneralId() string {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, _ := flake.NextID()
	return com.ToStr(id)
}
