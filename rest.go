package currencycom

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const DEFAULT_ENDPOINT string = "https://api-adapter.backend.currency.com"
const VERSION_API string = "v2"

var baseEndpoint string

func init() {
	baseEndpoint = DEFAULT_ENDPOINT
}

type RestAPI struct {
	apiKey   string
	secret   string
	Endpoint string
}

type requestArgs struct {
	httpMethod string
	endpoint   string
	methodName string
	params     map[string]string
	restApi    *RestAPI
}

func request(args *requestArgs) ([]byte, error) {
	url := args.endpoint + "/api/" + VERSION_API + "/" + args.methodName

	client := &http.Client{}
	req, err := http.NewRequest(args.httpMethod, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error in prepare request, %w", err)
	}

	query := req.URL.Query()
	for key, value := range args.params {
		query.Add(key, value)
	}

	if args.restApi != nil {
		req.Header.Set("X-MBX-APIKEY", args.restApi.apiKey)

		query.Add("timestamp", strconv.Itoa(int(time.Now().UnixMilli())))

		sig := hmac.New(sha256.New, []byte(args.restApi.secret))
		sig.Write([]byte(query.Encode()))
		query.Add("signature", hex.EncodeToString(sig.Sum(nil)))
	}

	req.URL.RawQuery = query.Encode()
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error in call, %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error in read, %w", err)
	}

	if resp.StatusCode >= 400 {
		return body, fmt.Errorf("bad response from server, " + resp.Status + ": " + string(body))
	}

	return body, nil
}

func ServerTime() (*ServerTimeResponse, error) {
	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   baseEndpoint,
		methodName: "time",
		params:     nil,
		restApi:    nil,
	})
	if err != nil {
		return nil, err
	}

	var out ServerTimeResponse
	err = json.Unmarshal(body, &out)

	return &out, err
}

func TradesAggregated(params *AggTradesRequest) ([]AggTrades, error) {
	if params == nil || params.Symbol == "" {
		return nil, fmt.Errorf("error params: Symbol need to set")
	}

	reqParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.EndTime != 0 {
		reqParams["endTime"] = strconv.FormatUint(uint64(params.EndTime), 10)
	}

	if params.Limit != 0 {
		reqParams["limit"] = strconv.FormatUint(uint64(params.Limit), 10)
	}

	if params.StartTime != 0 {
		reqParams["startTime"] = strconv.FormatUint(uint64(params.StartTime), 10)
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   baseEndpoint,
		methodName: "aggTrades",
		params:     reqParams,
		restApi:    nil,
	})
	if err != nil {
		return nil, err
	}

	var out []AggTrades
	err = json.Unmarshal(body, &out)

	return out, err
}

func OrderBook(params *DepthRequest) (*DepthResponse, error) {
	if params == nil || params.Symbol == "" {
		return nil, fmt.Errorf("error params: Symbol need to set")
	}

	reqParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.Limit != 0 {
		reqParams["limit"] = strconv.FormatUint(uint64(params.Limit), 10)
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   baseEndpoint,
		methodName: "depth",
		params:     reqParams,
		restApi:    nil,
	})
	if err != nil {
		return nil, err
	}

	var out DepthResponse
	err = json.Unmarshal(body, &out)

	return &out, err
}

func ExchangeInfo() (*ExchangeInfoResponse, error) {
	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   baseEndpoint,
		methodName: "exchangeInfo",
		params:     nil,
		restApi:    nil,
	})
	if err != nil {
		return nil, err
	}

	var out ExchangeInfoResponse
	err = json.Unmarshal(body, &out)

	return &out, err
}

func Klines(params *KLinesRequest) ([]KLinesResponseStruct, error) {
	if params == nil {
		return nil, fmt.Errorf("error params: Symbol and Interval need to set")
	}

	reqParams := map[string]string{
		"symbol":   params.Symbol,
		"interval": params.Interval,
	}

	if params.StartTime != 0 {
		reqParams["startTime"] = strconv.FormatUint(uint64(params.StartTime), 10)
	}

	if params.EndTime != 0 {
		reqParams["endTime"] = strconv.FormatUint(uint64(params.EndTime), 10)
	}

	if params.Limit != 0 {
		reqParams["limit"] = strconv.FormatUint(uint64(params.Limit), 10)
	}

	if params.Type != "" {
		reqParams["type"] = params.Type
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   baseEndpoint,
		methodName: "klines",
		params:     reqParams,
		restApi:    nil,
	})
	if err != nil {
		return nil, err
	}

	var middle [][6]interface{}
	err = json.Unmarshal(body, &middle)

	out := make([]KLinesResponseStruct, len(middle))
	for _, line := range middle {
		out = append(out, KLinesResponseStruct{
			OpenTime: line[0].(float64),
			Open:     line[1].(string),
			High:     line[2].(string),
			Low:      line[3].(string),
			Close:    line[4].(string),
			Volume:   line[5].(float64),
		})
	}

	return out, err
}

