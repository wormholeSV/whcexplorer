package api

type errCode int

var errNotFound = &Response{
	Code:    4004,
	Message: "unknown error",
	Result:  nil,
}

const (
	// consensus error

	// RPC error
	ErrGetSto errCode = 2100 + iota

	// User input error
	ErrInvalidParam errCode = 2200 + iota
	ErrInvalidBlockHashOrHeight
	ErrInvalidBlockHash
	ErrInvalidBlockHeight
	ErrInvalidBlockTimestamp
	ErrInvalidInteger
	ErrIncorrectPropertyQueryByKeyword

	// database
	ErrFetchBlockList errCode = 2300 + iota
	ErrFetchBlockListCount
	ErrFetchBlockByHash
	ErrNotFoundBlockHash
	ErrFetchBlockByHeight
	ErrNotFoundBlockHeight
	ErrNotFoundTransactionHash
	ErrFetchTransactionByHash
	ErrGetTransactionListCount
	ErrFetchTransactionListInBlockHeight
	ErrFetchTransactionListInBlockHash
	ErrGetNetworkInfo
	ErrGetTotalWormholeTxCount
	ErrGetWormholeTxCountInPeriod
	ErrGetPropertyNameByID
	ErrFetchBurnListCount
	ErrFetchLatestWhccountInfo
	ErrGetAddressDetail
	ErrGetAddressTransactions
	ErrGetAddressProperties
	ErrGetAddressPropertyTransactions
	ErrGetProperties
	ErrGetPropertyByID
	ErrGetPropertyChangeHistory
	ErrGetPropertyTransactions
	ErrListPropertiesCount
	ErrListProperties
	ErrGetPropertyByName

	// server or handle error
	ErrSetUTCLocation
)

var errMapping = map[errCode]*Response{
	// consensus error

	// RPC error
	ErrGetSto: {Code: ErrGetSto, Message: "Get sto transaction via rpc failed"},

	// User input error
	ErrInvalidParam:             {Code: ErrInvalidParam, Message: "input param error"},
	ErrInvalidBlockHashOrHeight: {Code: ErrInvalidBlockHashOrHeight, Message: "invalid block hash or block height"},
	ErrInvalidBlockHash:         {Code: ErrInvalidBlockHash, Message: "invalid block hash"},
	ErrInvalidBlockHeight:       {Code: ErrInvalidBlockHeight, Message: "invalid block height"},
	ErrInvalidBlockTimestamp:    {Code: ErrInvalidBlockTimestamp, Message: "invalid block timestamp"},
	ErrInvalidInteger:           {Code: ErrInvalidInteger, Message: "Invalid integer"},

	// database
	ErrFetchBlockList:                    {Code: ErrFetchBlockList, Message: "Fetch block list failed"},
	ErrFetchBlockListCount:               {Code: ErrFetchBlockListCount, Message: "Fetch block list count failed"},
	ErrFetchBlockByHash:                  {Code: ErrFetchBlockByHash, Message: "Fetch block information via hash failed"},
	ErrNotFoundBlockHash:                 {Code: ErrNotFoundBlockHash, Message: "The specified block hash not found"},
	ErrFetchBlockByHeight:                {Code: ErrFetchBlockByHeight, Message: "Fetch block information via height failed"},
	ErrNotFoundBlockHeight:               {Code: ErrNotFoundBlockHeight, Message: "The specified block height not found"},
	ErrNotFoundTransactionHash:           {Code: ErrNotFoundTransactionHash, Message: "The specified transaction hash not found"},
	ErrGetTransactionListCount:           {Code: ErrGetTransactionListCount, Message: "Get transaction list count failed"},
	ErrFetchTransactionByHash:            {Code: ErrFetchTransactionByHash, Message: "Fetch transaction information via hash failed"},
	ErrFetchTransactionListInBlockHeight: {Code: ErrFetchTransactionListInBlockHeight, Message: "Fetch transaction list in the specified block height failed"},
	ErrFetchTransactionListInBlockHash:   {Code: ErrFetchTransactionListInBlockHash, Message: "Fetch transaction list in the specified block hash failed"},
	ErrGetNetworkInfo:                    {Code: ErrGetNetworkInfo, Message: "Get network info failed"},
	ErrGetTotalWormholeTxCount:           {Code: ErrGetTotalWormholeTxCount, Message: "Get total wormhole transactions count failed"},
	ErrGetWormholeTxCountInPeriod:        {Code: ErrGetWormholeTxCountInPeriod, Message: "Get wormhole transactions count in period failed"},
	ErrGetPropertyNameByID:               {Code: ErrGetPropertyNameByID, Message: "Get property name by id failed"},
	ErrFetchLatestWhccountInfo:           {Code: ErrFetchLatestWhccountInfo, Message: "Fetch latest wormhole info failed"},
	ErrGetAddressDetail:				  {Code:ErrGetAddressDetail, Message: "Fetch address balance failed"},
	ErrGetAddressTransactions:			  {Code:ErrGetAddressTransactions, Message: "Fetch address transactions list failed"},
	ErrGetAddressProperties:			  {Code:ErrGetAddressProperties, Message: "Fetch address properties failed"},
	ErrGetAddressPropertyTransactions:	  {Code:ErrGetAddressPropertyTransactions, Message: "Fetch property transactions failed"},
	ErrGetProperties:				      {Code:ErrGetProperties, Message: "Fetch properties failed"},
	ErrGetPropertyByID:				      {Code:ErrGetPropertyByID, Message: "Fetch property detail failed"},
	ErrGetPropertyChangeHistory:		  {Code:ErrGetPropertyChangeHistory, Message: "Fetch property histories list failed"},
	ErrGetPropertyTransactions:			  {Code:ErrGetPropertyTransactions, Message: "Fetch property transactions list failed"},

	// server or handle error
	ErrSetUTCLocation: {Code: ErrSetUTCLocation, Message: "Set UTC time failed"},
}

func apiErrorWithMsg(code errCode, msg string) *Response {
	r, ok := errMapping[code]
	if !ok {
		return errNotFound
	}

	if msg != "" {
		// copy the origin response
		customMsg := *r
		customMsg.Message = msg
		return &customMsg
	}

	return r
}

func apiError(code errCode) *Response {
	r, ok := errMapping[code]
	if !ok {
		return errNotFound
	}

	return r
}

func apiSuccess(data interface{}) *Response {
	return &Response{
		Code:    0,
		Message: "",
		Result:  data,
	}
}
