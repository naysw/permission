package handlers

import (
	"net/http"

	"github.com/naysw/permission/api/rest/dto"
	"github.com/naysw/permission/api/rest/res"
	"github.com/naysw/permission/internal/usecase"
)

type PolicyHandler struct {
	policyUsecase *usecase.PolicyUsecase
}

func NewPolicyHandler(pu *usecase.PolicyUsecase) *PolicyHandler {
	return &PolicyHandler{
		policyUsecase: pu,
	}
}

func (h PolicyHandler) GetList(w http.ResponseWriter, r *http.Request) {
	ps, err := h.policyUsecase.GetList(r.Context(), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res.NewOK(w, res.WithData(ps))
}

func (h PolicyHandler) GetByID(w http.ResponseWriter, r *http.Request) {
}

func (h PolicyHandler) Create(w http.ResponseWriter, r *http.Request) {
	var cp dto.CreatePolicy
	if err := parseBody(r, &cp); err != nil {
		res.NewBadRequest(w, res.WithMessage(err.Error()))
		return
	}
	if fes := cp.Validate(); len(fes) > 0 {
		res.NewUnprocessableEntity(w, res.WithErrors(fes))
		return
	}

	p, err := h.policyUsecase.CreatePolicy(
		r.Context(),
		usecase.CreatePolicyInput{
			Name:        cp.Name,
			Description: cp.Description,
			Document:    cp.Document,
		},
	)
	if err != nil {
		res.NewHttpErr(w, err)
		return
	}

	res.NewCreated(w, res.WithData(map[string]interface{}{
		"id": p,
	}))
}

func (h PolicyHandler) Update(w http.ResponseWriter, r *http.Request) {
}

func (h PolicyHandler) Delete(w http.ResponseWriter, r *http.Request) {
}

func (h PolicyHandler) Attach(w http.ResponseWriter, r *http.Request) {
	var req dto.AttachPolicy

	if err := parseBody(r, &req); err != nil {
		res.NewBadRequest(w, res.WithMessage(err.Error()))
		return
	}
	if fes := req.Validate(); len(fes) > 0 {
		res.NewUnprocessableEntity(w, res.WithErrors(fes))
		return
	}

	if err := h.policyUsecase.AttachPolicy(
		r.Context(),
		usecase.AttachPolicyInput{
			PrincipalID:   req.PrincipalID,
			PrincipalType: req.PrincipalType,
			PolicyIDs:     req.PolicyIDs,
		}); err != nil {
		res.NewHttpErr(w, err)
		return
	}

	res.NewOK(w)
}

func (h PolicyHandler) Detach(w http.ResponseWriter, r *http.Request) {
	var req dto.DetachPolicy

	if err := parseBody(r, &req); err != nil {
		res.NewBadRequest(w, res.WithMessage(err.Error()))
		return
	}
	if fes := req.Validate(); len(fes) > 0 {
		res.NewUnprocessableEntity(w, res.WithErrors(fes))
		return
	}

	if err := h.policyUsecase.DetachPolicy(
		r.Context(),
		usecase.DetachPolicyInput{
			PrincipalID:   req.PrincipalID,
			PrincipalType: req.PrincipalType,
			PolicyIDs:     req.PolicyIDs,
		}); err != nil {
		res.NewHttpErr(w, err)
		return
	}

	res.NewOK(w)
}
