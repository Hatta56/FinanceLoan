package module

import (
	"Finance/api/service"
)

type ServiceModule interface {
	GetCustomerService() service.CustomerService
	GetLimitLoanService() service.LimitLoanService
	GetTransactionService() service.TransactionService
}
