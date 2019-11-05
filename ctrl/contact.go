/*
@desc : Created by San on 2019/11/3 23:58
*/
package ctrl

import (
	"chatting/args"
	"chatting/service"
	"chatting/util"
	"net/http"
)

var contactService service.ContactService

func LoadFriend(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg
	util.Bind(req, &arg)

	users := contactService.SearchFriend(arg.Userid)
	util.RespOKList(w, users, len(users))
}

func LoadCommunity(w http.ResponseWriter, req *http.Request) {
	var arg args.ContactArg

	util.Bind(req, &arg)
	communities := contactService.SearchCommunity(arg.Userid)
	util.RespOKList(w, communities, len(communities))
}
