package define

type AcountData struct {
	AcountID  string `json:"acount_id"`
	Name      string `json:"name"`
	EmailAddr string `json:"email_address"`
	Age       int    `json:"age"`
}
