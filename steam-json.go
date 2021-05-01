package waxpeer

import "time"

type getSteamItemsResponse struct {
	Success bool         `json:"success"`
	Items   []*steamItem `json:"items"`
}

type steamItem struct {
	Name       string      `json:"name"`
	Average    int64       `json:"average"`
	GameID     int64       `json:"game_id"`
	Type       interface{} `json:"type"`
	Collection interface{} `json:"collection"`
	RuName     interface{} `json:"ru_name"`
}

type checkTradelinkResponse struct {
	Success   bool        `json:"success"`
	Msg       string      `json:"msg"`
	Info      interface{} `json:"info"`
	Link      string      `json:"link"`
	Token     string      `json:"token"`
	Steamid32 string      `json:"steamid32"`
	Steamid64 string      `json:"steamid64"`
}

type buyresponce struct {
	Success  bool   `json:"success"`
	ID       int64  `json:"id"`
	Price    int64  `json:"price"`
	Msg      string `json:"msg"`
	ErrorMsg string `json:"error_msg"`
}

type pricesResponse struct {
	Success bool     `json:"success"`
	Items   []*price `json:"items"`
}

type price struct {
	Name  string `json:"name"`
	Min   int64  `json:"min"`
	Avg   int64  `json:"avg"`
	Max   int64  `json:"max"`
	Count int64  `json:"count"`
}

type accountHistoryResponse struct {
	Success bool              `json:"success"`
	History []*accountHistory `json:"history"`
}

type accountHistory struct {
	TradeID   interface{} `json:"trade_id"`
	Token     string      `json:"token"`
	Partner   int64       `json:"partner"`
	Created   time.Time   `json:"created"`
	SendUntil time.Time   `json:"send_until"`
	Reason    interface{} `json:"reason"`
	ID        int64       `json:"id"`
	ItemID    string      `json:"item_id"`
	Image     string      `json:"image"`
	Price     int64       `json:"price"`
	Name      string      `json:"name"`
	Status    int64       `json:"status"`
}

type readyToTransferP2PResponse struct {
	Success bool                        `json:"success"`
	Trades  []*tradesReadyToTransferP2P `json:"trades"`
	Msg     string                      `json:"msg"`
}

type tradesReadyToTransferP2P struct {
	ID           string `json:"id"`
	CostumID     string `json:"costum_id"`
	TradeID      int64  `json:"trade_id"`
	Status       string `json:"status"`
	TradeMessage string `json:"trade_message"`
	Tradelink    string `json:"tradelink"`
	Done         bool   `json:"done"`
	ForSteamid32 string `json:"for_steamid32"`
	ForSteamid64 string `json:"for_steamid64"`
	Created      string `json:"created"`
	SendUntil    string `json:"send_until"`
	Items        []*struct {
		ID         int64  `json:"id"`
		ItemID     string `json:"item_id"`
		GiveAmount int64  `json:"give_amount"`
		Image      string `json:"image"`
		Price      int64  `json:"price"`
		Game       string `json:"game"`
		Name       string `json:"name"`
		Status     int64  `json:"status"`
	} `json:"items"`
}

type itemAvailableResponse struct {
	Success bool             `json:"success"`
	Items   []*itemAvailable `json:"items"`
}

type itemAvailable struct {
	ItemID  string `json:"item_id"`
	Selling bool   `json:"selling"`
	Price   int64  `json:"price"`
	Name    string `json:"name"`
	Image   string `json:"image"`
}

type pricesFilterResponse struct {
	Success bool           `json:"success"`
	Items   []*priceFilter `json:"items"`
}

type priceFilter struct {
	ItemID     string  `json:"item_id"`
	Brand      string  `json:"brand"`
	Image      string  `json:"image"`
	Price      int64   `json:"price"`
	Name       string  `json:"name"`
	Float      float64 `json:"float"`
	BestDeals  int64   `json:"best_deals"`
	Discount   int64   `json:"discount"`
	SteamPrice int64   `json:"steam_price"`
	Type       string  `json:"type"`
}

type fetchMyInventoryResponse struct {
	Success             bool   `json:"success"`
	TotalInventoryCount int64  `json:"total_inventory_count"`
	Msg                 string `json:"msg"`
}

type sellEditResponse struct {
	Success bool               `json:"success"`
	Updated []*sellEditItem    `json:"updated"`
	Failed  []*sellFailedItem  `json:"failed"`
	Removed []*sellRemovedItem `json:"removed"`
}

type sellEditItem struct {
	ItemID string `json:"item_id"`
	Price  string `json:"price"`
}

type sellFailedItem struct {
	ItemID int    `json:"item_id"`
	Msg    string `json:"msg"`
}

type sellRemovedItem struct {
	Price  string `json:"price"`
	ItemID int    `json:"item_id"`
}

type sellResponse struct {
	Success bool          `json:"success"`
	Msg     string        `json:"msg"`
	Listed  []*listedItem `json:"listed"`
	Failed  []*failedItem `json:"failed"`
}

type listedItem struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	ItemID   int64  `json:"item_id"`
	Position int    `json:"position"`
}

type failedItem struct {
	Name   string `json:"name"`
	Price  int    `json:"price"`
	ItemID int64  `json:"item_id"`
	Msg    string `json:"msg"`
}

type sellOrdersResponse struct {
	Success bool          `json:"success"`
	Items   []*sellOrders `json:"items"`
}

type sellOrders struct {
	ItemID     int64     `json:"item_id,int64"`
	Price      int       `json:"price"`
	Date       time.Time `json:"date"`
	Position   int       `json:"position"`
	Name       string    `json:"name"`
	SteamPrice struct {
		Average int    `json:"average"`
		Current int    `json:"current"`
		Img     string `json:"img"`
	} `json:"steam_price"`
}

type sellItemsResponse struct {
	Success bool         `json:"success"`
	Items   []*sellItems `json:"items"`
	Count   int          `json:"count"`
}

type sellItems struct {
	ItemID     int64  `json:"item_id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	SteamPrice struct {
		Average      int64  `json:"average"`
		Current      int64  `json:"current"`
		Img          string `json:"img"`
		LowestPrice  int64  `json:"lowest_price"`
		HighestOffer int64  `json:"highest_offer"`
	} `json:"steam_price"`
}

type pricesNameResponse struct {
	Success bool         `json:"success"`
	Items   []*priceName `json:"items"`
}

type priceName struct {
	Name   string `json:"name"`
	Price  int64  `json:"price"`
	Image  string `json:"image"`
	ItemID string `json:"item_id"`
}

type sellRemoveResponse struct {
	Success bool     `json:"success"`
	Count   int      `json:"count"`
	Removed *[]int64 `json:"removed"`
}

type sellRemoveAllResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Count   int    `json:"count"`
}

type accountHistoryIDresponce struct {
	Success bool                `json:"success"`
	Trades  []*accountHistoryID `json:"trades"`
}

type accountHistoryID struct {
	ID           uint64      `json:"id,uint64"`
	Price        uint64      `json:"price,uint64"`
	Name         string      `json:"name"`
	Status       int         `json:"status"`
	ProjectID    string      `json:"project_id"`
	CustomID     string      `json:"custom_id"`
	TradeID      interface{} `json:"trade_id"`
	Done         bool        `json:"done"`
	ForSteamid64 string      `json:"for_steamid64"`
	Reason       interface{} `json:"reason"`
	SendUntil    int         `json:"send_until"`
	LastUpdated  int         `json:"last_updated"`
	Counter      int         `json:"counter"`
	Msg          string      `json:"msg"`
}
