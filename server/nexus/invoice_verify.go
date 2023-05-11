package nexus

import (
	"encoding/json"
	"github.com/wangcb/chg-sdk/chg"
	"github.com/wangcb/chg-sdk/http"
	"github.com/wangcb/chg-sdk/request"
	"github.com/wangcb/chg-sdk/response"
)

type InvoiceVerify struct {
}

func NewInvoiceVerify(config *chg.Config) *InvoiceVerify {
	config.InitConfig()
	return &InvoiceVerify{}
}

func (t *InvoiceVerify) Medical(url string) (*response.InvoiceMedical, error) {
	req := http.Request{
		Method: "POST",
		URL:    "/api/invoice/medical",
		Body: map[string]interface{}{
			"url": url,
		},
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return nil, err
	}
	var data response.InvoiceMedical
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	return &data, nil
}

func (t *InvoiceVerify) VAT(url string) (*response.InvoiceVAT, error) {
	req := http.Request{
		Method: "POST",
		URL:    "/api/invoice/vat",
		Body: map[string]interface{}{
			"url": url,
		},
	}
	res, err := request.Do(req, chg.Configure.NexusUrl)
	if err != nil {
		return nil, err
	}
	var data response.InvoiceVAT
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	return &data, nil
}
