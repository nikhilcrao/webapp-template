package handlers

import (
	"webapp/server/models"
)

type UserHandler struct {
	BaseHandler[models.User]
}

func (h UserHandler) QueryUserScoped() bool {
	return false
}

type AccountHandler struct {
	BaseHandler[models.Account]
}

type CategoryHandler struct {
	BaseHandler[models.Category]
}

type MerchantHandler struct {
	BaseHandler[models.Merchant]
}

type TransactionHandler struct {
	BaseHandler[models.Transaction]
}

type RuleHandler struct {
	BaseHandler[models.Rule]
}
