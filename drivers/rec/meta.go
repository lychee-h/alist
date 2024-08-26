package _rec

import (
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/op"
)

type Addition struct {
	// Usually one of two
	// driver.RootPath
	// driver.RootID
	// define other
	Username    string `json:"username" required:"true"`
	Password    string `json:"password" required:"true"`
	ResultInput string `json:"resultInput" required:"true"`
	GroupNumber string `json:"group_number" required:"true"`
	// Cookie   string `json:"cookie" help:"Fill in the cookie if need captcha"`
	driver.RootID
	// Field string `json:"field" type:"select" required:"true" options:"a,b,c" default:"a"`
}

var config = driver.Config{
	// Name:              "Template",
	// LocalSort:         false,
	// OnlyLocal:         false,
	// OnlyProxy:         false,
	// NoCache:           false,
	// NoUpload:          false,
	// NeedMs:            false,
	// DefaultRoot:       "root, / or other",
	// CheckStatus:       false,
	// Alert:             "",
	// NoOverwriteUpload: false,
	Name:        "RecCloud",
	LocalSort:   true,
	DefaultRoot: "0", // 群盘根目录
	Alert:       `info|check groupId.`,
}

func init() {
	op.RegisterDriver(func() driver.Driver {
		return &RecCloud{}
	})
}
