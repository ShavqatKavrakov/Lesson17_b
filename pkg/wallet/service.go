package wallet

import (
	"errors"
	"io/ioutil"
	"strconv"

	"github.com/ShavqatKavrakov/Lesson17_b/pkg/types"
	"github.com/google/uuid"
)

var ErrPhoneRegistered = errors.New("phone already registered")
var ErrAccountNotFound = errors.New("account not found")
var ErrAmountMostBePositive = errors.New("amount must be greater than zero")
var ErrNotEnouthBalance = errors.New("not enough balance in account")
var ErrPaymentNotFound = errors.New("payment not found")
var ErrFavoriteNotFound = errors.New("favorite not found")

type Service struct {
	nextAccountId int64
	accounts      []*types.Account
	payments      []*types.Payment
	favorites     []*types.Favorite
}

func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, acc := range s.accounts {
		if acc.Phone == phone {
			return nil, ErrPhoneRegistered
		}
	}
	s.nextAccountId++
	account := &types.Account{
		ID:      s.nextAccountId,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)
	return account, nil
}

func (s *Service) FindAccountById(accountId int64) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.ID == accountId {
			return account, nil
		}
	}
	return nil, ErrAccountNotFound
}

func (s *Service) Deposit(accountId int64, amount types.Money) (*types.Account, error) {
	if amount <= 0 {
		return nil, ErrAmountMostBePositive
	}
	account, err := s.FindAccountById(accountId)
	if err != nil {
		return nil, err
	}
	account.Balance += amount
	return account, nil
}

func (s *Service) Pay(acountId int64, category types.PaymentCategory, amount types.Money) (*types.Payment, error) {
	if amount < 0 {
		return nil, ErrAmountMostBePositive
	}
	account, err := s.FindAccountById(acountId)
	if err != nil {
		return nil, err
	}
	if account.Balance < amount {
		return nil, ErrNotEnouthBalance
	}
	account.Balance -= amount

	paymentID := uuid.New().String()
	payment := &types.Payment{
		ID:        paymentID,
		AccountID: acountId,
		Amount:    amount,
		Category:  category,
		Status:    types.PaymentStatusInProgress,
	}
	s.payments = append(s.payments, payment)
	return payment, nil
}

func (s *Service) FindPaymentById(paymentId string) (*types.Payment, error) {
	for _, payment := range s.payments {
		if payment.ID == paymentId {
			return payment, nil
		}
	}
	return nil, ErrPaymentNotFound
}
func (s *Service) ExportAccountHistory(accountID int64) ([]types.Payment, error) {
	_, err := s.FindAccountById(accountID)
	if err != nil {
		return nil, ErrAccountNotFound
	}
	var payments []types.Payment
	for _, payment := range s.payments {
		if payment.AccountID == accountID {
			payments = append(payments, *payment)
		}
	}
	return payments, nil
}
func (s *Service) HistoryToFile(payments []types.Payment, dir string, records int) error {
	var result string
	t := 0
	for index, payment := range payments {

		result += strconv.Itoa(int(payment.Amount)) + " " + string(payment.Category) + " " + string(payment.Status) + ";\n"
		if (index+1)%records == 0 {
			CreatePaymentsFile(result, dir, t)
			t++
			result = ""
		}
		if (len(payments)-t*records) < records && len(payments) != t*records && len(payments) == index+1 {
			CreatePaymentsFile(result, dir, t)
		}
	}
	return nil
}
func CreatePaymentsFile(result string, dir string, count int) error {
	if count == 0 {
		dir += "/payments" + ".dump"
	} else {
		dir += "/payments" + strconv.Itoa(count) + ".dump"
	}
	err := ioutil.WriteFile(dir, []byte(result), 0666)
	if err != nil {
		return err
	}
	return nil
}
