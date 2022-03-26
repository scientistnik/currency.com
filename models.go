package currencycom

type RejectReasonEnum string

//Enum:
//[ ACCOUNT_NOT_FOUND, CLOSED_MARKET, CLOSE_ONLY, ENGINE_BUSY, HEDGING_MODE_GSL, INSTRUMENT_NOT_AVAILABLE, INSTRUMENT_NOT_FOUND, INVALID_ORDER, INVALID_ORDER_QTY, INVALID_PRICE, LONG_ONLY, OFF_MARKET, ORDER_NOT_FOUND, ORIGINAL_GSL_UPDATE, POSITION_NOT_FOUND, RC_INSTRUMENT_CLIENT_MOP, RC_INSTRUMENT_GLOBAL_MOP, RC_NOT_ENOUGH_MARGIN, RC_NOT_FOUND, RC_NO_RATES, RC_SETTLEMENT, RC_UNKNOWN, REQUIRED_GSL, RISK_CHECK, THROTTLING, UNKNOWN ]

type DtoType string

//Enum:
//[ ORDER_CANCEL, ORDER_MODIFY, ORDER_NEW, POSITION_MODIFY ]

type DtoState string

//Enum:
//[ CANCELLED, PENDING, PROCESSED ]

type OrderSide string

// Enum:
// [ BUY, SELL ]

type OrderStatus string

// Enum:
// [ CANCELED, EXPIRED, FILLED, NEW, PARTIALLY_FILLED, PENDING_CANCEL, REJECTED ]

type OrderTimeInForce string

//Enum:
//[ FOK, GTC, IOC ]

type OrderType string

//Enum:
//[ LIMIT, LIMIT_MAKER, MARKET, STOP, STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, TAKE_PROFIT_LIMIT ]

type CurrencyType string

//_Enum:
//[ CRYPTO, EXCHANGE_TOKEN, FIAT, ICO, TOKEN, TOKENISED_SECURITY, UTILITY_TOKENS ]

type AssetType string

//Enum:
//[ BOND, COMMODITY, CREDIT, CRYPTOCURRENCY, CURRENCY, EQUITY, ICO, INDEX, INTEREST_RATE, OTHER_ASSET, REAL_ESTATE, UTILITY_TOKENS ]

type MarketModes string

//Enum:
//[ CLOSED_FOR_CORPORATE_ACTION, CLOSE_ONLY, HOLIDAY, LONG_ONLY, REGULAR, UNKNOWN, VIEW_AND_REQUEST, VIEW_ONLY ]

type MarketType string

//Enum:
//[ LEVERAGE, SPOT ]

type ExchangeStatus string

//Enum:
//[ AUCTION_MATCH, BREAK, END_OF_DAY, HALT, POST_TRADING, PRE_TRADING, TRADING ]

type PositionDtoState string

//Enum:
//[ ACTIVE, INACTIVE, INVALID ]

type PositionDtoType string

//Enum:
//[ HEDGE, NET ]

type ExchangeType string

//Enum:
//[ GTC, IOC ]

type ReportSource string

//Enum:
//[ CLOSE_OUT, DEALER, SL, SYSTEM, TP, USER ]

type ReportStatus string

//Enum:
//[ CLOSED, DIVIDEND, MODIFIED, MODIFY_REJECT, OPENED, SWAP ]

type AccountBalance struct {
	AccountId          string  `json:"accountId"`
	Asset              string  `json:"asset"`
	CollateralCurrency bool    `json:"collateralCurrency"`
	Default            bool    `json:"default"`
	Free               float64 `json:"free"`
	Locked             float64 `json:"locked"`
}

type AccountRequest struct {
	RecvWindow      int64 //maximum: 60000, exclusiveMaximum: false
	ShowZeroBalance bool
}

type AccountResponse struct {
	AffiliateId      string           `json:"affiliateId"`
	Balances         []AccountBalance `json:"balances"`
	BuyerCommission  float64          `json:"buyerCommission"`
	CanDeposit       bool             `json:"canDeposit"`
	CanTrade         bool             `json:"canTrade"`
	CanWithdraw      bool             `json:"canWithdraw"`
	MakerCommission  float64          `json:"makerCommission"`
	SellerCommission float64          `json:"sellerCommission"`
	TakerCommission  float64          `json:"takerCommission"`
	UpdateTime       int64            `json:"updateTime"`
	UserId           int64            `json:"userId"`
}

