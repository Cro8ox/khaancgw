package structs

type (
	StatementResponse struct {
		Account      string        `json:"account"`
		IBAN         string        `json:"iban"`
		Currency     string        `json:"currency"`
		CustomerName string        `json:"customerName"`
		ProductName  string        `json:"productName"`
		Branch       string        `json:"branch"`
		BranchName   string        `json:"branchName"`
		BeginBalance float64       `json:"beginBalance"`
		EndBalance   float64       `json:"endBalance"`
		BeginDate    string        `json:"beginDate"`
		EndDate      string        `json:"endDate"`
		Total        Total         `json:"total"`
		Transactions []Transaction `json:"transactions"`
	}
	Total struct {
		Count  uint64  `json:"count"`
		Credit float64 `json:"credit"`
		Debit  float64 `json:"debit"`
	}
	Transaction struct {
		Record         uint64  `json:"record"`
		TranDate       string  `json:"tranDate"`
		PostDate       string  `json:"postDate"`
		Time           string  `json:"time"`
		Branch         string  `json:"branch"`
		Teller         string  `json:"teller"`
		Journal        uint64  `json:"journal"`
		Code           uint64  `json:"code"`
		Amount         float64 `json:"amount"`
		Balance        float64 `json:"balance"`
		Debit          float64 `json:"debit"`
		Correction     uint64  `json:"correction"`
		Description    string  `json:"description"`
		RelatedAccount string  `json:"relatedAccount"`
	}
)
