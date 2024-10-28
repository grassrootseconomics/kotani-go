package kotani

type (
	Currency    string
	CountryCode string
	Network     string
	IDType      string
	Chain       string
	Token       string
	Status      string
)

const (
	KES Currency = "KES"
	UGX Currency = "UGX"

	KE CountryCode = "KE"
	UG CountryCode = "UG"

	MPESA Network = "MPESA"
	MTN   Network = "MTN"

	NATIONAL_ID IDType = "NATIONAL_ID"

	CELO Chain = "CELO"

	CUSD Token = "CUSD"
	CKES Token = "CKES"
	USDC Token = "USDC"
	USDT Token = "USDT"

	PENDING        Status = "PENDING"
	SUCCESSFUL     Status = "SUCCESSFUL"
	FAILED         Status = "FAILED"
	CANCELLED      Status = "CANCELLED"
	REVERSED       Status = "REVERSED"
	IN_PROGRESS    Status = "IN_PROGRESS"
	INITIATED      Status = "INITIATED"
	ERROR_OCCURRED Status = "ERROR_OCCURRED"
	DECLINED       Status = "DECLINED"
	EXPIRED        Status = "EXPIRED"
	REQUIRE_REVIEW Status = "REQUIRE_REVIEW"
)
