// The fastest unofficial Schwab TraderAPI wrapper
// Copyright (C) 2024 Samuel Troyer <github.com/samjtro>
// See the GNU General Public License for more details
package schwab

import (
	"fmt"
	"net/http"

	"github.com/bytedance/sonic"
)

var (
	accountEndpoint string = "https://api.schwabapi.com/trader/v1"

	// Accounts
	endpointAccountNumbers string = accountEndpoint + "/accounts/accountNumbers"
	endpointAccounts       string = accountEndpoint + "/accounts"
	endpointAccount        string = accountEndpoint + "/accounts/%s"
	endpointUserPreference string = accountEndpoint + "/userPreference"

	// Orders
	endpointOrders        string = accountEndpoint + "/orders"
	endpointAccountOrders string = accountEndpoint + "/accounts/%s/orders"
	endpointAccountOrder  string = accountEndpoint + "/accounts/%s/orders/%s"
	endpointPreviewOrder  string = accountEndpoint + "/accounts/%s/previewOrder"

	// Transactions
	endpointTransactions string = accountEndpoint + "/accounts/%s/transactions"
	endpointTransaction  string = accountEndpoint + "/accounts/%s/transactions/%s"
)

type Transaction struct {
	ActivityID     int `json:"ActivityId"`
	Time           string
	User           User
	Description    string
	AccountNumber  string
	Type           string
	Status         string
	SubAccount     string
	TradeDate      string
	SettlementDate string
	PositionId     int
	OrderId        int
	NetAmount      int
	ActivityType   string
	TransferItems  TransferItems
}

type User struct {
	CdDomainId     string
	Login          string
	Type           string
	UserId         int
	SystemUserName string
	FirstName      string
	LastName       string
	BrokerRepCode  string
}

type TransferItems struct {
	Instrument     TransactionInstrument
	Amount         int
	Cost           int
	Price          int
	FeeType        string
	PositionEffect string
}

type TransactionInstrument struct {
	Cusip        string
	Symbol       string
	Description  string
	InstrumentId int
	NetChange    int
	Type         string
}

type AccountNumbers struct {
	AccountNumber string
	HashValue     string
}

type Account struct {
	Type                    string
	AccountNumber           string
	RoundTrips              int
	IsDayTrader             bool
	IsClosingOnlyRestricted bool
	PFCBFlag                bool
	Positions               []Position
	InitialBalances         InitialBalance
	CurrentBalances         CurrentBalance
	ProjectedBalances       ProjectedBalance
	AggregatedBalance       AggregatedBalance
}

type Position struct {
	ShortQuantity                int
	AveragePrice                 float64
	CurrentDayProfitLoss         float64
	LongQuantity                 int
	SettledLongQuantity          int
	SettledShortQuantity         int
	AgedQuantity                 int
	Instrument                   AccountInstrument
	MarketValue                  float64
	MaintenanceRequirement       float64
	AverageLongPrice             float64
	AverageShortPrice            float64
	TaxLotAverageLongPrice       float64
	TaxLotAverageShortPrice      float64
	LongOpenProfitLoss           float64
	ShortOpenProfitLoss          float64
	PreviousSessionLongQuantity  int
	PreviousSessionShortQuantity int
	CurrentDayCost               float64
}

type AccountInstrument struct {
	Cusip        string
	Symbol       string
	Description  string
	InstrumentID int
	NetChange    float64
	Type         string
}

