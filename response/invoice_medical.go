package response

type InvoiceMedical struct {
	NoteCode                string         `json:"note_code"`                  // 发票代码
	NoteNo                  string         `json:"note_no"`                    // 发票号码
	BillingDate             string         `json:"billing_date"`               // 开票日期
	CheckCode               string         `json:"check_code"`                 // 校验码
	ChargingUnits           string         `json:"charging_units"`             // 收费单位
	Checksum                Checksum       `json:"checksum"`                   // 校验和
	CostCategories          []CostCategory `json:"cost_categories"`            // 费用分类列表
	CostDetailList          []CostDetail   `json:"cost_detail_list"`           // 费用明细列表
	HospitalName            string         `json:"hospital_name"`              // 医院名称
	HospitalNo              string         `json:"hospital_no"`                // 医院编号
	InvoicePartyCode        string         `json:"invoice_party_code"`         // 发票方代码
	MedicalInsuranceNo      string         `json:"medical_insurance_no"`       // 医疗保险号码
	MedicalInsuranceType    string         `json:"medical_insuranceType"`      // 医疗保险类型
	MedicalOrganizationType string         `json:"medical_organization_type"`  // 医疗机构类型
	NoteTitle               string         `json:"note_title"`                 // 发票标题
	PatientGender           int            `json:"patient_gender"`             // 患者性别
	PatientName             string         `json:"patient_name"`               // 患者姓名
	Payee                   string         `json:"payee"`                      // 收款人
	PaymentChannel          string         `json:"payment_channel"`            // 付款渠道
	PaymentsInfo            []PaymentInfo  `json:"payments_info"`              // 付款信息列表
	Reviewer                string         `json:"reviewer"`                   // 复核人
	ServiceSerialNo         string         `json:"service_serial_no"`          // 服务流水号
	TotalCost               float64        `json:"total_cost"`                 // 总费用
	TreatmentDate           string         `json:"treatment_date"`             // 就诊日期
	UnifiedSocialCreditCode string         `json:"unified_social_credit_code"` // 统一社会信用代码
	WorkUnit                string         `json:"work_unit"`                  // 工作单位
	InvoiceUrl              string         `json:"invoice_url"`                // 发票地址
	OriginResult            string         `json:"origin_result"`
}

type Checksum struct {
	CostCategories int `json:"cost_categories"`  // 费用分类数量
	CostDetailList int `json:"cost_detail_list"` // 费用明细数量
}

type CostCategory struct {
	Cost float64 `json:"cost"` // 费用金额
	Name string  `json:"name"` // 费用分类名称
}

type CostDetail struct {
	Amount        float64 `json:"amount"`      // 数量
	Class         string  `json:"class"`       // 分类
	Cost          float64 `json:"cost"`        // 费用金额
	ItemCoding    string  `json:"item_coding"` // 项目编码
	MedicalLevel  string  `json:"medical_level"`
	MedicalType   string  `json:"medical_type"` // 医疗类型
	Name          string  `json:"name"`         // 名称
	OriginOCRName string  `json:"origin_ocr_name"`
	SelfPay       string  `json:"self_pay"`       // 自费金额
	SelfPayRatio  string  `json:"self_pay_ratio"` // 自费比例
	Spec          string  `json:"spec"`           // 规格
	Unit          string  `json:"unit"`           // 单位
	UnitPrice     string  `json:"unit_price"`     // 单价
}

type PaymentInfo struct {
	Cost float64 `json:"cost"` // 付款金额
	Name string  `json:"name"` // 付款方式名称
}