func PriceChange(params *BySymbolRequest) (*Ticker24hr, error) {
	if params == nil || params.Symbol == "" {
		return nil, fmt.Errorf("error params: Symbol need to set")
	}

	reqParams := map[string]string{
		"symbol": params.Symbol,
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   baseEndpoint,
		methodName: "ticker/24hr",
		params:     reqParams,
		restApi:    nil,
	})
	if err != nil {
		return nil, err
	}

	var out Ticker24hr
	err = json.Unmarshal(body, &out)

	return &out, err
}

func NewRestAPI(apiKey string, secret string, endpoint string) *RestAPI {
	if endpoint == "" {
		endpoint = baseEndpoint
	}

	return &RestAPI{apiKey: apiKey, secret: secret, Endpoint: endpoint}
}

func (r RestAPI) AccountInfo(params *AccountRequest) (*AccountResponse, error) {
	var reqParams map[string]string

	if params != nil {
		reqParams = map[string]string{
			"showZeroBalance": strconv.FormatBool(params.ShowZeroBalance),
		}

		if params.RecvWindow != 0 {
			reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
		}
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   r.Endpoint,
		methodName: "account",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out AccountResponse

	err = json.Unmarshal(body, &out)

	return &out, err
}

func (r RestAPI) TradingPositionClose(params *CloseTradingPositionRequest) (*TradingPositionCloseAllResponse, error) {
	if params == nil || params.PositionId == "" {
		return nil, fmt.Errorf("error params: PositionId need to set")
	}

	reqParams := map[string]string{
		"positionId": params.PositionId,
	}

	if params.RecvWindow != 0 {
		reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
	}

	body, err := request(&requestArgs{
		httpMethod: "POST",
		endpoint:   r.Endpoint,
		methodName: "closeTradingPosition",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out TradingPositionCloseAllResponse

	err = json.Unmarshal(body, &out)

	return &out, err
}

func (r RestAPI) ListOfCurrencies(params *SignedRequest) ([]CurrencyDtoResponse, error) {
	var reqParams map[string]string

	if params != nil {
		reqParams = map[string]string{
			"recvWindow": strconv.FormatUint(uint64(params.RecvWindow), 10),
		}
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   r.Endpoint,
		methodName: "currencies",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out []CurrencyDtoResponse

	err = json.Unmarshal(body, &out)

	return out, err
}

func (r RestAPI) StringOfAddress(params *BlockchainAddressRequest) (*BlockchainAddressGetResponse, error) {
	if params == nil || params.Coin == "" {
		return nil, fmt.Errorf("error params: Coin need to set")
	}

	reqParams := map[string]string{
		"coin": params.Coin,
	}

	if params.RecvWindow != 0 {
		reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   r.Endpoint,
		methodName: "depositAddress",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out BlockchainAddressGetResponse

	err = json.Unmarshal(body, &out)

	return &out, err
}

func (r RestAPI) ListOfDeposits(params *TransactionsRequest) ([]TransactionDTOResponse, error) {
	var reqParams map[string]string

	if params != nil {
		reqParams = make(map[string]string)

		if params.RecvWindow != 0 {
			reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
		}

		if params.StartTime != 0 {
			reqParams["startTime"] = strconv.FormatUint(uint64(params.StartTime), 10)
		}

		if params.EndTime != 0 {
			reqParams["endTime"] = strconv.FormatUint(uint64(params.EndTime), 10)
		}

		if params.Limit != 0 {
			reqParams["limit"] = strconv.FormatUint(uint64(params.Limit), 10)
		}
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   r.Endpoint,
		methodName: "deposits",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out []TransactionDTOResponse

	err = json.Unmarshal(body, &out)

	return out, err
}

func (r RestAPI) ListOfLedgers(params *TransactionsRequest) ([]TransactionDTOResponse, error) {
	var reqParams map[string]string

	if params != nil {
		reqParams = make(map[string]string)

		if params.StartTime != 0 {
			reqParams["startTime"] = strconv.FormatUint(uint64(params.StartTime), 10)
		}

		if params.EndTime != 0 {
			reqParams["endTime"] = strconv.FormatUint(uint64(params.EndTime), 10)
		}

		if params.Limit != 0 {
			reqParams["limit"] = strconv.FormatUint(uint64(params.Limit), 10)
		}

		if params.RecvWindow != 0 {
			reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
		}
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   r.Endpoint,
		methodName: "ledger",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out []TransactionDTOResponse

	err = json.Unmarshal(body, &out)

	return out, err
}

func (r RestAPI) LeverageSettings(params *LeverageSettingsRequest) (*LeverageSettingsResponse, error) {
	if params == nil || params.Symbol == "" {
		return nil, fmt.Errorf("error params: Symbol need to set")
	}

	reqParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.RecvWindow != 0 {
		reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   r.Endpoint,
		methodName: "leverageSettings",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out LeverageSettingsResponse

	err = json.Unmarshal(body, &out)

	return &out, err
}

func (r RestAPI) ListOfTrades(params *AllMyTradesRequest) ([]MyTradesResponse, error) {
	if params == nil || params.Symbol == "" {
		return nil, fmt.Errorf("error params: Symbol need to set")
	}

	reqParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.StartTime != 0 {
		reqParams["startTime"] = strconv.FormatUint(uint64(params.StartTime), 10)
	}

	if params.EndTime != 0 {
		reqParams["endTime"] = strconv.FormatUint(uint64(params.EndTime), 10)
	}

	if params.Limit != 0 {
		reqParams["limit"] = strconv.FormatUint(uint64(params.Limit), 10)
	}

	if params.RecvWindow != 0 {
		reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   r.Endpoint,
		methodName: "myTrades",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out []MyTradesResponse

	err = json.Unmarshal(body, &out)

	return out, err
}

func (r RestAPI) ListOfOpenOrder(params *PositionHistoryRequest) ([]QueryOrderResponse, error) {
	var reqParams map[string]string

	if params != nil {
		reqParams = make(map[string]string)

		if params.Symbol != "" {
			reqParams["symbol"] = params.Symbol
		}

		if params.RecvWindow != 0 {
			reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
		}
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   r.Endpoint,
		methodName: "openOrders",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out []QueryOrderResponse

	err = json.Unmarshal(body, &out)

	return out, err
}

func (r RestAPI) CreateOrder(params *CreateOrderRequest) (*NewOrderResponseRESULT, error) {
	if params == nil {
		return nil, fmt.Errorf("error params: Symbol, Quantity, Side, Type need to set")
	}

	reqParams := map[string]string{
		"symbol":   params.Symbol,
		"quantity": strconv.FormatFloat(params.Quantity, 'f', 8, 64),
		"side":     params.Side,
		"type":     params.Type,
	}

	if params.AccountId != 0 {
		reqParams["accountId"] = strconv.FormatUint(uint64(params.AccountId), 10)
	}

	if params.ExpireTimestamp != 0 {
		reqParams["expireTimestamp"] = strconv.FormatUint(uint64(params.ExpireTimestamp), 10)
	}

	if params.GuaranteedStopLoss {
		reqParams["guaranteedStopLoss"] = strconv.FormatBool(params.GuaranteedStopLoss)
	}

	if params.Leverage != 0 {
		reqParams["leverage"] = strconv.FormatInt(int64(params.Leverage), 10)
	}

	if params.NewOrderRespType != "" {
		reqParams["newOrderRespType"] = params.NewOrderRespType
	}

	if params.Price != 0 {
		reqParams["price"] = strconv.FormatFloat(params.Price, 'f', 8, 64)
	}

	if params.RecvWindow != 0 {
		reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
	}

	if params.StopLoss != 0 {
		reqParams["stopLoss"] = strconv.FormatFloat(params.StopLoss, 'f', 8, 64)
	}

	if params.TakeProfit != 0 {
		reqParams["takeProfit"] = strconv.FormatFloat(params.TakeProfit, 'f', 8, 64)
	}

	body, err := request(&requestArgs{
		httpMethod: "POST",
		endpoint:   r.Endpoint,
		methodName: "order",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out NewOrderResponseRESULT

	err = json.Unmarshal(body, &out)

	return &out, err
}

func (r RestAPI) CancelOrder(params *CancelOrderRequest) (*CancelOrderResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("error params: Symbol, OrderId need to set")
	}

	reqParams := map[string]string{
		"symbol":  params.Symbol,
		"orderId": params.OrderId,
	}

	if params.RecvWindow != 0 {
		reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
	}

	body, err := request(&requestArgs{
		httpMethod: "DELETE",
		endpoint:   r.Endpoint,
		methodName: "order",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out CancelOrderResponse

	err = json.Unmarshal(body, &out)

	return &out, err
}

func (r RestAPI) ListOfLeverageTrades(params *SignedRequest) (*TradingPositionListResponse, error) {
	var reqParams map[string]string

	if params != nil && params.RecvWindow != 0 {
		reqParams = make(map[string]string)
		reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   r.Endpoint,
		methodName: "tradingPositions",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out TradingPositionListResponse

	err = json.Unmarshal(body, &out)

	return &out, err
}

func (r RestAPI) ListOfHistoricalPositions(params *PositionHistoryRequest) (*TradingPositionHistoryResponse, error) {
	var reqParams map[string]string

	if params != nil {
		reqParams = make(map[string]string)

		if params.RecvWindow != 0 {
			reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
		}

		if params.Symbol != "" {
			reqParams["symbol"] = params.Symbol
		}

		if params.Limit != 0 {
			reqParams["limit"] = strconv.FormatInt(int64(params.Limit), 10)
		}
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   r.Endpoint,
		methodName: "tradingPositionsHistory",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out TradingPositionHistoryResponse

	err = json.Unmarshal(body, &out)

	return &out, err
}

func (r RestAPI) ListOfTransactions(params *TransactionsRequest) ([]TransactionDTOResponse, error) {
	var reqParams map[string]string

	if params != nil {
		reqParams = make(map[string]string)

		if params.RecvWindow != 0 {
			reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
		}

		if params.StartTime != 0 {
			reqParams["startTime"] = strconv.FormatInt(params.StartTime, 10)
		}

		if params.EndTime != 0 {
			reqParams["endTime"] = strconv.FormatInt(params.EndTime, 10)
		}

		if params.Limit != 0 {
			reqParams["limit"] = strconv.FormatInt(int64(params.Limit), 10)
		}
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   r.Endpoint,
		methodName: "transactions",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out []TransactionDTOResponse

	err = json.Unmarshal(body, &out)

	return out, err
}

func (r RestAPI) LeverageOrdersEdit(params *UpdateTradingOrderRequest) (*TradingOrderUpdateResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("error params: OrderId need to set")
	}

	reqParams := map[string]string{
		"orderId": params.OrderId,
	}

	if params.RecvWindow != 0 {
		reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
	}

	if params.ExpireTimestamp != 0 {
		reqParams["expireTimestamp"] = strconv.FormatInt(params.ExpireTimestamp, 10)
	}

	if params.GuaranteedStopLoss {
		reqParams["guaranteedStopLoss"] = strconv.FormatBool(params.GuaranteedStopLoss)
	}

	if params.NewPrice != 0 {
		reqParams["newPrice"] = strconv.FormatFloat(params.NewPrice, 'f', 8, 64)
	}

	if params.StopLoss != 0 {
		reqParams["stopLoss"] = strconv.FormatFloat(params.StopLoss, 'f', 8, 64)
	}

	if params.TakeProfit != 0 {
		reqParams["takeProfit"] = strconv.FormatFloat(params.TakeProfit, 'f', 8, 64)
	}

	body, err := request(&requestArgs{
		httpMethod: "POST",
		endpoint:   r.Endpoint,
		methodName: "updateTradingOrder",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out TradingOrderUpdateResponse

	err = json.Unmarshal(body, &out)

	return &out, err
}

func (r RestAPI) LeverageTradeEdit(params *UpdateTradingPositionRequest) (*TradingPositionUpdateResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("error params: PositionId need to set")
	}

	reqParams := map[string]string{
		"positionId": params.PositionId,
	}

	if params.RecvWindow != 0 {
		reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
	}

	if params.GuaranteedStopLoss {
		reqParams["guaranteedStopLoss"] = strconv.FormatBool(params.GuaranteedStopLoss)
	}

	if params.StopLoss != 0 {
		reqParams["stopLoss"] = strconv.FormatFloat(params.StopLoss, 'f', 8, 64)
	}

	if params.TakeProfit != 0 {
		reqParams["takeProfit"] = strconv.FormatFloat(params.TakeProfit, 'f', 8, 64)
	}

	body, err := request(&requestArgs{
		httpMethod: "POST",
		endpoint:   r.Endpoint,
		methodName: "updateTradingPosition",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out TradingPositionUpdateResponse

	err = json.Unmarshal(body, &out)

	return &out, err
}

func (r RestAPI) ListOfWithdrawals(params *TransactionsRequest) ([]TransactionDTOResponse, error) {
	var reqParams map[string]string

	if params != nil {
		reqParams = make(map[string]string)

		if params.RecvWindow != 0 {
			reqParams["recvWindow"] = strconv.FormatUint(uint64(params.RecvWindow), 10)
		}

		if params.StartTime != 0 {
			reqParams["startTime"] = strconv.FormatInt(params.StartTime, 10)
		}

		if params.EndTime != 0 {
			reqParams["endTime"] = strconv.FormatInt(params.EndTime, 10)
		}

		if params.Limit != 0 {
			reqParams["limit"] = strconv.FormatInt(int64(params.Limit), 10)
		}
	}

	body, err := request(&requestArgs{
		httpMethod: "GET",
		endpoint:   r.Endpoint,
		methodName: "withdrawals",
		params:     reqParams,
		restApi:    &r,
	})
	if err != nil {
		return nil, err
	}

	var out []TransactionDTOResponse

	err = json.Unmarshal(body, &out)

	return out, err
}
