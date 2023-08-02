package response

// Icd 疾病查询
type Icd struct {
	Id           int64  `json:"ID"`
	PyShort      string `json:"py_short"`
	PyFull       string `json:"py_full"`
	ChapterCode  string `json:"chapter_code"`
	ChapterScope string `json:"chapter_scope"`
	ChapterName  string `json:"chapter_name"`
	SectionScope string `json:"section_scope"`
	SectionName  string `json:"section_name"`
	CategoryCode string `json:"category_code"`
	CategoryName string `json:"category_name"`
	SubCode      string `json:"sub_code"`
	SubName      string `json:"sub_name"`
	IcdCode      string `json:"icd_code"`
	IcdName      string `json:"icd_name"`
}