type AggTrades struct {
	Timestamp int64  `json:"T"`
	Aggregate int64  `json:"a"`
	IsMaker   bool   `json:"m"`
	Price     string `json:"p"`
	Quantity  string `json:"q"`
}

type AggTradesRequest struct {
	EndTime   int64
	Limit     int32
	StartTime int64
	Symbol    string //*
}

type AggTradesResponse struct {
	AggTrades []AggTrades
}

type AllMyTradesRequest struct {
	EndTime    int64
	Limit      int32
	RecvWindow int64 //maximum: 60000, exclusiveMaximum: false
	StartTime  int64
	Symbol     string //*
}

type AllMyTradesResponse struct {
	MyTrades []MyTradesResponse
}

type BlockchainAddressGetResponse struct {
	Address        string `json:"address"`
	AddressLegacy  string `json:"addressLegacy"`
	DestinationTag string `json:"destinationTag"`
}

type BlockchainAddressRequest struct {
	Coin       string //*
	RecvWindow int64  //maximum: 60000, exclusiveMaximum: false
}

type BySymbolRequest struct {
	Symbol string
}

type CancelOrderRequest struct {
	OrderId    string //*
	RecvWindow int64  //maximum: 60000, exclusiveMaximum: false
	Symbol     string //*
}

type CancelOrderResponse struct {
	ExecutedQty string `json:"executedQty"`
	OrderId     string `json:"orderId"`
	OrigQty     string `json:"origQty"`
	Price       string `json:"price"`
	Side        string `json:"side"`   //OrderSide
	Status      string `json:"status"` //OrderStatus
	Symbol      string `json:"symbol"`
	TimeInForce string `json:"timeInForce"` //OrderTimeInForce
	Type        string `json:"type"`        //OrderType
}

type CloseTradingPositionRequest struct {
	PositionId string //*
	RecvWindow int64  //maximum: 60000, exclusiveMaximum: false
}

type CreateOrderRequest struct {
	AccountId          int64
	ExpireTimestamp    int64
	GuaranteedStopLoss bool
	Leverage           int32
	NewOrderRespType   string
	Price              float64
	Quantity           float64 //*
	RecvWindow         int64   //maximum: 60000, exclusiveMaximum: false
	Side               string  //*
	StopLoss           float64
	Symbol             string //*
	TakeProfit         float64
	Type               string //*
}

type CurrencyDtoResponse struct {
	CommissionFixed   float64 `json:"commissionFixed"`
	CommissionMin     float64 `json:"commissionMin"`
	CommissionPercent float64 `json:"commissionPercent"`
	DisplaySymbol     string  `json:"displaySymbol"`
	MaxWithdrawal     float64 `json:"maxWithdrawal"`
	MinDeposit        float64 `json:"minDeposit"`
	MinWithdrawal     float64 `json:"minWithdrawal"`
	Name              string  `json:"name"`
	Precision         int32   `json:"precision"`
	Type              string  `json:"type"` //CurrencyType
}

type CurrencyResponse struct {
	Currencies []CurrencyDtoResponse `json:"currencies"`
}

type DepthRequest struct {
	Limit  int32
	Symbol string //*
}

type DepthResponse struct {
	Asks         [][]float64 `json:"asks"`
	Bids         [][]float64 `json:"bids"`
	LastUpdateId int64       `json:"lastUpdateId"`
}

type EmptyRequest struct{}
type ExchangeFilter struct{}

type ExchangeInfoResponse struct {
	ExchangeFilters []ExchangeFilter     `json:"exchangeFilters"`
	RateLimits      []RateLimits         `json:"rateLimits"`
	ServerTime      int64                `json:"serverTime"`
	Symbols         []ExchangeSymbolInfo `json:"symbols"`
	Timezone        string               `json:"timezone"`
}

