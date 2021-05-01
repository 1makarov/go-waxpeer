![WaxPeer](https://pbs.twimg.com/profile_banners/1146046554563895296/1562074133/1080x360)

## Installation
```sh
go get github.com/1makarov/go-waxpeer
```

##Initialization
```go
session := CreateSession(WAXPEER_API)
```

##Fetching your account info
```go
user, err := session.GetAccountInformation()
```

##Transfer of balance between users
```go
err := session.AccountTransfer(AccountTransferConfig{
    Amount: 1000,
    SteamId: 76561198322519016,
})
```

##Get history
```go
history, err := session.AccountHistory(AccountHistoryConfig{
    Skip:    0, 
    Token:   "", 
    Partner: "",
})
```

##Get history by ID
```go
history, err := session.AccountHistoryID(&[]string{
    "waxpeer",
    "red",
    "123",
})
```

##Check your account for availability
```go
err := session.AccountReloadInventory()
```

##Install steamApiKey
```go
err := session.AccountSetSteamApiKey(STEAM_API_KEY)
```

##Install tradelink
```go
err := session.AccountSetTradelink(Tradelink)
```

##Buying item with ID
```go
err := session.BuyID(BuyIDConfig{
    ProjectId: "",
    ItemId:    23786954,
    Token:     "2dl-u2kT",
    Partner:   "362253288",
    Price:     1000,
})
```

##Buying item with Name
```go
err := session.BuyName(BuyNameConfig{
    ProjectId: "",
    Name:      "AK-47 | Redline (Field-Tested)",
    Token:     "2dl-u2kT",
    Partner:   "362253288",
    Price:     10000,
})
```

##Checking tradelink
```go
err := session.CheckTradelink(Tradelink)
```

##Checking for the availability of item by ID
```go
value, err := session.ItemAvailable(&[]uint64{
    22564358567, 
    22769563455,
})
```

##Create buy order
```go
order, err := session.OrderCreate(OrderCreateConfig{
    Name: "AK-47 | Redline (Field-Tested)",
    Price: 10000,
    Amount: 5,
})
```

##Get open buy orders
```go
order, err := session.OrderOpen(OrderOpenConfig{
    Name: "",
    Skip: 0,
})
```


##Edit buy order
```go
err := session.OrderEdit(OrderEditConfig{
    ID: 1298065,
    Price: 11000,
    Amount: 4,
})
```

##Remove buy order
```go
err := session.OrderRemove(&[]uint64{
    1298065,
    1235064,
    1290062,
})
```

##Remove all buy orders
```go
err := session.OrderRemoveAll()
```

##Order History
```go
orderHistory, err := session.OrderHistory(OrderHistoryConfig{
    Skip: 0,
})
```

##Get price lists and the quantity of each item
```go
itemPrices, err := session.Prices(PricesConfig{
    Game: "csgo",
    MinPrice: 1000,
    MaxPrice: 100000,
    Search: "",
})
```

##Get prices with additional filtering
```go
itemPrices, err := session.PricesFilter(PricesFilterConfig{
    Skip: 0,
    Auto: true,
    Search: "",
    Brand: "knife",
    Order: "desc",
    OrderBy: "price",
    Exterior: "MW",
    By: "",
    Limit: 0,
    Sort: "profit",
    MaxPrice: 0,
    MinPrice: 0,
    Discount: 10,
    Minified: true, 
    Game: "csgo",
})
```

##Get the price of items by name
```go
itemPrices, err := session.PricesName(&[]string{
    "Tec-9 | Decimator (Minimal Wear)",
    "USP-S | Orion (Factory New)",
    "Special Agent Ava | FBI",
})
```

##Get prices on steam by appid
```go
steamPrices, err := session.PricesSteam(730)
```

##Sell item
```go
sellResponce, err := session.Sell(&[]SellItemConfig{
    {ItemID: 23495634332, Price: 3454},
    {ItemID: 23434127874, Price: 8766},
    {ItemID: 28454583412, Price: 91233},
})
```

##Editing the price of an item on sale
```go
editResponce, err := session.SellEdit(&[]SellItemConfig{
    {ItemID: 23495634332, Price: 3500},
    {ItemID: 23434127874, Price: 8900},
    {ItemID: 28454583412, Price: 90000},
})
```

##Get a list of items available for sale
```go
itemsResponce, err := session.SellItems(SellItemsConfig{
    Skip: 0,
    Game: 730,
})
```

##Get items that are on sale
```go
itemsResponce, err := session.SellOrders()
```

##Remove items from sale by ID
```go
err := session.SellRemove(&[]uint64{
    23495634332,
    23434127874,
    28454583412,
})
```

##Remove all items from sale
```go
err := session.SellRemoveAll()
```