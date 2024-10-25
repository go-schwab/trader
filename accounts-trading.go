package trader

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
)

/* TODO:
[ ] http.NewRequest -> agent.client
*/

var (
	accountEndpoint string = "https://api.schwabapi.com/trader/v1"

	// Accounts
	endpointAccountNumbers string = accountEndpoint + "/accounts/accountNumbers"
	endpointAccounts       string = accountEndpoint + "/accounts"
	endpointAccount        string = accountEndpoint + "/accounts/%s"
	// endpointUserPreference string = accountEndpoint + "/userPreference"

	// Orders
	endpointOrders        string = accountEndpoint + "/orders"
	endpointAccountOrders string = accountEndpoint + "/accounts/%s/orders"
	endpointAccountOrder  string = accountEndpoint + "/accounts/%s/orders/%s"
	// endpointPreviewOrder  string = accountEndpoint + "/accounts/%s/previewOrder"

	// Transactions
	// endpointTransactions string = accountEndpoint + "/accounts/%s/transactions"
	endpointTransaction string = accountEndpoint + "/accounts/%s/transactions/%s"
)

type Transaction struct {
	ActivityId     int
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
	Instrument     InstrumentRef
	Amount         int
	Cost           int
	Price          int
	FeeType        string
	PositionEffect string
}

type InstrumentRef struct {
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

type FullOrder struct {
	Session                  string
	Duration                 string
	OrderType                string
	CancelTime               string
	ComplexOrderStrategyType string
	Quantity                 int
	FilledQuantity           int
	RemainingQuantity        int
	RequestedDestination     string
	DestinationLinkName      string
	ReleaseTime              string
	StopPrice                int
	StopPriceLinkBasis       string
	StopPriceLinkType        string
	StopPriceOffset          int
	StopType                 string
	Price                    string
	TaxLotMethod             string
	OrderLegCollection       []FullOrderLeg
	ActivationPrice          int
	SpecialInstruction       string
	OrderStrategyType        string
	OrderId                  int
	Cancelable               bool
	Editable                 bool
	Status                   string
	EnteredTime              string
	CloseTime                string
	Tag                      string
	AccountNumber            int
	OrderActivityCollection  []FullOrderActivity
	ReplacingOrderCollection string
	ChildOrderStrategies     string
	StatusDescription        string
}

type FullOrderActivity struct {
	ActivityType           string
	ExecutionType          string
	Quantity               int
	OrderRemainingQuantity int
	ExecutionLegs          []FullExecutionLeg
}

type FullExecutionLeg struct {
	LegId             int
	Price             int
	Quantity          int
	MismarkedQuantity int
	InstrumentId      int
	Time              string
}

type FullOrderLeg struct {
	OrderLegType   string
	LegId          int
	Instrument     InstrumentRef
	Instruction    string
	PositionEffect string
	Quantity       int
	QuantityType   string
	DivCapGains    string
	ToSymbol       string
}

type SingleLegOrder struct {
	OrderType   string `default:"MARKET"`
	Session     string `default:"NORMAL"`
	Duration    string `default:"DAY"`
	Strategy    string `default:"SINGLE"`
	Instruction string
	Quantity    float32
	Instrument  SimpleOrderInstrument
}

type MultiLegOrder struct {
	OrderType          string // LIMIT, MARKET
	Session            string // NORMAL
	Duration           string // DAY
	Strategy           string // SINGLE
	OrderLegCollection []SimpleOrderLeg
}

type SimpleOrderLeg struct {
	Instruction string
	Quantity    float32
	Instrument  SimpleOrderInstrument
}

type SimpleOrderInstrument struct {
	Symbol    string
	AssetType string // EQUITY
}

type (
	SingleLegOrderComposition      func(order *SingleLegOrder)
	MultiLegSimpleOrderComposition func(order *MultiLegOrder)
)

// Create a new Market order
func CreateSingleLegOrder(opts ...SingleLegOrderComposition) *SingleLegOrder {
	order := &SingleLegOrder{OrderType: "MARKET"}
	for _, opt := range opts {
		opt(order)
	}
	return order
}

// Set SingleLegOrder.OrderType
func OrderType(t string) SingleLegOrderComposition {
	return func(order *SingleLegOrder) {
		order.OrderType = t
	}
}

// Set SingleLegOrder.Session
func Session(session string) SingleLegOrderComposition {
	return func(order *SingleLegOrder) {
		order.Session = session
	}
}

// Set SingleLegOrder.Duration
func Duration(duration string) SingleLegOrderComposition {
	return func(order *SingleLegOrder) {
		order.Duration = duration
	}
}

// Set SingleLegOrder.Strategy
func Strategy(strategy string) SingleLegOrderComposition {
	return func(order *SingleLegOrder) {
		order.Strategy = strategy
	}
}

// Set SingleLegOrder.Instruction
func Instruction(instruction string) SingleLegOrderComposition {
	return func(order *SingleLegOrder) {
		order.Instruction = instruction
	}
}

// Set SingleLegOrder.Quantity
func Quantity(quantity float32) SingleLegOrderComposition {
	return func(order *SingleLegOrder) {
		order.Quantity = quantity
	}
}

// Set SingleLegOrder.Instrument
func Instrument(instrument SimpleOrderInstrument) SingleLegOrderComposition {
	return func(order *SingleLegOrder) {
		order.Instrument = instrument
	}
}

var OrderTemplate = `
{
  "orderType": "%s",
  "session": "%s",
  "duration": "%s",
  "orderStrategyType": "%s",
  "orderLegCollection": [
    %s
  ]
}
`

var LegTemplate = `
{
  "instruction": "%s",
  "quantity": %f,
  "instrument": {
    "symbol": "%s",
    "assetType": "%s"
  }
},
`

var LegTemplateLast = `
{
  "instruction": "%s",
  "quantity": %f,
  "instrument": {
    "symbol": "%s",
    "assetType": "%s"
  }
},
`

func marshalSingleLegOrder(order *SingleLegOrder) string {
	return fmt.Sprintf(OrderTemplate, order.OrderType, order.Session, order.Duration, order.Strategy, fmt.Sprintf(LegTemplate, order.Instruction, order.Quantity, order.Instrument.Symbol, order.Instrument.AssetType))
}

func marshalMultiLegOrder(order *MultiLegOrder) string {
	var legs string
	// UNTESTED
	for i, leg := range order.OrderLegCollection {
		if i != order.OrderLegCollection.length-1 {
			legs += fmt.Sprintf(LegTemplate, order.Instruction, order.Quantity, order.Instrument.Symbol, order.Instrument.AssetType)
		} else {
			legs += fmt.Sprintf(LegTemplateLast, order.Instruction, order.Quantity, order.Instrument.Symbol, order.Instrument.AssetType)
		}
	}
	return fmt.Sprintf(OrderTemplate)
}

// Submit a single-leg order for the specified encrypted account ID
func (agent *Agent) SubmitSingleLegOrder(hashValue string, order *SingleLegOrder) error {
	orderJson := marshalSingleLegOrder(order)
	req, err := http.NewRequest("POST", fmt.Sprintf(endpointAccountOrders, hashValue), strings.NewReader(orderJson))
	isErrNil(err)
	req.Header.Set("Content-Type", "application/json")
	_, err = agent.Handler(req)
	isErrNil(err)
	return nil
}

// Get a specific order by account number & order ID
func (agent *Agent) GetOrder(accountNumber, orderID string) (FullOrder, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(endpointAccountOrder, accountNumber, orderID), nil)
	isErrNil(err)
	resp, err := agent.Handler(req)
	isErrNil(err)
	var order FullOrder
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	isErrNil(err)
	err = sonic.Unmarshal(body, &order)
	isErrNil(err)
	return order, nil
}

