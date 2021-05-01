package waxpeer

import "time"

type accountInformationResponse struct {
	Success bool                `json:"success"`
	User    *accountInformation `json:"user"`
}

type accountInformation struct {
	Wallet     int64       `json:"wallet"`
	ID         string      `json:"id"`
	UserID     string      `json:"user_id"`
	ID64       string      `json:"id64"`
	BtcWallet  string      `json:"btc_wallet"`
	UsdtWallet string      `json:"usdt_wallet"`
	Avatar     string      `json:"avatar"`
	Proxy      interface{} `json:"proxy"`
	Rank       int64       `json:"rank"`
	Shop       string      `json:"shop"`
	Ref        interface{} `json:"ref"`
	Name       string      `json:"name"`
	SellStatus bool        `json:"sell_status"`
	SellFees   float64     `json:"sell_fees"`
	CanP2P     bool        `json:"can_p2p"`
	Login      string      `json:"login"`
	EthWallet  string      `json:"eth_wallet"`
	Tradelink  string      `json:"tradelink"`
}

type orderOpenResponse struct {
	Success bool         `json:"success"`
	Offers  []*orderOpen `json:"offers"`
	Count   int64        `json:"count"`
}

type orderOpen struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Price  string `json:"price"`
	Amount int64  `json:"amount"`
	Filled int64  `json:"filled"`
}

type orderRemoveResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Removed int64  `json:"removed"`
}

type orderRemoveAllresponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Count   int64  `json:"count"`
}

type orderHistoryResponse struct {
	Success bool            `json:"success"`
	History []*orderHistory `json:"history"`
	Count   int64           `json:"count"`
}

type orderHistory struct {
	ID          int64     `json:"id"`
	ItemName    string    `json:"item_name"`
	Price       string    `json:"price"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
}

type accountSetSteamApiKeyResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type orderEditResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	ID      int64  `json:"id"`
	Price   int64  `json:"price"`
	Amount  int64  `json:"amount"`
}

type orderCreateResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	ID      int64  `json:"id"`
	Filled  int    `json:"filled"`
}

type accountSetTradelinkResponse struct {
	Success   bool   `json:"success"`
	Link      string `json:"link"`
	Token     string `json:"token"`
	Steamid32 string `json:"steamid32"`
}

type transferResponse struct {
	Success bool   `json:"success"`
	Count   int    `json:"count"`
	Msg     string `json:"msg"`
}
