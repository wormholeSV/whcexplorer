package common

type errCode int

const (
	Success = 0

	ErrUnSupportTxType errCode = 2000 + iota
	ErrGetTxType
	ErrUnSupportTxVersion
	ErrUnSupportEcosystem
	ErrUnSupportCrowdSaleSP
	ErrUnSupportPrecision
	ErrIncorrectAddress
	ErrEmptyAddressList

	// RPC error
	ErrCreatePayload
	ErrCreateRawTxInput
	ErrCreateRawTxReference
	ErrCreateRawTxOpReturn
	ErrCreateRawTxChange
	ErrSendRawTransaction
	ErrWhcDecodeTransaction

	// Number error
	ErrConvertFloat64
	ErrConvertInt
	ErrIncorrectAmount

	ErrFormItems
	ErrEmptyQueryParam
	ErrCategoryNotFound
	ErrHexStringFormat
	ErrHash256Format
	ErrTxDeserialize
	ErrCanNotGetUtxo
	ErrCanNotGetInputs
	ErrDecodeRawTransaction

	// handle
	ErrInsertBCH
	ErrInsertWhc

	// database
	ErrGetHistoryList
	ErrEmptyHistoryList
	ErrEmptyHistoryDetail
	ErrEmptyHistorySpDetail
	ErrGetPropertyByAddress
	ErrListProperties
	ErrGetAllCrowdSale
	ErrQueryTotal
	ErrQueryTransactions

	// account
	ErrChallenge
	ErrCreateWallet
	ErrUpdateWallet
	ErrLogin
	ErrNotUUID
	ErrVerify

	// server error
	ErrGetProperties
	ErrGetBalanceFromRedis
	ErrGetBalanceFromDatabase
	ErrEmptyBalance
)

type Response struct {
	Code    errCode     `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

var errNotFound = Response{
	Code:    4004,
	Message: "unknown error",
	Result:  nil,
}

var errMapping = map[errCode]*Response{

}
func ApiErrorWithMsg(code errCode, msg string) *Response {
	r, ok := errMapping[code]
	if !ok {
		return &errNotFound
	}

	if msg != "" {
		// copy the origin response
		customMsg := *r
		customMsg.Message = msg
		return &customMsg
	}

	return r
}

func ApiError(code errCode) *Response {
	r, ok := errMapping[code]
	if !ok {
		return &errNotFound
	}

	return r
}

func ApiSuccess(data interface{}) *Response {
	return &Response{
		Code:    0,
		Message: "",
		Result:  data,
	}
}