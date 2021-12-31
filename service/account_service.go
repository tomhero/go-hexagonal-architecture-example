package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"strings"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAcccount(customerID int, req NewAccountRequest) (*AccountResponse, error) {
	// NOTE : Validation here
	if req.Amount < 3000 {
		return nil, errs.NewValidationError("amount at least 3,000")
	}

	if strings.ToLower(req.Accountype) != "saving" && strings.ToLower(req.Accountype) != "checking" {
		return nil, errs.NewValidationError("account type should be `saving` or `checking`")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"), // NOTE : `2006-01-02 15:04:05` เป็นแค่การบอก format เท่านั้น!!!
		Accountype:  req.Accountype,
		Amount:      req.Amount,
	}

	newAcc, err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	responseData := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		Accountype:  newAcc.Accountype,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}

	return &responseData, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accRepo.GetAll(customerID)
	if err != nil {
		// NOTE : ตรงนี้คือ Technical error
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	responseData := []AccountResponse{}
	for _, account := range accounts {
		responseData = append(responseData, AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			Accountype:  account.Accountype,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}
	return responseData, nil
}
