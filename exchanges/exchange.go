package exchange

import (
	"log"
	"time"

	"github.com/thrasher-/gocryptotrader/common"
	"github.com/thrasher-/gocryptotrader/config"
	"github.com/thrasher-/gocryptotrader/currency/pair"
	"github.com/thrasher-/gocryptotrader/exchanges/orderbook"
	"github.com/thrasher-/gocryptotrader/exchanges/ticker"
)

const (
	WarningBase64DecryptSecretKeyFailed = "WARNING -- Exchange %s unable to base64 decode secret key.. Disabling Authenticated API support."
	ErrExchangeNotFound                 = "Exchange not found in dataset."
)

//ExchangeAccountInfo : Generic type to hold each exchange's holdings in all enabled currencies
type ExchangeAccountInfo struct {
	ExchangeName string
	Currencies   []ExchangeAccountCurrencyInfo
}

//ExchangeAccountCurrencyInfo : Sub type to store currency name and value
type ExchangeAccountCurrencyInfo struct {
	CurrencyName string
	TotalValue   float64
	Hold         float64
}

type ExchangeBase struct {
	Name                        string
	Enabled                     bool
	Verbose                     bool
	Websocket                   bool
	RESTPollingDelay            time.Duration
	AuthenticatedAPISupport     bool
	APISecret, APIKey, ClientID string
	TakerFee, MakerFee, Fee     float64
	BaseCurrencies              []string
	AvailablePairs              []string
	EnabledPairs                []string
	WebsocketURL                string
	APIUrl                      string
}

//IBotExchange : Enforces standard functions for all exchanges supported in gocryptotrader
type IBotExchange interface {
	Setup(exch config.ExchangeConfig)
	Start()
	SetDefaults()
	GetName() string
	IsEnabled() bool
	GetTickerPrice(currency pair.CurrencyPair) (ticker.TickerPrice, error)
	GetOrderbookEx(currency pair.CurrencyPair) (orderbook.OrderbookBase, error)
	GetEnabledCurrencies() []string
	GetExchangeAccountInfo() (ExchangeAccountInfo, error)
}

func (e *ExchangeBase) GetName() string {
	return e.Name
}
func (e *ExchangeBase) GetEnabledCurrencies() []string {
	return e.EnabledPairs
}
func (e *ExchangeBase) SetEnabled(enabled bool) {
	e.Enabled = enabled
}

func (e *ExchangeBase) IsEnabled() bool {
	return e.Enabled
}

func (e *ExchangeBase) SetAPIKeys(APIKey, APISecret, ClientID string, b64Decode bool) {
	e.APIKey = APIKey
	e.ClientID = ClientID

	if b64Decode {
		result, err := common.Base64Decode(APISecret)
		if err != nil {
			e.AuthenticatedAPISupport = false
			log.Printf(WarningBase64DecryptSecretKeyFailed, e.Name)
		}
		e.APISecret = string(result)
	} else {
		e.APISecret = APISecret
	}
}

func (e *ExchangeBase) UpdateAvailableCurrencies(exchangeProducts []string) error {
	exchangeProducts = common.SplitStrings(common.StringToUpper(common.JoinStrings(exchangeProducts, ",")), ",")
	diff := common.StringSliceDifference(e.AvailablePairs, exchangeProducts)
	if len(diff) > 0 {
		cfg := config.GetConfig()
		exch, err := cfg.GetExchangeConfig(e.Name)
		if err != nil {
			return err
		} else {
			log.Printf("%s Updating available pairs. Difference: %s.\n", e.Name, diff)
			exch.AvailablePairs = common.JoinStrings(exchangeProducts, ",")
			cfg.UpdateExchangeConfig(exch)
		}
	}
	return nil
}
