package model

import (
	"github.com/jinzhu/gorm"
	
)

//User table
type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"isAdmin"`
}

//Customer table
type Customer struct {
	CustomerID        int    `gorm:"primary_key" json:"customerId"`
	CustomerFirstName string `json:"customerFirstName"`
	CustomerLastName  string `json:"customerLastName"`
	Email             string `json:"email"`
	ContactNumber     string `json:"contactNumber"`
	AddressLine       string `json:"addressLine"`
	City              string `json:"city"`
	PostalCode        string `json:"postalCode"`
	State             string `json:"state"`
	Country           string `json:"country"`
	UID               uint   `json:"uid"`
}

//Category Table
type Category struct {
	CID                 int    `gorm:"primary_key" json:"categoryId"`
	CategoryName        string `json:"categoryName"`
	Active              bool   `json:"active"`
	Level               int    `json:"level"`
	ParentID            int    `json:"parentId"`
	Icon                int    `json:"icon"`
	Path                string `json:"path"`
	IncludeInDrawerMenu bool   `json:"includeInDrawerMenu"`
	UID                 int    `json:"uid"`
}

//OptionValue : Option values table
type OptionValue struct {
	OptionValueID    int    `gorm:"primary_key" json:"optionValueId"`
	OptionID         int    `json:"optionId"`
	OptionValueName  string `json:"optionValueName"`
	OptionValuePrice string `json:"optionValuePrice"`
	OptionOptionID   int    //forign key
}

//Option Table
type Option struct {
	OptionID     int           `gorm:"primary_key" json:"optionId"`
	OptionName   string        `json:"optionName"`
	Type         string        `json:"type"`
	OptionValues []OptionValue `json:"optionValues"`
	SortOrder    int           `json:"sortOrder"`
	ProductPID   int           //foriegn key
}

//Tax :tax table
type Tax struct {
	TaxID      int    `json:"taxId"`
	TaxName    string `json:"taxName"`
	Enabled    bool   `json:"enabled"`
	Type       string `json:"type"`
	TaxRate    string `json:"taxRate"`
	ProductPID int    //forignKey
}

//Product table
type Product struct {
	PID                   int                    `gorm:"primary_key" json:"pid"`
	ProductName           string                 `json:"productName"`
	ProductShortDes       string                 `json:"productShortDes"`
	Sku                   string                 `json:"sku"`
	IsEnabled             bool                   `json:"isEnabled"`
	Price                 float32                `json:"price"`
	FormattedPrice        string                 `json:"formattedPrice"`
	SpecialPrice          float32                `json:"specialPrice"`
	FormattedSpecialPrice string                 `json:"formattedSpecialPrice"`
	IsTaxableGoodsApplied bool                   `json:"isTaxableGoodsApplied"`
	ProductTax            Tax                    `json:"productTax"`
	TrackInventory        bool                   `json:"trackInventory"`
	Quantity              float32                `json:"quantity"`
	Stock                 bool                   `json:"stock"`
	Discount              float32                `json:"discount"`
	FormattedDiscount     string                 `json:"formattedDiscount"`
	Image                 string                 `json:"image"`
	Weight                string                 `json:"weight"`
	BarCode               int64                  `json:"barCode"`
	ProductCategories     []ProductCategoryModel `json:"productCategories"`
	Options               []Option               `json:"options"`
}

//TotalModel data : part of CartModel table
type TotalModel struct {
	SubTotal               string  `json:"subTotal"`
	FormatedSubTotal       string  `json:"formatedSubTotal"`
	Qty                    string  `json:"qty"`
	Tax                    string  `json:"tax"`
	FormatedTax            string  `json:"formatedTax"`
	Discount               string  `json:"discount"`
	TotalDiscountByProduct float32 `json:"totalDiscountByProduct"`
	FormatedDiscount       string  `json:"formatedDiscount"`
	GrandTotal             string  `json:"grandTotal"`
	FormatedGrandTotal     string  `json:"formatedGrandTotal"`
	RoundTotal             string  `json:"roundTotal"`
	FormatedRoundTotal     string  `json:"formatedRoundTotal"`
	DisplayError           bool    `json:"displayError"`
}

//CartModel : part of HoldCart table
type CartModel struct {
	Products           []Product  `json:"products"`
	Totals             TotalModel `json:"totals"`
	Customer           Customer   `json:"customer"`
	HoldCartHoldCartID int        //forignkey
	OrderOrderID       int        //forignkey
	UID                int        `json:"uid"`
}

//CashModel : part of Order table
type CashModel struct {
	Total                  string `json:"total"`
	FormatedTotal          string `json:"formatedTotal"`
	CollectedCash          string `json:"collectedCash"`
	FormattedCollectedCash string `json:"formattedCollectedCash"`
	Note                   string `json:"note"`
	FormattedChangeDue     string `json:"formattedChangeDue"`
	ChangeDue              string `json:"changeDue"`
	ChangeDueVisibility    bool   `json:"changeDueVisibility"`
	DisplayError           bool   `json:"displayError"`
	OrderOrderID           int    //forignkey
}

//HoldCart table
type HoldCart struct {
	HoldCartID int       `gorm:"primary_key" json:"holdCartId"`
	Time       string    `json:"time"`
	Date       string    `json:"date"`
	CartData   CartModel `json:"cartData"`
	Qty        string    `json:"qty"`
	IsSynced   string    `json:"isSynced"`
	UID        int       `json:"uid"`
}

//Order table
type Order struct {
	OrderID         int       `gorm:"primary_key" json:"orderId"`
	Time            string    `json:"time"`
	Date            string    `json:"date"`
	CartData        CartModel `json:"cartData"`
	Qty             string    `json:"qty"`
	CashData        CashModel `json:"cashData"`
	IsSynced        string    `json:"isSynced"`
	IsReturned      bool      `json:"isReturned"`
	RefundedOrderID string    `json:"refundedOrderId"`
	UID             int       `json:"uid"`
}

//ProductCategoryModel table : less values from category table
type ProductCategoryModel struct {
	CID        string `json:"cid"`
	Name       string `json:"name"`
	ProductPID int    //forignkey
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(
		&User{},
		&TotalModel{},
		&Tax{},
		&ProductCategoryModel{},
		&Product{},
		&Order{},
		&OptionValue{},
		&Option{},
		&HoldCart{},
		&Category{},
		&CashModel{},
		&CartModel{},
		&Customer{})
	return db
}