// fromEnteredTime, toEnteredTime format:
// yyyy-MM-ddTHH:mm:ss.SSSZ
func (agent *Agent) GetAccountOrders(accountNumber, fromEnteredTime, toEnteredTime string) ([]FullOrder, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(endpointAccountOrders, accountNumber), nil)
	isErrNil(err)
	q := req.URL.Query()
	q.Add("fromEnteredTime", fromEnteredTime)
	q.Add("toEnteredTime", toEnteredTime)
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	isErrNil(err)
	var orders []FullOrder
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	isErrNil(err)
	err = sonic.Unmarshal(body, &orders)
	isErrNil(err)
	return orders, nil
}

// WIP:
// fromEnteredTime, toEnteredTime format:
// yyyy-MM-ddTHH:mm:ss.SSSZ
func (agent *Agent) GetAllOrders(fromEnteredTime, toEnteredTime string) ([]FullOrder, error) {
	req, err := http.NewRequest("GET", endpointOrders, nil)
	isErrNil(err)
	q := req.URL.Query()
	q.Add("fromEnteredTime", fromEnteredTime)
	q.Add("toEnteredTime", toEnteredTime)
	req.URL.RawQuery = q.Encode()
	resp, err := agent.Handler(req)
	isErrNil(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	isErrNil(err)
	var orders []FullOrder
	/*err = sonic.Unmarshal(body, &orders)
	isErrNil(err)*/
	fmt.Println(body)
	return orders, nil
}

// Get encrypted account numbers for trading
func (agent *Agent) GetAccountNumbers() ([]AccountNumbers, error) {
	req, err := http.NewRequest("GET", endpointAccountNumbers, nil)
	isErrNil(err)
	resp, err := agent.Handler(req)
	isErrNil(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	isErrNil(err)
	var accountNumbers []AccountNumbers
	err = sonic.Unmarshal(body, &accountNumbers)
	isErrNil(err)
	return accountNumbers, nil
}

// Get all accounts associated with the user logged in
func (agent *Agent) GetAccounts() ([]Account, error) {
	req, err := http.NewRequest("GET", endpointAccounts, nil)
	isErrNil(err)
	resp, err := agent.Handler(req)
	isErrNil(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	isErrNil(err)
	var accounts []Account
	err = sonic.Unmarshal(body, &accounts)
	isErrNil(err)
	return accounts, nil
}

// Get account by encrypted account id
func (agent *Agent) GetAccount(id string) (Account, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(endpointAccount, id), nil)
	isErrNil(err)
	resp, err := agent.Handler(req)
	isErrNil(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	isErrNil(err)
	var account Account
	err = sonic.Unmarshal(body, &account)
	isErrNil(err)
	return account, nil
}

// Get all transactions for the user logged in
// func (agent *Agent) GetTransactions() ([]Transaction, error) {}

// Get a transaction for a specific account id
func (agent *Agent) GetTransaction(accountNumber, transactionId string) (Transaction, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(endpointTransaction, accountNumber, transactionId), nil)
	isErrNil(err)
	resp, err := agent.Handler(req)
	isErrNil(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	isErrNil(err)
	var transaction Transaction
	err = sonic.Unmarshal(body, &transaction)
	isErrNil(err)
	return transaction, nil
}
