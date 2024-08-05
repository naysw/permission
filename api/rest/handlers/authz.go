package handlers

import (
	"fmt"
	"net/http"

	"github.com/naysw/permission/api/rest/dto"
	"github.com/naysw/permission/api/rest/res"
	"github.com/naysw/permission/internal/usecase"
)

type AuthzHandler struct {
	roleUsecase   *usecase.RoleUsecase
	policyUsecase *usecase.PolicyUsecase
}

func NewAuthzHandler(ru *usecase.RoleUsecase, pu *usecase.PolicyUsecase) AuthzHandler {
	return AuthzHandler{
		roleUsecase:   ru,
		policyUsecase: pu,
	}
}

func (h AuthzHandler) IsAuthorized(w http.ResponseWriter, r *http.Request) {
	var req dto.AuthorizedDto
	if err := parseBody(r, &req); err != nil {
		res.NewBadRequest(w, res.WithMessage(err.Error()))
		return
	}
	if fes := req.Validate(); len(fes) > 0 {
		res.NewUnprocessableEntity(w, res.WithErrors(fes))
		return
	}

	result, err := h.policyUsecase.Authorized(
		r.Context(),
		req.Entities,
		req.Request,
	)
	if err != nil {
		fmt.Println(err)
		res.NewHttpErr(w, err)
		return
	}

	res.NewOK(w, res.WithData(map[string]interface{}{
		"allow": result,
	}))
}
