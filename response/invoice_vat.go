package response

type InvoiceVAT struct {
	NoteCode       string        `json:"note_code"`        // 发票代码
	NoteNo         string        `json:"note_no"`          // 发票号码
	CheckCode      string        `json:"check_code"`       // 校验码
	BillingDate    string        `json:"billing_date"`     // 开票日期
	PurchaserName  string        `json:"purchaser_name"`   // 购买方名称
	PurchaserTaxNo string        `json:"purchaser_tax_no"` // 购买方税号
	SellerName     string        `json:"seller_name"`      // 销售方名称
	SellerTaxNo    string        `json:"seller_tax_no"`    // 销售方税号
	InvoiceType    string        `json:"invoice_type"`     // 发票类型
	InvoiceItems   []InvoiceItem `json:"invoice_items"`    // 发票明细列表
	TotalTaxAmount float64       `json:"total_tax_amount"` // 税额总额
	InvoiceAmount  float64       `json:"invoice_amount"`   // 发票金额
	TotalAmount    float64       `json:"total_amount"`     // 总金额(发票金额+税额)
	OriginResult   string        `json:"origin_result"`
}

type InvoiceItem struct {
	Name          string  `json:"name"`          // 商品名称
	Specification string  `json:"specification"` // 商品规格
	Unit          string  `json:"unit"`          // 商品单位
	Quantity      float64 `json:"quantity"`      // 商品数量
	Price         float64 `json:"price"`         // 商品单价
	Amount        float64 `json:"amount"`        // 商品金额
	TaxRate       float64 `json:"tax_rate"`      // 商品税率
	TaxAmount     float64 `json:"tax_amount"`    // 商品税额
}
