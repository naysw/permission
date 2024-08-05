package handlers

import (
	"net/http"

	"github.com/naysw/permission/internal/usecase"
)

type RoleHandler struct {
	roleUsecase *usecase.RoleUsecase
}

func NewRoleHandler(ru *usecase.RoleUsecase) *RoleHandler {
	return &RoleHandler{
		roleUsecase: ru,
	}
}

func (h RoleHandler) GetList(w http.ResponseWriter, r *http.Request) {
}

func (h RoleHandler) GetByID(w http.ResponseWriter, r *http.Request) {
}

func (h RoleHandler) Create(w http.ResponseWriter, r *http.Request) {
}

func (h RoleHandler) Update(w http.ResponseWriter, r *http.Request) {
}

func (h RoleHandler) Delete(w http.ResponseWriter, r *http.Request) {
}
