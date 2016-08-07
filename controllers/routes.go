package controllers

import "net/http"

//Route : description d'un route http pour l'api rest
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes : un tableau de Route
type Routes []Route

var routes = Routes{
	// Route{
	// 	"Index",
	// 	http.MethodGet,
	// 	"/",
	// 	HomePage,
	// },
	Route{
		"allBank",
		http.MethodGet,
		"/bank",
		GetAllBanK,
	}, Route{
		"CreateBank",
		http.MethodPost,
		"/bank",
		CreateBank,
	}, Route{
		"UpdateBank",
		http.MethodPut,
		"/bank",
		UpdateBank,
	}, Route{
		"DeleteBankID",
		http.MethodDelete,
		"/bank/{id}",
		DeleteBankID,
	}, Route{
		"getBankByID",
		http.MethodGet,
		"/bank/{bankid}",
		SearchBankByID,
	}, Route{
		"getbankByName",
		http.MethodGet,
		"/bank/name/{name}",
		SearchBankByName,
	},
	// utilisateur
	Route{
		"allUser",
		http.MethodGet,
		"/user",
		GetAllUser,
	}, Route{
		"CreateUser",
		http.MethodPost,
		"/user",
		CreateUser,
	}, Route{
		"UpdateUser",
		http.MethodPut,
		"/user",
		UpdateUser,
	}, Route{
		"DeleteUserID",
		http.MethodDelete,
		"/user/{id}",
		DeleteUserID,
	},
	// Account
	Route{
		"allUser",
		http.MethodGet,
		"/user",
		GetAllUser,
	}, Route{
		"CreateUser",
		http.MethodPost,
		"/user",
		CreateUser,
	}, Route{
		"UpdateUser",
		http.MethodPut,
		"/user",
		UpdateUser,
	}, Route{
		"DeleteUserID",
		http.MethodDelete,
		"/user/{id}",
		DeleteUserID,
	},
	// transaction
	Route{
		"allTransaction",
		http.MethodGet,
		"/Transaction",
		GetAllTransaction,
	}, Route{
		"CreateTransaction",
		http.MethodPost,
		"/Transaction",
		CreateTransaction,
	}, Route{
		"UpdateTransaction",
		http.MethodPut,
		"/Transaction",
		UpdateTransaction,
	}, Route{
		"DeleteTransactionID",
		http.MethodDelete,
		"/Transaction/{id}",
		DeleteTransactionID,
	},
	// type de transaction
	Route{
		"allTransactionType",
		http.MethodGet,
		"/TransactionType",
		GetAllTransactionType,
	}, Route{
		"CreateTransactionType",
		http.MethodPost,
		"/TransactionType",
		CreateTransactionType,
	}, Route{
		"UpdateTransactionType",
		http.MethodPut,
		"/TransactionType",
		UpdateTransactionType,
	}, Route{
		"DeleteTransactionTypeID",
		http.MethodDelete,
		"/TransactionType/{id}",
		DeleteTransactionTypeID,
	},
}
