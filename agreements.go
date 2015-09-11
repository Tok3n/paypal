package paypal
import (
	"fmt"
)

type (
	BillinAgreementReq struct {
		Name string `json:"name"`
		Description string `json:"description"`
		StartDate string`json:"start_date"`
		Payer *Payer `json:"payer"`
		Plan *Plan `json:"plan"`
	}
)

func (c *Client) CreateAgreement(bar *BillinAgreementReq)(*Agreement, error){
	req, err := NewRequest("POST", fmt.Sprintf("%s/payments/billing-agreements", c.APIBase), bar)
	if err != nil {
		return nil, err
	}

	v := &Agreement{}

	err = c.SendWithAuth(req, v)
		if err != nil {
		return nil, err
	}

	return v, nil
}