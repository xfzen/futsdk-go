package futsdk

import (
	"context"
	"testing"

	futsdk "github.com/xfzen/futsdk/futuapi"

	"github.com/zeromicro/go-zero/core/logx"
)

func init() {
	opend = futsdk.NewFutuAPI()
	err := opend.Connect(context.TODO(), ":11111")
	if err != nil {
		logx.Errorf("opend.Connect err: %v", err)
		return
	}

	logx.SetUp(logx.LogConf{Mode: "console", Level: "debug", Encoding: "plain"})
}

func TestGetUserSecurity(t *testing.T) {
	GetUserSecurity("美股")
}

func TestFetchSZHistoryKLine(t *testing.T) {
	FetchSZHistoryKLine("601238", "2022-09-01", "2023-01-12")
}

func TestFetchUSHistoryKLine(t *testing.T) {
	// 正常美股代码
	FetchUSHistoryKLine("AAPL", "2022-09-01", "2023-01-12")
	// 异常代码
	FetchUSHistoryKLine("INVALID", "2022-09-01", "2023-01-12")
	// 异常时间区间
	FetchUSHistoryKLine("AAPL", "2023-01-12", "2022-09-01")
}

func TestFetchHKHistoryKLine(t *testing.T) {
	// 正常港股代码
	FetchHKHistoryKLine("00700", "2022-09-01", "2023-01-12")
	// 异常代码
	FetchHKHistoryKLine("INVALID", "2022-09-01", "2023-01-12")
	// 异常时间区间
	FetchHKHistoryKLine("00700", "2023-01-12", "2022-09-01")
}
