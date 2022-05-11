package request_model

import "httpWeb_IM/util"

type ContactArg struct {
	util.PageArg
	Userid int64 `json:"userid" form:"userid"`
	Dstid  int64 `json:"dstid" form:"dstid"`
}
