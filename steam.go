package waxpeer

import (
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"net/url"
	"strconv"
)

const (
	steamGetSteamItems      = defaultURL + "get-steam-items?"
	steamCheckTradelink     = defaultURL + "check-tradelink?"
	steamBuyOneP2P          = defaultURL + "buy-one-p2p?"
	steamBuyOneP2PName      = defaultURL + "buy-one-p2p-name?"
	steamPrices             = defaultURL + "prices?"
	steamReadyToTransferP2P = defaultURL + "ready-to-transfer-p2p?"
	steamCheckAvailability  = defaultURL + "check-availability?"
	steamGetItemsList       = defaultURL + "get-items-list?"
	steamFetchMyInventory   = defaultURL + "fetch-my-inventory?"
	steamEditItem           = defaultURL + "edit-items?"
	steamListItem           = defaultURL + "list-items-steam?"
	steamGetInventory       = defaultURL + "get-my-inventory?"
	steamSearchItemsByName  = defaultURL + "search-items-by-name?"
	steamRemoveItems        = defaultURL + "remove-items?"
	steamRemoveAllItems     = defaultURL + "remove-all?"
	steamCheckManyProjectId = defaultURL + "check-many-project-id?"
)

// appId: 730,570
func (s *Session) PricesSteam(appId uint64) ([]*steamItem, error) {
	bodyRequest := url.Values{
		"api":  {s.WaxpeerApiKey},
		"game": {strconv.FormatUint(appId, 10)},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamGetSteamItems + bodyRequest.Encode())
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body getSteamItemsResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return body.Items, nil
}

type checkTradelink struct {
	Tradelink string `json:"tradelink"`
}

// https://steamcommunity.com/tradeoffer/new/?partner=111&token=111
func (s *Session) CheckTradelink(tradelink string) (*checkTradelinkResponse, error) {
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	bodyRequestJson, err := json.Marshal(checkTradelink{Tradelink: tradelink})
	if err != nil {
		return nil, err
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamCheckTradelink + bodyRequest.Encode())
	request.Header.SetMethod("POST")
	request.Header.SetContentType("application/json")
	request.SetBody(bodyRequestJson)
	response := fasthttp.AcquireResponse()
	if err = fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body checkTradelinkResponse
	if err = json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, errors.New(body.Msg)
	}
	return &body, nil
}

type PricesConfig struct {
	Game     string // csgo, dota2
	MinPrice uint64
	MaxPrice uint64
	Search   string // search by name ex: 'hardened'.
}