type ExchangeSymbolInfo struct {
	AssetType          string         `json:"assetType"` // AssetType
	BaseAsset          string         `json:"baseAsset"`
	BaseAssetPrecision int32          `json:"baseAssetPrecision"`
	Country            string         `json:"country"`
	ExchangeFee        float64        `json:"exchangeFee"`
	Filters            []SymbolFilter `json:"filters"`
	Industry           string         `json:"industry"`
	LongRate           float64        `json:"longRate"`
	MakerFee           float64        `json:"makerFee"`
	MarketModes        []string       `json:"marketModes"` //MarketModes
	MarketType         string         `json:"marketType"`  //MarketType
	MaxSLGap           float64        `json:"maxSLGap"`
	MaxTPGap           float64        `json:"maxTPGap"`
	MinSLGap           float64        `json:"minSLGap"`
	MinTPGap           float64        `json:"minTPGap"`
	Name               string         `json:"name"`
	OrderTypes         []string       `json:"orderTypes"` //OrderType
	QuoteAsset         string         `json:"quoteAsset"`
	QuoteAssetId       string         `json:"quoteAssetId"`
	QuotePrecision     int32          `json:"quotePrecision"`
	Sector             string         `json:"sector"`
	ShortRate          float64        `json:"shortRate"`
	Status             string         `json:"status"` //ExchangeStatus
	SwapChargeInterval int64          `json:"swapChargeInterval"`
	Symbol             string         `json:"symbol"`
	TakerFee           float64        `json:"takerFee"`
	TickSize           float64        `json:"tickSize"`
	TickValue          float64        `json:"tickValue"`
	TradingFee         float64        `json:"tradingFee"`
	TradingHours       string         `json:"tradingHours"`
}

type InternalQuote struct {
	bid        float64
	bidQty     float64
	ofr        float64
	ofrQty     float64
	symbolName string
	timestamp  int64
}

type KLinesRequest struct {
	EndTime   int64
	Interval  string //*
	Limit     int32
	StartTime int64
	Symbol    string //*
	Type      string
}

type KLinesResponse struct {
	Lines [][6]interface{}
}

type KLinesResponseStruct struct {
	OpenTime float64
	Open     string
	High     string
	Low      string
	Close    string
	Volume   float64
}

type LeverageSettingsRequest struct {
	RecvWindow int64  //maximum: 60000, exclusiveMaximum: false
	Symbol     string //*
}

type LeverageSettingsResponse struct {
	Value  int32   `json:"value"`
	Values []int32 `json:"values"`
}

type MarketDepthData struct {
	bid map[string]float64
	ofr map[string]float64
	ts  int64
}

type MarketDepthEvent struct {
	data   MarketDepthData
	symbol string
}

