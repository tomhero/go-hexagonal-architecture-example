package repository

type accountRepositoryMock struct {
	accounts []Account
	account  Account
}

func NewAccountRepositoryMock() accountRepositoryMock {
	// NOTE : Mock account data here
	accounts := []Account{
		{AccountID: 1, CustomerID: 1000, Accountype: "saving", OpeningDate: "2022-01-01 00:00:00", Amount: 5000, Status: 1},
		{AccountID: 2, CustomerID: 1001, Accountype: "checking", OpeningDate: "2022-01-01 00:00:01", Amount: 4500, Status: 0},
		{AccountID: 3, CustomerID: 1000, Accountype: "checking", OpeningDate: "2022-01-02 00:00:03", Amount: 9999, Status: 0},
	}
	return accountRepositoryMock{accounts: accounts, account: accounts[0]}
}

func (r accountRepositoryMock) GetAll(customerID int) ([]Account, error) {
	customerAccounts := []Account{}
	for _, acc := range r.accounts {
		if acc.CustomerID == customerID {
			customerAccounts = append(customerAccounts, acc)
		}
	}
	return customerAccounts, nil
}

func (r accountRepositoryMock) Create(acc Account) (*Account, error) {
	acc.AccountID = 4
	return &acc, nil
}