// get lowest price and amount of items
func (s *Session) Prices(c PricesConfig) ([]*price, error) {
	bodyRequest := url.Values{
		"api":    {s.WaxpeerApiKey},
		"game":   {c.Game},
		"search": {c.Search},
	}
	if c.MaxPrice != 0 {
		bodyRequest.Add("max_price", strconv.FormatUint(c.MaxPrice, 10))
	}
	if c.MinPrice != 0 {
		bodyRequest.Add("min_price", strconv.FormatUint(c.MinPrice, 10))
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamPrices + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body pricesResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return body.Items, nil
}

// fetch trades that need to be sent, we recommend sending a trade once. You should be making this request at least every minute in order to be online
func AccountReadyToTransferP2P(SteamApiKey string) ([]*tradesReadyToTransferP2P, error) {
	bodyRequest := url.Values{
		"steam_api": {SteamApiKey},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamReadyToTransferP2P + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body readyToTransferP2PResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, errors.New(body.Msg)
	}
	return body.Trades, nil
}

// fetches items based on the item_id passed in query
func (s *Session) ItemAvailable(idArray *[]uint64) ([]*itemAvailable, error) {
	if len(*idArray) > 100 {
		return nil, max100Elements
	}
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	for _, id := range *idArray {
		bodyRequest.Add("item_id", strconv.FormatUint(id, 10))
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamCheckAvailability + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body itemAvailableResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return body.Items, nil
}

type PricesFilterConfig struct {
	Skip     uint64 // how many items to skip, to skip previous items
	Auto     bool   // either true or false. If you pass true, then it will return items instantly available for withdrawing
	Search   string // search items by name, ex: asiimov
	Brand    string // item category, ex: key,rifle,knife
	Order    string // ex: desc or asc
	OrderBy  string // ex: price,profit,best_deals
	Exterior string // ex: FN, MW,FT,WW,BS
	By       string // only fetch items from certain users by passing UUID from their profile page
	Limit    uint64 // how many items we would like to fetch
	Sort     string // ex: profit, desc, asc, best_deals
	MaxPrice uint64 // 1$ = 1000
	MinPrice uint64 // 1$ = 1000
	Discount uint64 // if you pass this parameter for example 10, then it will show items with discount 10% or higher
	Minified bool   // if you pass this you will receive additional info like float
	Game     string // ex: csgo, dota2
}

// fetches items based on the game you pass as a query
func (s *Session) PricesFilter(c PricesFilterConfig) ([]*priceFilter, error) {
	bodyRequest := url.Values{
		"api":      {s.WaxpeerApiKey},
		"search":   {c.Search},
		"brand":    {c.Brand},
		"order":    {c.Order},
		"order_by": {c.OrderBy},
		"exterior": {c.Exterior},
		"by":       {c.By},
		"sort":     {c.Sort},
		"game":     {c.Game},
	}
	if c.MaxPrice != 0 {
		bodyRequest.Add("max_price", strconv.FormatUint(c.MaxPrice, 10))
	}
	if c.MinPrice != 0 {
		bodyRequest.Add("min_price", strconv.FormatUint(c.MinPrice, 10))
	}
	if c.Limit != 0 {
		bodyRequest.Add("limit", strconv.FormatUint(c.Limit, 10))
	}
	if c.Discount != 0 {
		bodyRequest.Add("discount", strconv.FormatUint(c.Discount, 10))
	}
	if c.Auto {
		bodyRequest.Add("auto", "1")
	}
	if c.Minified {
		bodyRequest.Add("minified", "1")
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamGetItemsList + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body pricesFilterResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return body.Items, nil
}

// fetch my inventory
func (s *Session) AccountReloadInventory() error {
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamFetchMyInventory + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return err
	}
	var body fetchMyInventoryResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return err
	}
	if body.Success != true {
		return errors.New(body.Msg)
	}
	return nil
}

type requestItem struct {
	Items []SellItemConfig `json:"items"`
}

type SellItemConfig struct {
	ItemID int64 `json:"item_id"`
	Price  int64 `json:"price"`
}

// edit price for listed items
func (s *Session) SellEdit(c *[]SellItemConfig) (*sellEditResponse, error) {
	if len(*c) > 50 {
		return nil, max50Elements
	}
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	bodyRequestJson, err := json.Marshal(requestItem{Items: *c})
	if err != nil {
		return nil, err
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetMethod("POST")
	request.Header.SetRequestURI(steamEditItem + bodyRequest.Encode())
	request.Header.SetContentType("application/json")
	request.SetBody(bodyRequestJson)
	response := fasthttp.AcquireResponse()
	if err = fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body sellEditResponse
	if err = json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return &body, nil
}

// sell items
func (s *Session) Sell(c *[]SellItemConfig) (*sellResponse, error) {
	if len(*c) > 50 {
		return nil, max50Elements
	}
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	bodyRequestJson, err := json.Marshal(requestItem{Items: *c})
	if err != nil {
		return nil, err
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetMethod("POST")
	request.Header.SetRequestURI(steamListItem + bodyRequest.Encode())
	request.Header.SetContentType("application/json")
	request.SetBody(bodyRequestJson)
	response := fasthttp.AcquireResponse()
	if err = fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body sellResponse
	if err = json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return &body, nil
}

//get skins on sale
func (s *Session) SellOrders() ([]*sellOrders, error) {
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamListItem + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body sellOrdersResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return body.Items, nil
}

type SellItemsConfig struct {
	Skip uint64 // skip amount of items (Pagination) since default gives only an array with a length of 30
	Game uint64 // ex: 730
}

//get items that you can list for sale
func (s *Session) SellItems(c SellItemsConfig) ([]*sellItems, error) {
	bodyRequest := url.Values{
		"api":  {s.WaxpeerApiKey},
		"skip": {strconv.FormatUint(c.Skip, 10)},
		"game": {strconv.FormatUint(c.Game, 10)},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamGetInventory + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body sellItemsResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return body.Items, nil
}

// getting the cost by name
func (s *Session) PricesName(nameArray *[]string) ([]*priceName, error) {
	if len(*nameArray) > 100 {
		return nil, max100Elements
	}
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	for _, name := range *nameArray {
		bodyRequest.Add("names", name)
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamSearchItemsByName + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body pricesNameResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return body.Items, nil
}

//remove items
func (s *Session) SellRemove(idArray *[]uint64) error {
	if len(*idArray) > 1000 {
		return max1000Elements
	}
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	for _, id := range *idArray {
		bodyRequest.Add("id", strconv.FormatUint(id, 10))
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamRemoveItems + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return err
	}
	var body sellRemoveResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return err
	}
	if body.Success != true {
		return wrongApiKey
	}
	return nil
}

//remove all items
func (s *Session) SellRemoveAll() error {
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamRemoveAllItems + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return err
	}
	var body sellRemoveAllResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return err
	}
	if body.Success != true {
		return wrongApiKey
	}
	return nil
}

// account history by id
func (s *Session) AccountHistoryID(idArray *[]string) ([]*accountHistoryID, error) {
	if len(*idArray) > 100 {
		return nil, max100Elements
	}
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	for _, id := range *idArray {
		bodyRequest.Add("id", id)
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamCheckManyProjectId + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body accountHistoryIDresponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return body.Trades, nil
}

type BuyNameConfig struct {
	ProjectId string // your unique ID, max 50 symbols, it will be possible to track your trade
	Name      string // name of an item you can also pass without extension which will buy the cheapest available
	Token     string // token parameter from steam tradelink
	Price     uint64 // item price | 1$=1000
	Partner   string // partner parameter from steam tradelink
}

// buy item and send to specific tradelink
func (s *Session) BuyName(c BuyNameConfig) error {
	bodyRequest := url.Values{
		"api":        {s.WaxpeerApiKey},
		"project_id": {c.ProjectId},
		"name":       {c.Name},
		"token":      {c.Token},
		"price":      {strconv.FormatUint(c.Price, 10)},
		"partner":    {c.Partner},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamBuyOneP2PName + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return err
	}
	var body buyresponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return err
	}
	if body.Success != true {
		return errors.New(body.Msg)
	}
	return nil
}

type BuyIDConfig struct {
	ProjectId string // your unique ID, max 50 symbols, it will be possible to track your trade
	ItemId    uint64 // item id from fetching our items
	Token     string // token parameter from steam tradelink
	Price     uint64 // item price | 1$=1000
	Partner   string // partner parameter from steam tradelink
}

// buy item and send to specific tradelink
func (s *Session) BuyID(c BuyIDConfig) error {
	bodyRequest := url.Values{
		"api":        {s.WaxpeerApiKey},
		"project_id": {c.ProjectId},
		"token":      {c.Token},
		"partner":    {c.Partner},
		"item_id":    {strconv.FormatUint(c.ItemId, 10)},
		"price":      {strconv.FormatUint(c.Price, 10)},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(steamBuyOneP2P + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return err
	}
	var body buyresponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return err
	}
	if body.Success != true {
		return errors.New(body.Msg)
	}
	return nil
}
