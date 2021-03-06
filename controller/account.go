package controller

import (
	"github.com/bottos-project/BlockExplorer/common"
	"github.com/bottos-project/BlockExplorer/db"
	"github.com/bottos-project/BlockExplorer/module"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func AccountList(c *gin.Context) {
	var params module.ReqAccountList
	var accounts []module.ResAccountList

	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	start := params.Start
	length := params.Length

	accountModule := mongoIns.AccountCollection()
	totalCount, _ := accountModule.Count()
	start, length = paging(start, length, totalCount)
	if length <= 0 {
		common.ResponseSuccess(c, "accounts find success", module.ResPageList{TotalRecordes: totalCount})
		return
	}

	if err := accountModule.Find(bson.M{}).Skip(start).Limit(length).All(&accounts); err != nil {
		common.ResponseErr(c, "accounts find failed", err)
		return
	}

	res := module.ResPageList{
		Data:          accounts,
		TotalRecordes: totalCount,
	}

	common.ResponseSuccess(c, "accounts find success", res)
}

func AccountDetail(c *gin.Context) {
	var params module.ReqAccountInfo
	var account module.ResAccountDetail

	if err := c.BindJSON(&params); err != nil {
		common.ResponseErr(c, "params error", err)
		return
	}

	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	accountModule := mongoIns.AccountCollection()
	if err := accountModule.Find(bson.M{"account_name": params.AccountName}).One(&account); err != nil {
		common.ResponseErr(c, "account find failed", err)
		return
	}

	// transactions
	trxModule := mongoIns.TransactionCollection()
	// transaction out count
	out := bson.M{"method": "transfer", "param.from": params.AccountName}
	send, _ := trxModule.Find(out).Count()
	// transaction in
	in := bson.M{"method": "transfer", "param.to": params.AccountName}
	receiveCount, _ := trxModule.Find(in).Count()

	sender := bson.M{"sender": params.AccountName}
	trxCount, _ := trxModule.Find(sender).Count()
	// set account info
	account.SendCount = send
	account.ReceiveCount = receiveCount
	account.TradeCount = trxCount
	res := module.ResDataStruct{
		Data: account,
	}
	common.ResponseSuccess(c, "account find success", res)
}

// 查询账户实时总数
func AccountTotal(c *gin.Context) {
	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)
	accountModule := mongoIns.AccountCollection()
	totalCount, err := accountModule.Count()
	if err != nil {
		common.ResponseErr(c, "account total search failed", err)
		return
	}

	type total struct {
		Count int `json:"rtCustCount"`
	}

	res := total{
		Count: totalCount,
	}
	common.ResponseSuccess(c, "success", res)
}