type MyTradesResponse struct {
	Buyer           bool   `json:"buyer"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Id              string `json:"id"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	Maker           bool   `json:"maker"`
	OrderId         string `json:"orderId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	QuoteQty        string `json:"quoteQty"`
	Symbol          string `json:"symbol"`
	Time            int64  `json:"time"`
}

type NewOrderResponseRESULT struct {
	ExecutedQty        string  `json:"executedQty"`
	ExpireTimestamp    int64   `json:"expireTimestamp"`
	GuaranteedStopLoss bool    `json:"guaranteedStopLoss"`
	Margin             float64 `json:"margin"`
	OrderId            string  `json:"orderId"`
	OrigQty            string  `json:"origQty"`
	Price              string  `json:"price"`
	RejectMessage      string  `json:"rejectMessage"`
	Side               string  `json:"side"` //OrderSide
	StopLoss           float64 `json:"stopLoss"`
	Symbol             string  `json:"symbol"`
	TakeProfit         float64 `json:"takeProfit"`
	TimeInForce        string  `json:"timeInForce"` //OrderTimeInForce
	TransactTime       int64   `json:"transactTime"`
	Type               string  `json:"type"` //OrderType
}

type OHLCBar struct {
	c        float64
	h        float64
	interval string
	l        float64
	o        float64
	symbol   string
	t        int64
	_type    string
}

type OHLCSubscribeRequest struct {
	intervals []string //Identifies intervals for subscription. Available: 1m, 5m, 15m, 30m, 1h, 4h, 1d, 1w. Default: 1m.
	symbols   []string //Identifies symbols for subscription.
	_type     string   //Type of candlestick. Available: classic, heikin-ashi.
}

type OpenOrdersReponse struct {
	openOrders []QueryOrderResponse
}

type PingRequest struct{}
type PingResponse struct{}

type PositionDto struct {
	AccountId          string  `json:"accountId"`     //*
	ClosePrice         float64 `json:"closePrice"`    //*
	CloseQuantity      float64 `json:"closeQuantity"` //*
	CloseTimestamp     int64   `json:"closeTimestamp"`
	Cost               float64 `json:"cost"`
	CreatedTimestamp   int64   `json:"createdTimestamp"` //*
	Currency           string  `json:"currency"`         //*
	Dividend           float64 `json:"dividend"`
	Fee                float64 `json:"fee"`
	GuaranteedStopLoss bool    `json:"guaranteedStopLoss"`
	Id                 string  `json:"id"`            //* uuid
	InstrumentId       int64   `json:"instrumentId"`  //*
	Margin             float64 `json:"margin"`        //*
	OpenPrice          float64 `json:"openPrice"`     //*
	OpenQuantity       float64 `json:"openQuantity"`  //*
	OpenTimestamp      int64   `json:"openTimestamp"` //*
	OrderId            string  `json:"orderId"`       //* uuid
	Rpl                float64 `json:"rpl"`
	RplConverted       float64 `json:"rplConverted"`
	State              string  `json:"state"` //* PositionDtoState
	StopLoss           float64 `json:"stopLoss"`
	Swap               float64 `json:"swap"`
	SwapConverted      float64 `json:"swapConverted"`
	Symbol             string  `json:"symbol"`
	TakeProfit         float64 `json:"takeProfit"`
	Type               string  `json:"type"` //PositionDtoType
	Upl                float64 `json:"upl"`
	UplConverted       float64 `json:"uplConverted"`
}

type PositionExecutionReportDto struct {
	AccountCurrency  string             `json:"accountCurrency"`  //*
	AccountId        int64              `json:"accountId"`        //*
	CreatedTimestamp int64              `json:"createdTimestamp"` //*
	Currency         string             `json:"currency"`         //*
	ExecTimestamp    int64              `json:"execTimestamp"`    //*
	ExecutionType    string             `json:"executionType"`    //ExchangeType
	Fee              float64            `json:"fee"`
	FeeDetails       map[string]float64 `json:"feeDetails"`
	FxRate           float64            `json:"fxRate"`
	GSL              bool               `json:"gSL"`
	InstrumentId     int64              `json:"instrumentId"` //*
	PositionId       string             `json:"positionId"`   //*
	Price            float64            `json:"price"`
	Quantity         float64            `json:"quantity"`
	RejectReason     string             `json:"rejectReason"` //RejectReasonEnum
	Rpl              float64            `json:"rpl"`
	RplConverted     float64            `json:"rplConverted"`
	Source           string             `json:"source"` //* ReportSource
	Status           string             `json:"status"` //* ReportStatus
	StopLoss         float64            `json:"stopLoss"`
	Swap             float64            `json:"swap"`
	SwapConverted    float64            `json:"swapConverted"`
	Symbol           string             `json:"symbol"`
	TakeProfit       float64            `json:"takeProfit"`
}

type PositionHistoryRequest struct {
	Limit      int32
	RecvWindow int64 //maximum: 60000, exclusiveMaximum: false
	Symbol     string
}

type QueryOrderResponse struct {
	AccountId          string  `json:"accountId"`
	ExecutedQty        string  `json:"executedQty"`
	ExpireTimestamp    int64   `json:"expireTimestamp"`
	QuaranteedStopLoss bool    `json:"guaranteedStopLoss"`
	IcebergQty         string  `json:"icebergQty"`
	Leverage           bool    `json:"leverage"`
	Margin             float64 `json:"margin"`
	OrderId            string  `json:"orderId"`
	OrigQty            string  `json:"origQty"`
	Price              string  `json:"price"`
	Side               string  `json:"side"`   //OrderSide
	Status             string  `json:"status"` //OrderStatus
	StopLoss           float64 `json:"stopLoss"`
	Symbol             string  `json:"symbol"`
	TakeProfit         float64 `json:"takeProfit"`
	Time               int64   `json:"time"`
	TimeInForce        string  `json:"timeInForce"` //OrderTimeInForce
	Type               string  `json:"type"`        //OrderType
	UpdateTime         int64   `json:"updateTime"`
	Working            bool    `json:"working"`
}

type RateLimits struct {
	Interval      string `json:"interval"`
	IntervalNum   int32  `json:"intervalNum"`
	Limit         int32  `json:"limit"`
	RateLimitType string `json:"rateLimitType"`
}

type RequestDto struct {
	AccountId        string `json:"accountId"`        //*
	CreatedTimestamp int64  `json:"createdTimestamp"` //*
	Id               int64  `json:"id"`               //*
	InstrumentId     int64  `json:"instrumentId"`     //*
	OrderId          string `json:"orderId"`
	PositionId       string `json:"positionId"`
	RejectReason     string `json:"rejectReason"` //RejectReasonEnum
	RqBody           string `json:"rqBody"`       //*
	RqType           string `json:"rqType"`       //*
	State            string `json:"state"`        //*
}

type ServerTimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}

type SignedBySymbolRequest struct {
	apiKey     string //*
	recvWindow int64  //maximum: 60000, exclusiveMaximum: false
	signature  string //*
	symbol     string
	timestamp  int64 //*
}

type SignedRequest struct {
	//apiKey     string //*
	RecvWindow int64 //maximum: 60000, exclusiveMaximum: false
	//signature  string //*
	//timestamp  int64  //*
}

type SubscribeRequest struct {
	symbols []string //Identifies symbols for subscription.
}

type SubscribeResponse struct {
	subscriptions map[string]string
}

type SymbolFilter struct {
	FilterType string `json:"filterType"`
}

type Ticker24HResponse struct {
	tickers []Ticker24hr
}

type Ticker24hr struct {
	AskPrice           string `json:"askPrice"`
	BidPrice           string `json:"bidPrice"`
	CloseTime          int64  `json:"closeTime"`
	HighPrice          string `json:"highPrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	LowPrice           string `json:"lowPrice"`
	OpenPrice          string `json:"openPrice"`
	OpenTime           int64  `json:"openTime"`
	PrevClosePrice     string `json:"prevClosePrice"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	QuoteVolume        string `json:"quoteVolume"`
	Symbol             string `json:"symbol"`
	Volume             string `json:"volume"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
}

