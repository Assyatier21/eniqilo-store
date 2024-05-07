package api

import (
	"net/http"

	"github.com/backend-magang/eniqilo-store/models"
	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/backend-magang/eniqilo-store/utils/helper"
	"github.com/backend-magang/eniqilo-store/utils/pkg"
	"github.com/labstack/echo/v4"
)

func (h *handler) RegisterStaff(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := entity.RegisterStaffRequest{}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.RegisterStaff(ctx, request)
	return helper.WriteResponse(c, resp)
}

// func (h *handler) LoginUser(c echo.Context) (err error) {
// 	ctx, cancel := helper.GetContext()
// 	defer cancel()

// 	request := entity.LoginUserRequest{}
// 	err = pkg.BindValidate(c, &request)
// 	if err != nil {
// 		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
// 	}

// 	resp := h.usecase.LoginUser(ctx, request)
// 	return helper.WriteResponse(c, resp)
// }
