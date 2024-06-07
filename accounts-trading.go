// The fastest unofficial Schwab TraderAPI wrapper
// Copyright (C) 2024 Samuel Troyer <github.com/samjtro>
// See the GNU General Public License for more details
package schwab

var (
	Endpoint string = "https://api.schwabapi.com/trader/v1"

	// Accounts
	Endpoint_accountNumbers string = Endpoint + "/accounts/accountNumbers"
	Endpoint_accounts       string = Endpoint + "/accounts"
	Endpoint_account        string = Endpoint + "/accounts/%s"
	Endpoint_userPreference string = Endpoint + "/userPreference"

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
