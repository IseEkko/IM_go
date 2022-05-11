package service

import (
	"errors"
	"fmt"
	"httpWeb_IM/global"
	"httpWeb_IM/model"
)

type ContactService struct {
}

//添加好友服务
//传入参数：用户id，添加的id
/**
功能描述：
  对于这个功能，我们需要添加两条记录，一个是自己的还有一个是对方的，因为我们添加用户是双方的。
这里使用到了事物的回滚。
*/
func (contact *ContactService) UserContact(userid, dtid int64) (err error) {
	fmt.Println(userid == dtid)
	if userid == dtid {
		return errors.New("不能添加自己为好友")
	}
	//判断是不是已经加了好友了
	var tep = model.Contact{
		Ownerid: userid,
		Dstobj:  dtid,
		Cate:    model.CONCAT_CATE_USER,
	}

	findUserContactOk := global.DB.Where("ownerid = ? AND dstobj = ? AND cate = ?", userid, dtid, model.CONCAT_CATE_USER).Find(&tep)
	if findUserContactOk.RowsAffected > 0 {
		return errors.New("已经添加好友")
	}

	tx := global.DB.Begin()
	fmt.Println(tep)
	if creatErr := tx.Create(&tep); creatErr.Error != nil {
		tx.Rollback()
		return creatErr.Error
	}
	tep = model.Contact{
		Ownerid: dtid,
		Dstobj:  userid,
		Cate:    model.CONCAT_CATE_USER,
	}
	if creatErr := tx.Create(&tep); creatErr.Error != nil {
		tx.Rollback()
		return creatErr.Error
	}

	if errs := tx.Commit(); errs != nil {
		return errs.Error
	}
	return nil
}

//查询用户
/***
功能详情：
  在这里我们需要知道的事情，这里我们首先要查询出，这个用户有那些好友，然后查询出好友的详情信息
*/
func (contact *ContactService) SeachContactUser(userid int64) []model.User {
	contacts := make([]model.Contact, 0)
	global.DB.Where("ownerid = ? AND cate = ?", userid, model.CONCAT_CATE_USER).Find(&contacts)
	DtsId := make([]int64, 0)
	res := make([]model.User, 0)
	for _, v := range contacts {
		DtsId = append(DtsId, v.Dstobj)
	}
	if len(DtsId) == 0 {
		return res
	}
	global.DB.Where("id IN ?", DtsId).Find(&res)
	return res
}
