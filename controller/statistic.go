package controller

import (
	"github.com/bottos-project/BlockExplorer/common"
	"github.com/bottos-project/BlockExplorer/db"
	"github.com/bottos-project/BlockExplorer/module"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func StatisticTransaction(c *gin.Context) {
	//var params module.ReqQueryStaticInfo
	var trxStatistics []module.ResStatisticTransaction

	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	sModule := mongoIns.StatisticCollection()
	count, err := sModule.Find(bson.M{}).Count()
	if err != nil {
		common.ResponseErr(c, "transaction statistic failed", err)
		return
	}

	start := count - 14
	if start < 0 {
		start = 0
	}

	if err := sModule.Find(bson.M{}).Skip(start).Limit(14).All(&trxStatistics); err != nil {
		common.ResponseErr(c, "transaction statistic failed", err)
		return
	}
	common.ResponseSuccess(c, "transaction statistic success", trxStatistics)
}

func StatisticAccount(c *gin.Context) {
	//var params module.ReqQueryStaticInfo
	var accountStatistic []module.ResStatisticAccount

	mongoIns, err := db.NewDBCollection()
	if err != nil {
		common.ResponseErr(c, "connect mongoDB error", err)
		return
	}

	defer db.CloseSession(mongoIns)

	sModule := mongoIns.StatisticCollection()
	count, err := sModule.Find(bson.M{}).Count()
	if err != nil {
		common.ResponseErr(c, "account statistic failed", err)
		return
	}

	start := count - 7
	if start < 0 {
		start = 0
	}

	if err := sModule.Find(bson.M{}).Skip(start).Limit(7).All(&accountStatistic); err != nil {
		common.ResponseErr(c, "account statistic failed", err)
		return
	}

	common.ResponseSuccess(c, "account statistic success", accountStatistic)
}