type TradeEventReq struct {
	id      int32
	orderId string
	price   float64
	size    float64
	symbol  string
	ts      int64
}

type TradeEventRes struct {
	buyer   bool
	id      int32
	orderId string
	price   float64
	size    float64
	symbol  string
	ts      int64
}

type TradingOrderUpdateResponse struct {
	RequestId int64  `json:"requestId"` //*
	State     string `json:"state"`     //* DtoState
}

type TradingPositionCloseAllResponse struct {
	Request []RequestDto `json:"request"`
}

type TradingPositionHistoryResponse struct {
	History []PositionExecutionReportDto `json:"history"`
}

type TradingPositionListResponse struct {
	Positions []PositionDto `json:"positions"`
}

type TradingPositionUpdateResponse struct {
	RequestId int64  `json:"requestId"` //*
	State     string `json:"state"`     //* DtoState
}

type TransactionDTOResponse struct {
	Amount                    float64 `json:"amount"`
	Balance                   float64 `json:"balance"`
	BlockchainTransactionHash string  `json:"blockchainTransactionHash"`
	Commission                float64 `json:"commission"`
	Currency                  string  `json:"currency"`
	Id                        int64   `json:"id"`
	PaymentMethod             string  `json:"paymentMethod"`
	Status                    string  `json:"status"`
	Timestamp                 int64   `json:"timestamp"`
	Type                      string  `json:"type"`
}

type TransactionsRequest struct {
	EndTime    int64
	Limit      int32
	RecvWindow int64 //maximum: 60000, exclusiveMaximum: false
	StartTime  int64
}

type TransactionsResponse struct {
	transactions []TransactionDTOResponse
}

type UpdateTradingOrderRequest struct {
	ExpireTimestamp    int64
	GuaranteedStopLoss bool
	NewPrice           float64
	OrderId            string //*
	RecvWindow         int64  //maximum: 60000, exclusiveMaximum: false
	StopLoss           float64
	TakeProfit         float64
}

type UpdateTradingPositionRequest struct {
	GuaranteedStopLoss bool
	PositionId         string //*
	RecvWindow         int64  //maximum: 60000, exclusiveMaximum: false
	StopLoss           float64
	TakeProfit         float64
}
