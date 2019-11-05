/*
@desc : Created by San on 2019/11/3 23:59
*/
package service

import (
	"chatting/model"
	"errors"
	"time"
)

type ContactService struct {
}

func (service *ContactService) AddFriend(userid, dstid int64) error {
	if userid == dstid {
		return errors.New("不能添加自己为好友呀")
	}

	//	判断是否已经加了好友
	tmp := model.Contact{}
	// 查询是否已经是好友
	DbEngin.Where("ownerid = ?", userid).And("dstid = ?", dstid).And("cate = ?", model.CONCAT_CATE_USER).Get(&tmp)

	if tmp.Id > 0 {
		return errors.New("该用户已经被添加过了")
	}

	// 事务
	session := DbEngin.NewSession()
	session.Begin()
	_, e2 := session.InsertOne(model.Contact{
		Ownerid:  userid,
		Dstobj:   dstid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})

	// 插入对方的
	_, e3 := session.InsertOne(model.Contact{
		Ownerid:  dstid,
		Dstobj:   userid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})

	// 没有错误
	if e2 == nil && e3 == nil {
		session.Commit()
		return nil
	} else {
		session.Rollback()
		if e2 != nil {
			return e2
		} else {
			return e3
		}
	}
}

// 查找群
func (service *ContactService) SearchCommunity(userId int64) []model.Community {
	contacts := make([]model.Contact, 0)
	comIds := make([]int64, 0)

	DbEngin.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&contacts)
	for _, v := range contacts {
		comIds = append(comIds, v.Dstobj)
	}

	coms := make([]model.Community, 0)
	if len(comIds) == 0 {
		return coms
	}

	DbEngin.In("id", comIds).Find(&coms)
	return coms

}

// 查找好友
func (service *ContactService) SearchFriend(userId int64) []model.User {
	contacts := make([]model.Contact, 0)
	objIds := make([]int64, 0)
	DbEngin.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_USER).Find(&contacts)
	for _, v := range contacts {
		objIds = append(objIds, v.Dstobj)
	}

	coms := make([]model.User, 0)
	if len(objIds) == 0 {
		return coms
	}

	DbEngin.In("id", objIds).Find(&coms)
	return coms
}
