package request_model

import "httpWeb_IM/util"

type User_seach_Contact struct {
	util.PageArg
	Userid int64 `json:"userid" form:"userid"`
}
