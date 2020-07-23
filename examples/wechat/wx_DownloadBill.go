package wechat

import (
	"fmt"

	"github.com/liao725636367/gopay"
	"github.com/liao725636367/gopay/wechat"
	"github.com/iGoogle-ink/gotil"
)

func DownloadBill() {
	//初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("sign_type", wechat.SignType_MD5)
	bm.Set("bill_date", "20190722")
	bm.Set("bill_type", "ALL")

	//请求下载对账单，成功后得到结果（string类型字符串）
	wxRsp, err := client.DownloadBill(bm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("wxRsp：", wxRsp)
}
