package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mingrammer/go-todo-rest-api-example/app/handler"
	"github.com/mingrammer/go-todo-rest-api-example/app/model"
	"github.com/mingrammer/go-todo-rest-api-example/config"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize initializes the app with predefined configuration
func (a *App) Initialize(config *config.Config) {
	// dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
	// 	config.DB.Username,
	// 	config.DB.Password,
	// 	config.DB.Host,
	// 	config.DB.Port,
	// 	config.DB.Name,
	// 	config.DB.Charset)

	//db, err := gorm.Open(config.DB.Dialect, dbURI)
	db, err := gorm.Open(config.DB.Dialect, config.DB.Username+":"+config.DB.Password+"@/"+config.DB.Name+"?charset="+config.DB.Charset+"&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", conf.MySQL.UserName+":"+conf.MySQL.Password+"@/"+conf.MySQL.DB+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// See "Important settings" section.
	//db.SetConnMaxLifetime(time.Minute * 3)
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)
	if err != nil {
		log.Fatal(err)
		log.Fatal("Could not connect database")
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the Users
	a.Get("/api/users", a.handleRequest(handler.GetAllUsers))
	a.Post("/api/user", a.handleRequest(handler.CreateUser))
	a.Get("/api/user/{id}", a.handleRequest(handler.GetUser))
	a.Put("/api/user/{id}", a.handleRequest(handler.UpdateUser))
	a.Delete("/api/user/{id}", a.handleRequest(handler.DeleteUser))

	// Routing for handling the Customers
	a.Get("/api/customers", a.handleRequest(handler.GetAllCustomers))
	a.Post("/api/customer", a.handleRequest(handler.CreateCustomer))
	a.Get("/api/customer/{id}", a.handleRequest(handler.GetCustomer))
	a.Get("/api/customers/{uid}", a.handleRequest(handler.GetCustomersByUser))
	a.Put("/api/customer/{id}", a.handleRequest(handler.UpdateCustomer))
	a.Delete("/api/customer/{id}", a.handleRequest(handler.DeleteCustomer))

	// Routing for handling the Customers
	a.Get("/api/products", a.handleRequest(handler.GetAllProducts))
	a.Post("/api/product", a.handleRequest(handler.CreateProduct))
	a.Get("/api/product/{id}", a.handleRequest(handler.GetProduct))
	a.Get("/api/products/{uid}", a.handleRequest(handler.GetProductsByUser))
	a.Put("/api/product/{id}", a.handleRequest(handler.UpdateProduct))
	a.Delete("/api/product/{id}", a.handleRequest(handler.DeleteProduct))

	a.Get("/api/orders", a.handleRequest(handler.GetAllOrders))
	a.Post("/api/order", a.handleRequest(handler.CreateOrder))
	a.Get("/api/order/{id}", a.handleRequest(handler.GetOrder))
	a.Get("/api/orders/{uid}", a.handleRequest(handler.GetOrdersByUser))
	a.Put("/api/order/{id}", a.handleRequest(handler.UpdateOrder))
	a.Delete("/api/order/{id}", a.handleRequest(handler.DeleteOrder))

	// Routing for handling the Customers
	a.Get("/api/cart_items", a.handleRequest(handler.GetAllCartItems))
	a.Post("/api/cart_item", a.handleRequest(handler.CreateCartItem))
	a.Get("/api/cart_item/{id}", a.handleRequest(handler.GetCartItem))
	a.Get("/api/cart_items/{uid}", a.handleRequest(handler.GetCartItemsByUser))
	a.Put("/api/cart_item/{id}", a.handleRequest(handler.UpdateCartItem))
	a.Delete("/api/cart_item/{id}", a.handleRequest(handler.DeleteCartItem))

	// Routing for handling the Customers
	a.Get("/api/holdcart_items", a.handleRequest(handler.GetAllHoldCartItems))
	a.Post("/api/holdcart_item", a.handleRequest(handler.CreateHoldCartItem))
	a.Get("/api/holdcart_item/{id}", a.handleRequest(handler.GetHoldCartItem))
	a.Get("/api/holdcart_items/{uid}", a.handleRequest(handler.GetHoldCartItemsByUser))
	a.Put("/api/holdcart_item/{id}", a.handleRequest(handler.UpdateHoldCartItem))
	a.Delete("/api/holdcart_item/{id}", a.handleRequest(handler.DeleteHoldCartItem))

	// Routing for handling the Category
	a.Get("/api/categories", a.handleRequest(handler.GetAllCustomers))
	a.Post("/api/category", a.handleRequest(handler.CreateCustomer))
	a.Get("/api/category/{id}", a.handleRequest(handler.GetCustomer))
	a.Get("/api/customers/{uid}", a.handleRequest(handler.GetCustomersByUser))
	a.Put("/api/category/{id}", a.handleRequest(handler.UpdateCustomer))
	a.Delete("/api/category/{id}", a.handleRequest(handler.DeleteCustomer))

	// 	// Routing for handling the tasks
	// 	a.Get("/api/projects/{title}/tasks", a.handleRequest(handler.GetAllTasks))
	// 	a.Post("/api/projects/{title}/tasks", a.handleRequest(handler.CreateTask))
	// 	a.Get("/api/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.GetTask))
	// 	a.Put("/api/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.UpdateTask))
	// 	a.Delete("/api/projects/{title}/tasks/{id:[0-9]+}", a.handleRequest(handler.DeleteTask))
	// 	a.Put("/api/projects/{title}/tasks/{id:[0-9]+}/complete", a.handleRequest(handler.CompleteTask))
	// 	a.Delete("/api/projects/{title}/tasks/{id:[0-9]+}/complete", a.handleRequest(handler.UndoTask))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.DB, w, r)
	}
}
