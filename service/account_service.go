package service

import (
	"hexagonal/errs"
	"hexagonal/logs"
	"hexagonal/repository"
	"strings"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (a accountService) NewAccount(id int, req NewAccountRequest) (*AccountResponse, error) {

	if req.Amount < 5000 {
		return nil, errs.NewValidatonError("amount at least 5,000")

	}
	if strings.ToLower(req.AccountType) != "saving" && strings.ToLower(req.AccountType) != "checking" {
		return nil, errs.NewValidatonError("account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  id,
		OpeningDate: time.Now().Format("2006-01-2 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      1,
	}

	newAcc, err := a.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}

	return &response, nil
}

func (a accountService) GetAccounts(id int) ([]AccountResponse, error) {
	accounts, err := a.accRepo.GetAll(id)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	res := []AccountResponse{}
	for _, acc := range accounts {

		res = append(res, AccountResponse{
			AccountID:   acc.AccountID,
			OpeningDate: acc.OpeningDate,
			AccountType: acc.AccountType,
			Amount:      acc.Amount,
			Status:      acc.Status,
		})
	}
	return res, nil
}
