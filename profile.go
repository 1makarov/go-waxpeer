package waxpeer

import (
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"net/url"
	"strconv"
)

const (
	defaultURL = "https://api.waxpeer.com/v1/"

	profileAccountInformation = defaultURL + "user?"
	profileHistory            = defaultURL + "history?"
	profileYourBuyOrders      = defaultURL + "buy-orders?"
	profileRemoveBuyOrder     = defaultURL + "remove-buy-order?"
	profileRemoveAllOrders    = defaultURL + "remove-all-orders?"
	profileBuyOrderHistory    = defaultURL + "buy-order-history?"
	profileSetSteamApiKey     = defaultURL + "set-my-steamapi?"
	profileEditBuyOrder       = defaultURL + "edit-buy-order?"
	profileCreateBuyOrder     = defaultURL + "create-buy-order?"
	profileChangeTradelink    = defaultURL + "change-tradelink?"
	profileSendBalance        = defaultURL + "transfer-money?"
)

var (
	wrongApiKey     = errors.New("your api key is not working")
	max50Elements   = errors.New("maximum number of elements 50")
	max100Elements  = errors.New("maximum number of elements 100")
	max1000Elements = errors.New("maximum number of elements 1000")
	wrongSteamId    = errors.New("this SteamID was not found")
)

type Session struct {
	WaxpeerApiKey string // apiKey Waxpeer
}

func CreateSession(WaxpeerApiKey string) *Session {
	return &Session{WaxpeerApiKey: WaxpeerApiKey}
}

//Get Account Information
func (s *Session) AccountInformation() (*accountInformation, error) {
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(profileAccountInformation + bodyRequest.Encode())
	responce := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, responce); err != nil {
		return nil, err
	}
	var body accountInformationResponse
	if err := json.Unmarshal(responce.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return body.User, nil
}

type OrderOpenConfig struct {
	Name string // Item name to filter
	Skip uint64 // we will return 100 orders, use skip to get others
}

//Get open buy orders
func (s *Session) OrderOpen(c OrderOpenConfig) ([]*orderOpen, error) {
	bodyRequest := url.Values{
		"api":  {s.WaxpeerApiKey},
		"name": {c.Name},
		"skip": {strconv.FormatUint(c.Skip, 10)},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(profileYourBuyOrders + bodyRequest.Encode())
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body orderOpenResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return body.Offers, nil
}

func (s *Session) AccountSetSteamApiKey(steamApiKey string) error {
	bodyRequest := url.Values{
		"api":       {s.WaxpeerApiKey},
		"steam_api": {steamApiKey},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(profileSetSteamApiKey + bodyRequest.Encode())
	responce := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, responce); err != nil {
		return err
	}
	var body accountSetSteamApiKeyResponse
	if err := json.Unmarshal(responce.Body(), &body); err != nil {
		return err
	}
	if body.Success != true {
		return errors.New(body.Msg)
	}
	return nil
}

func (s *Session) AccountSetTradelink(tradelink string) error {
	bodyRequest := url.Values{
		"api":       {s.WaxpeerApiKey},
		"tradelink": {tradelink},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetMethod("POST")
	request.Header.SetRequestURI(profileChangeTradelink + bodyRequest.Encode())
	responce := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, responce); err != nil {
		return err
	}
	var body accountSetTradelinkResponse
	if err := json.Unmarshal(responce.Body(), &body); err != nil {
		return err
	}
	if body.Success != true {
		return wrongApiKey
	}
	return nil
}

type AccountTransferConfig struct {
	SteamId uint64 // the id on which the translation is being made
	Amount  uint64 // 1$ = 1000
}

// sending funds between Waxpeer users
func (s *Session) AccountTransfer(c AccountTransferConfig) error {
	bodyRequest := url.Values{
		"api":      {s.WaxpeerApiKey},
		"steam_id": {strconv.FormatUint(c.SteamId, 10)},
		"amount":   {strconv.FormatUint(c.Amount, 10)},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetMethod("POST")
	request.Header.SetRequestURI(profileSendBalance + bodyRequest.Encode())
	responce := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, responce); err != nil {
		return err
	}
	var body transferResponse
	if err := json.Unmarshal(responce.Body(), &body); err != nil {
		return wrongSteamId
	}
	if body.Success != true {
		return errors.New(body.Msg)
	}
	return nil
}

// Orders
// Remove buy orders // MAX 50
func (s *Session) OrderRemove(idArray *[]uint64) error {
	if len(*idArray) > 50 {
		return max50Elements
	}
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	for _, id := range *idArray {
		bodyRequest.Add("id", strconv.FormatUint(id, 10))
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(profileRemoveBuyOrder + bodyRequest.Encode())
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return err
	}
	var body orderRemoveResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return err
	}
	if body.Success != true {
		return errors.New(body.Msg)
	}
	return nil
}

