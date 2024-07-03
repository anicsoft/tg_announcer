package api

import (
	"github.com/gin-gonic/gin"
)

// UploadImage godoc
//
//	@Summary		Upload an image
//	@Description	Uploads an image for an announcement to S3 and updates the entity's record with the S3 URL.
//	@Tags			announcements
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id				path		string				true	"announcements ID"
//	@Param			Authorization	header		string				true	"Authorization token"
//	@Param			image			formData	file				true	"Logo image file"
//	@Success		200				{object}	model.S3Response	"Successfully uploaded"
//	@Failure		400				{object}	HttpError			"Bad request"
//	@Failure		500				{object}	HttpError			"Internal server error"
//	@Router			/announcements/{id}/image [post]
func (a *BaseApi) UploadImage(ctx *gin.Context) {
	id := ctx.Param("id")
	header, err := ctx.FormFile("image")
	if err != nil {
		StatusBadRequest(ctx, err)
		return
	}

	s3URL, err := a.imageService.UploadAnnouncPictures(ctx, id, header)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, s3URL)
}

// UploadLogo godoc
//
//	@Summary		Upload a logo image
//	@Description	Uploads a logo image for a company to S3 and updates the entity's record with the S3 URL.
//	@Tags			companies
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			id				path		string				true	"company ID"
//	@Param			Authorization	header		string				true	"Authorization token"
//	@Param			image			formData	file				true	"Logo image file"
//	@Success		200				{object}	model.S3Response	"Successfully uploaded"
//	@Failure		400				{object}	HttpError			"Bad request"
//	@Failure		401				{object}	HttpError			"Unauthorized"
//	@Failure		403				{object}	HttpError			"Forbidden"
//	@Failure		404				{object}	HttpError			"Entity not found"
//	@Failure		500				{object}	HttpError			"Internal server error"
//	@Router			/companies/{id}/logo [post]
func (a *BaseApi) UploadLogo(ctx *gin.Context) {
	id := ctx.Param("id")
	header, err := ctx.FormFile("image")
	if err != nil {
		StatusBadRequest(ctx, err)
		return
	}

	s3URL, err := a.imageService.UploadLogo(ctx, id, header)
	if err != nil {
		StatusInternalServerError(ctx, err)
		return
	}

	StatusOK(ctx, s3URL)
}
