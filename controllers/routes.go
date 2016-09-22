package controllers

import "net/http"

//Route : description d'une route http pour l'api rest
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes : un tableau de Route
type Routes []Route

// getRouteWithAuth : retourne la liste des routes de l'application
// TODO : implementer l'utilisation du verbe HTTP options pour chaque ressource histoire de founir l'info
func getRouteWithAuth() Routes {
	return Routes{
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
			"/bank/{id}",
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
			"SearchUserByID",
			http.MethodGet,
			"/user/{id}",
			SearchUserByID,
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
			"allAccount",
			http.MethodGet,
			"/account",
			GetAllAccount,
		}, Route{
			"SearchAccountByID",
			http.MethodGet,
			"/account/{id}",
			SearchAccountByID,
		}, Route{
			"CreateAccount",
			http.MethodPost,
			"/account",
			CreateAccount,
		}, Route{
			"UpdateAccount",
			http.MethodPut,
			"/account",
			UpdateAccount,
		}, Route{
			"DeleteAccountID",
			http.MethodDelete,
			"/account/{id}",
			DeleteAccountID,
		}, Route{
			"DeleteAccountByNumber",
			http.MethodDelete,
			"/account/number/{number}",
			DeleteAccountByNumber,
		},
		// transaction
		Route{
			"allTransaction",
			http.MethodGet,
			"/transaction",
			GetAllTransaction,
		}, Route{
			"getTransactionByID",
			http.MethodGet,
			"/transaction/{id}",
			SearchTransactionByID,
		}, Route{
			"CreateTransaction",
			http.MethodPost,
			"/transaction",
			CreateTransaction,
		}, Route{
			"UpdateTransaction",
			http.MethodPut,
			"/transaction/{id}",
			UpdateTransaction,
		}, Route{
			"DeleteTransactionID",
			http.MethodDelete,
			"/transaction/{id}",
			DeleteTransactionID,
		},
		// type de transaction
		Route{
			"allTransactionType",
			http.MethodGet,
			"/transactionType",
			GetAllTransactionType,
		}, Route{
			"getTransactionTypeByID",
			http.MethodGet,
			"/transactionType/{id}",
			SearchTransactionTypeByID,
		}, Route{
			"CreateTransactionType",
			http.MethodPost,
			"/transactionType",
			CreateTransactionType,
		}, Route{
			"UpdateTransactionType",
			http.MethodPut,
			"/transactionType",
			UpdateTransactionType,
		}, Route{
			"DeleteTransactionTypeID",
			http.MethodDelete,
			"/transactionType/{id}",
			DeleteTransactionTypeID,
		},
	}
}

// getRouteWithoutAuth Liste des routes qui ne demandes pas d'authentification, ou qui la fournisse ;)
func getRouteWithoutAuth() Routes {
	return Routes{
		Route{
			"Index",
			http.MethodGet,
			"/",
			HomePage,
		}, Route{
			"Version",
			http.MethodGet,
			"/version",
			getVersion,
		}, Route{
			"UserAuthenticate",
			http.MethodPost,
			"/user/authenticate",
			UserAuthenticate,
		}, Route{
			"Logout",
			http.MethodDelete,
			"/user/logout",
			UserLogout,
		},
	}
}
