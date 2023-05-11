package response

type InvoiceVAT struct {
	NoteCode       string        `json:"noteCode"`       // 发票代码
	NoteNo         string        `json:"noteNo"`         // 发票号码
	CheckCode      string        `json:"checkCode"`      // 校验码
	BillingDate    string        `json:"billingDate"`    // 开票日期
	PurchaserName  string        `json:"purchaserName"`  // 购买方名称
	PurchaserTaxNo string        `json:"purchaserTaxNo"` // 购买方税号
	SellerName     string        `json:"sellerName"`     // 销售方名称
	SellerTaxNo    string        `json:"sellerTaxNo"`    // 销售方税号
	InvoiceType    string        `json:"invoiceType"`    // 发票类型
	InvoiceItems   []InvoiceItem `json:"invoiceItems"`   // 发票明细列表
	TotalTaxAmount float64       `json:"totalTaxAmount"` // 税额总额
	TotalAmount    float64       `json:"totalAmount"`    // 价税合计总额
	OriginResult   string        `json:"origin_result"`
}

type InvoiceItem struct {
	Name          string  `json:"name"`          // 商品名称
	Specification string  `json:"specification"` // 商品规格
	Unit          string  `json:"unit"`          // 商品单位
	Quantity      float64 `json:"quantity"`      // 商品数量
	Price         float64 `json:"price"`         // 商品单价
	Amount        float64 `json:"amount"`        // 商品金额
	TaxRate       float64 `json:"taxRate"`       // 商品税率
	TaxAmount     float64 `json:"taxAmount"`     // 商品税额
}
