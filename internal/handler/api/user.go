package api

import (
	"net/http"

	"github.com/backend-magang/eniqilo-store/models"
	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/backend-magang/eniqilo-store/utils/constant"
	"github.com/backend-magang/eniqilo-store/utils/helper"
	"github.com/backend-magang/eniqilo-store/utils/pkg"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
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

func (h *handler) LoginStaff(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := entity.LoginStaffRequest{}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.LoginStaff(ctx, request)
	return helper.WriteResponse(c, resp)
}

func (h *handler) RegisterCustomer(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := entity.RegisterCustomerRequest{}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.RegisterCustomer(ctx, request)
	return helper.WriteResponse(c, resp)
}

func (h *handler) GetListCustomer(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := entity.GetListUserRequest{
		PhoneNumber: c.QueryParam("phoneNumber"),
		Name:        c.QueryParam("name"),
		Limit:       c.QueryParam("limit"),
		Offset:      c.QueryParam("offset"),
		Role:        constant.ROLE_CUSTOMER,
	}

	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	if cast.ToInt(request.Limit) == 0 {
		request.Limit = "5"
	}

	if cast.ToInt(request.Offset) == 0 {
		request.Offset = "0"
	}

	resp := h.usecase.GetListCustomer(ctx, request)
	return helper.WriteResponse(c, resp)
}
