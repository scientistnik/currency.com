# currency.com

Libraty for work with currency.com exchange via [official API](https://currency.com/api). Implemented full REST API.

## Install

Run `go get github.com/scientistnik/currency.com`

## How to use

Create RestAPI instance ([how to get API key and secret](https://currency.com/api-get-started)) :

```go
import currencycom "github.com/scientistnik/currency.com"

const (
  ApiKey string = "<your_api_key>"
  Secret string = "<your_secret>"
  EndPoint string = ""
)

api := currencycom.NewRestAPI(ApiKey, Secret, EndPoint)

```

Look to [official swagger API](https://apitradedoc.currency.com/swagger-ui.html#/)

```go
serverTime, err := api.ServerTime()
accountInfo, err := api.AccountInfo(nil)
trades, err := currency.TradesAggregated(&currencycom.AggTradesRequest{Symbol: "BTC/USD"})
body, err := api.ListOfCurrencies(nil)
body, err := api.StringOfAddress(&currencycom.BlockchainAddressRequest{Coin: "BTC"})
body, err := api.ListOfDeposits(nil)
body, err := currencycom.OrderBook(&currencycom.DepthRequest{Symbol: "BTC/USD", Limit: 2})
body, err := currencycom.Exchangeinfo()
body, err := currencycom.Klines(&currencycom.KLinesRequest{Symbol: "BTC/USD", Interval: "1m"})
body, err := api.ListOfLedgers(nil)
body, err := api.LeverageSettings(&currencycom.LeverageSettingsRequest{Symbol: "BTC/USD_LEVERAGE"})
body, err := api.ListOfTrades(&currencycom.AllMyTradesRequest{Symbol: "BTC/USD"})
body, err := api.ListOfOpenOrder(&currency.PositionHistoryRequest{Symbol: "LTC/USD"})
body, err := api.CreateOrder()
body, err := api.CancelOrder()
body, err := currencycom.PriceChange(&currencycom.BySymbolRequest{Symbol: "BTC/USD"})
body, err := api.ListOfLeverageTrades(&currencycom.SignedRequest{})
body, err := api.ListOfHistoricalPositions(&currencycom.PositionHistoryRequest{})
body, err := api.ListOfTransactions(&currencycom.TransactionsRequest{})
body, err := api.ListOfWithdrawals(&currencycom.TransactionsRequest{})
```
