package account

var (
	Endpoint string = "https://api.schwabapi.com/trader/v1"

	// Accounts
	Endpoint_accountNumbers string = Endpoint + "/accounts/accountNumbers"
	Endpoint_accounts       string = Endpoint + "/accounts"
	Endpoint_account        string = Endpoint + "/accounts/%s"
	Endpoint_userPreference string = Endpoint + "/userPreference"
)
