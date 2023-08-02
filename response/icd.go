package response

type Icd struct {
	Id           int64  `json:"ID"`
	PyShort      string `json:"py_short"`
	PyFull       string `json:"py_full"`
	ChapterCode  string `json:"chapter_code"`
	ChapterScope string `json:"chapter_scope"`
	ChapterName  int    `json:"chapter_name"`
	SectionScope int    `json:"section_scope"`
	SectionName  int    `json:"section_name"`
	CategoryCode int    `json:"category_code"`
	CategoryName int    `json:"category_name"`
	SubCode      int    `json:"sub_code"`
	SubName      int    `json:"sub_name"`
	IcdCode      int    `json:"icd_code"`
	IcdName      int    `json:"icd_name"`
}
