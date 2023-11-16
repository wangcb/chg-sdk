package mall

import (
	"encoding/json"
	"errors"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
	"strconv"
)

type Order struct {
}

func NewOrder(config *chg.Config) *Order {
	config.InitConfig()
	return &Order{}
}

// GetUserOrders 获取用户订单
func (r *Order) GetUserOrders(page int, prePage int) (map[string]interface{}, error) {
	req := http.Request{
		Method: "GET",
		URL:    "/internal/orders",
		Body: map[string]interface{}{
			"page":     strconv.Itoa(page),
			"pre_page": strconv.Itoa(prePage),
		},
	}
	res, err := request.Do(req, chg.Configure.MallUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Msg)
	}
	var data map[string]interface{}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetUserOrderDetail 获取用户订单详情
func (r *Order) GetUserOrderDetail(orderNo string) (map[string]interface{}, error) {
	req := http.Request{
		Method: "GET",
		URL:    "/internal/orders/" + orderNo,
		Body:   map[string]interface{}{},
	}
	res, err := request.Do(req, chg.Configure.MallUrl)
	if err != nil {
		return nil, err
	}
	if res.Code != 200 {
		return nil, errors.New(res.Msg)
	}
	var data map[string]interface{}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
