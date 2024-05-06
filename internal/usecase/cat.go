package usecase

// func (u *usecase) GetListCat(ctx context.Context, req entity.GetListCatRequest) models.StandardResponseReq {
// 	var (
// 		listCats = []entity.GetListCatResponse{}
// 	)

// 	if err := builFilterAgeRequest(&req); err != nil {
// 		return models.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.FAILED, Error: err}
// 	}

// 	cats, err := u.repository.GetListCat(ctx, req)
// 	if err != nil {
// 		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED_GET_CATS, Error: err}
// 	}

// 	for _, cat := range cats {
// 		listCats = append(listCats, buildResponseCat(cat))
// 	}

// 	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS, Data: listCats}
// }

// func (u *usecase) CreateCat(ctx context.Context, req entity.CreateCatRequest) models.StandardResponseReq {
// 	now := time.Now()

// 	newCat := entity.Cat{
// 		UserID:           req.UserID,
// 		Name:             req.Name,
// 		Race:             req.Race,
// 		Sex:              req.Sex,
// 		Age:              req.AgeInMonth,
// 		Description:      req.Description,
// 		Images:           req.ImageUrls,
// 		IsAlreadyMatched: false,
// 		CreatedAt:        now,
// 		UpdatedAt:        now,
// 	}

// 	cat, err := u.repository.InsertCat(ctx, newCat)
// 	if err != nil {
// 		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
// 	}

// 	resp := entity.CreateCatResponse{
// 		ID:        cast.ToString(cat.ID),
// 		CreatedAt: cat.CreatedAt,
// 	}

// 	return models.StandardResponseReq{Code: http.StatusCreated, Message: constant.SUCCESS_ADD_CAT, Data: resp, Error: nil}
// }

// func (u *usecase) UpdateCat(ctx context.Context, req entity.UpdateCatRequest) models.StandardResponseReq {
// 	var (
// 		now    = time.Now()
// 		userId = req.UserID
// 		catId  = cast.ToInt(req.ID)
// 	)

// 	// find user's cat by id
// 	cat, err := u.repository.FindUserCatByID(ctx, userId, catId)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED_GET_USER_CAT, Error: err}
// 		}
// 		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
// 	}

// 	// if sex changed
// 	if req.Sex != cat.Sex {
// 		matchCats, err := u.repository.FindRequestedMatch(ctx, catId)
// 		if err != nil && err != sql.ErrNoRows {
// 			return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
// 		}

// 		// if there is a match request
// 		if len(matchCats) > 0 {
// 			return models.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.HAS_REQUESTED_MATCH, Error: err}
// 		}
// 	}

// 	updatedCat := entity.Cat{
// 		UserID:      userId,
// 		ID:          catId,
// 		Name:        req.Name,
// 		Race:        req.Race,
// 		Sex:         req.Sex,
// 		Age:         req.AgeInMonth,
// 		Description: req.Description,
// 		Images:      req.ImageUrls,
// 		UpdatedAt:   now,
// 	}

// 	cat, err = u.repository.UpdateCat(ctx, updatedCat)
// 	if err != nil {
// 		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
// 	}

// 	resp := entity.UpdateCatResponse{
// 		ID:          cast.ToString(cat.ID),
// 		Name:        cat.Name,
// 		Race:        cat.Race,
// 		Sex:         cat.Sex,
// 		AgeInMonth:  cat.Age,
// 		Description: cat.Description,
// 		ImageUrls:   cat.Images,
// 		CreatedAt:   cat.CreatedAt,
// 		UpdatedAt:   cat.UpdatedAt,
// 	}

// 	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_UPDATE_CAT, Data: resp}
// }

// func (u *usecase) DeleteCat(ctx context.Context, req entity.DeleteCatRequest) models.StandardResponseReq {
// 	var (
// 		now = time.Now()
// 	)

// 	// Find cat by id
// 	cat, err := u.repository.FindCatByID(ctx, cast.ToInt(req.ID))
// 	if err == sql.ErrNoRows {
// 		return models.StandardResponseReq{Code: http.StatusNotFound, Message: constant.FAILED_CAT_NOT_FOUND, Error: err}
// 	}

// 	// If cat isn't belong to user
// 	if req.UserID != cat.UserID {
// 		return models.StandardResponseReq{Code: http.StatusBadRequest, Message: constant.FAILED, Error: errors.New(constant.FAILED_CAT_USER_UNAUTHORIZED)}
// 	}

// 	cat.UpdatedAt = now
// 	cat.DeletedAt = sql.NullTime{Time: now, Valid: true}

// 	_, err = u.repository.DeleteCat(ctx, cat)
// 	if err != nil {
// 		return models.StandardResponseReq{Code: http.StatusInternalServerError, Message: constant.FAILED, Error: err}
// 	}

// 	return models.StandardResponseReq{Code: http.StatusOK, Message: constant.SUCCESS_DELETE_CAT}
// }
