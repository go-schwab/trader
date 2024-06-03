package account

var (
	// Orders
	Endpoint_orders        string = Endpoint + "/orders"
	Endpoint_accountOrders string = Endpoint + "/accounts/%s/orders"
	Endpoint_accountOrder  string = Endpoint + "/accounts/%s/orders/%s"
	Endpoint_previewOrder  string = Endpoint + "/accounts/%s/previewOrder"

	// Transactions
	Endpoint_transactions string = Endpoint + "/accounts/%s/transactions"
	Endpoint_transaction  string = Endpoint + "/accounts/%s/transactions/%s"
)

func SubmitLimitOrder()  {}
func SubmitMarketOrder() {}