// remove all buy orders
func (s *Session) OrderRemoveAll() error {
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(profileRemoveAllOrders + bodyRequest.Encode())
	responce := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, responce); err != nil {
		return err
	}
	if len(responce.Body()) == 0 {
		return nil
	}
	var body orderRemoveAllResponce
	if err := json.Unmarshal(responce.Body(), &body); err != nil {
		return err
	}
	if body.Success != true {
		return wrongApiKey
	}
	return nil
}

type OrderHistoryConfig struct {
	Skip uint64 // by default it will return 50 trades, use skip to get others
}

func (s *Session) OrderHistory(c OrderHistoryConfig) ([]*orderHistory, error) {
	bodyRequest := url.Values{
		"api":  {s.WaxpeerApiKey},
		"skip": {strconv.FormatUint(c.Skip, 10)},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(profileBuyOrderHistory + bodyRequest.Encode())
	responce := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, responce); err != nil {
		return nil, err
	}
	var body orderHistoryResponse
	if err := json.Unmarshal(responce.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return body.History, nil
}

type OrderEditConfig struct {
	ID     uint64 `json:"id"`     // buy order id
	Price  uint64 `json:"price"`  // new price or old price | 1$ = 1000
	Amount uint64 `json:"amount"` // new amount or old amount
}

func (s *Session) OrderEdit(c OrderEditConfig) error {
	bodyRequest := url.Values{
		"api": {s.WaxpeerApiKey},
	}
	bodyRequestJson, err := json.Marshal(c)
	if err != nil {
		return err
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetMethod("POST")
	request.Header.SetRequestURI(profileEditBuyOrder + bodyRequest.Encode())
	request.Header.SetContentType("application/json")
	request.SetBody(bodyRequestJson)
	responce := fasthttp.AcquireResponse()
	if err = fasthttp.Do(request, responce); err != nil {
		return err
	}
	var body orderEditResponse
	if err = json.Unmarshal(responce.Body(), &body); err != nil {
		return err
	}
	if body.Success != true {
		return errors.New(body.Msg)
	}
	return nil
}

type OrderCreateConfig struct {
	Name   string // name of item
	Price  uint64 // max price that you want to buy item for | 1$ = 1000
	Amount uint64 // amount of items
}

func (s *Session) OrderCreate(c OrderCreateConfig) (int64, error) {
	bodyRequest := url.Values{
		"api":    {s.WaxpeerApiKey},
		"name":   {c.Name},
		"price":  {strconv.FormatUint(c.Price, 10)},
		"amount": {strconv.FormatUint(c.Amount, 10)},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetMethod("POST")
	request.Header.SetRequestURI(profileCreateBuyOrder + bodyRequest.Encode())
	responce := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, responce); err != nil {
		return 0, err
	}
	var body orderCreateResponse
	if err := json.Unmarshal(responce.Body(), &body); err != nil {
		return 0, err
	}
	if body.Success != true {
		return 0, errors.New(body.Msg)
	}
	return body.ID, nil
}

type AccountHistoryConfig struct {
	Partner string // partner link from tradelink that you used to purchase an item or steamid32
	Token   string // token used to purchase an item
	Skip    uint64 // by default it only fetched 50 items,use skip to get other trades
}

// get recent purchases
func (s *Session) AccountHistory(c AccountHistoryConfig) ([]*accountHistory, error) {
	bodyRequest := url.Values{
		"api":     {s.WaxpeerApiKey},
		"partner": {c.Partner},
		"token":   {c.Token},
		"skip":    {strconv.FormatUint(c.Skip, 10)},
	}
	request := fasthttp.AcquireRequest()
	request.Header.SetRequestURI(profileHistory + bodyRequest.Encode())
	request.Header.SetMethod("GET")
	response := fasthttp.AcquireResponse()
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	var body accountHistoryResponse
	if err := json.Unmarshal(response.Body(), &body); err != nil {
		return nil, err
	}
	if body.Success != true {
		return nil, wrongApiKey
	}
	return body.History, nil
}
