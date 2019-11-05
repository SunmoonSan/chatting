/*
@desc : Created by San on 2019/11/4 00:26
*/
package args

type ContactArg struct {
	PageArg
	Userid int64 	`json:"userid" form:"userid"`
	Dstid int64 	`json:"dstid" form:"dstid"`
}