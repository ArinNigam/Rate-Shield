package api

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/x-sushant-x/RateShield/models"
	"github.com/x-sushant-x/RateShield/service"
	"github.com/x-sushant-x/RateShield/utils"
)

type RulesAPIHandler struct {
	rulesSvc service.RulesService
}

func NewRulesAPIHandler(svc service.RulesService) RulesAPIHandler {
	return RulesAPIHandler{
		rulesSvc: svc,
	}
}

func (h RulesAPIHandler) ListAllRules(w http.ResponseWriter, r *http.Request) {
	rules, err := h.rulesSvc.GetAllRules()
	if err != nil {
		utils.InternalError(w)
	}
	utils.SuccessResponse(rules, w)
}

func (h RulesAPIHandler) CreateOrUpdateRule(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		updateReq, err := utils.ParseAPIBody[models.Rule](r)
		if err != nil {
			utils.BadRequestError(w)
		}

		log.Info().Msgf("Request Body: %v", updateReq)

		err = h.rulesSvc.CreateOrUpdateRule(updateReq)
		if err != nil {
			utils.InternalError(w)
		}

		utils.SuccessResponse("Rule Created Successfully", w)
	} else {
		utils.MethodNotAllowedError(w)
	}
}

func (h RulesAPIHandler) DeleteRule(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		deleteReq, err := utils.ParseAPIBody[models.DeleteRuleDTO](r)
		if err != nil {
			utils.BadRequestError(w)
		}

		err = h.rulesSvc.DeleteRule(deleteReq.RuleKey)
		if err != nil {
			utils.InternalError(w)
		}

		utils.SuccessResponse("Rule Deleted Successfully", w)
	} else {
		utils.MethodNotAllowedError(w)
	}
}