type InitialBalance struct {
	AccruedInterest                  float64
	AvailableFundsNonMarginableTrade float64
	BondValue                        float64
	BuyingPower                      float64
	CashBalance                      float64
	CashAvailableForTrading          float64
	CashReceipts                     float64
	DayTradingBuyingPower            float64
	DayTradingBuyingPowerCall        float64
	DayTradingEquityCall             float64
	Equity                           float64
	EquityPercentage                 float64
	LiquidationValue                 float64
	LongMarginValue                  float64
	LongOptionMarketValue            float64
	LongStockValue                   float64
	MaintenanceCall                  float64
	MaintenanceRequirement           float64
	Margin                           float64
	MarginEquity                     float64
	MoneyMarketFund                  float64
	MutualFundValue                  float64
	RegTCall                         float64
	ShortMarginValue                 float64
	ShortOptionMarketValue           float64
	ShortStockValue                  float64
	TotalCash                        float64
	IsInCall                         float64
	UnsettledCash                    float64
	PendingDeposits                  float64
	MarginBalance                    float64
	ShortBalance                     float64
	AccountValue                     float64
}

/*type CurrentBalance struct {
	AvailableFunds                   float64
	AvailableFundsNonMarginableTrade float64
	BuyingPower                      float64
	BuyingPowerNonMarginableTrade    float64
	DayTradingBuyingPower            float64
	DayTradingBuyingPowerCall        float64
	Equity                           float64
	EquityPercentage                 float64
	LongMarginValue                  float64
	MaintenanceCall                  float64
	MaintenanceRequirement           float64
	MarginBalance                    float64
	RegTCall                         float64
	ShortBalance                     float64
	ShortMarginValue                 float64
	SMA                              float64
	IsInCall                         float64
	StockBuyingPower                 float64
	OptionBuyingPower                float64
}*/

type CurrentBalance struct {
	AccruedInterest       float64
	CashBalance           float64
	CashReceipts          float64
	LongOptionMarketValue float64
	LiquidationValue      float64
	SMA                   float64
}

type ProjectedBalance struct {
	AvailableFunds                   float64
	AvailableFundsNonMarginableTrade float64
	BuyingPower                      float64
	BuyingPowerNonMarginableTrade    float64
	DayTradingBuyingPower            float64
	DayTradingBuyingPowerCall        float64
	Equity                           float64
	EquityPercentage                 float64
	LongMarginValue                  float64
	MaintenanceCall                  float64
	MaintenanceRequirement           float64
	MarginBalance                    float64
	RegTCall                         float64
	ShortBalance                     float64
	ShortMarginValue                 float64
	SMA                              float64
	IsInCall                         float64
	StockBuyingPower                 float64
	OptionBuyingPower                float64
}

type AggregatedBalance struct {
	CurrentLiquidationValue float64
	LiquidationValue        float64
}

// Get encrypted account numbers for trading
func (agent *Agent) GetAccountNumbers() ([]AccountNumbers, error) {
	req, err := http.NewRequest("GET", endpointAccountNumbers, nil)
	check(err)
	body, err := agent.Handler(req)
	check(err)
	var accountNumbers []AccountNumbers
	err = sonic.Unmarshal([]byte(body), &accountNumbers)
	check(err)
	return accountNumbers, nil
}

// Get all accounts associated with the user logged in
func (agent *Agent) GetAccounts() ([]Account, error) {
	req, err := http.NewRequest("GET", endpointAccounts, nil)
	check(err)
	body, err := agent.Handler(req)
	check(err)
	var accounts []Account
	err = sonic.Unmarshal([]byte(body), &accounts)
	check(err)
	return accounts, nil
}

// Get account by encrypted account id
func (agent *Agent) GetAccount(id string) (Account, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(endpointAccount, id), nil)
	check(err)
	body, err := agent.Handler(req)
	check(err)
	var account Account
	err = sonic.Unmarshal([]byte(body), &account)
	check(err)
	return account, nil
}

// Get all transactions for the user logged in
//func (agent *Agent) GetTransactions() ([]Transaction, error) {}

func (agent *Agent) GetTransaction(accountNumber, transactionId string) (Transaction, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(endpointTransaction, accountNumber, transactionId), nil)
	check(err)
	body, err := agent.Handler(req)
	check(err)
	var transaction Transaction
	err = sonic.Unmarshal([]byte(body), &transaction)
	check(err)
	return transaction, nil
}

func SubmitLimitOrder() {}

func SubmitMarketOrder() {}
