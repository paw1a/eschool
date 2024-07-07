package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func (h *Handler) initPaymentRoutes(api *gin.RouterGroup) {
	payment := api.Group("/payment")
	{
		payment.POST("/webhook", h.processCoursePayment)
		authenticated := payment.Group("/", h.verifyToken)
		{
			authenticated.GET("/courses/:id", h.getCoursePaymentUrl)
		}
	}
}

func (h *Handler) getCoursePaymentUrl(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	userID, err := getIdFromRequestContext(context)
	if err != nil {
		h.errorResponse(context, UnauthorizedError)
		return
	}

	url, err := h.paymentService.GetCoursePaymentUrl(
		context.Request.Context(), userID, courseID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.successResponse(context, url.String())
}

func (h *Handler) processCoursePayment(context *gin.Context) {
	key := context.PostForm("label")
	paid := context.PostForm("withdraw_amount")

	paidDigits := strings.Split(paid, ".")
	paidInt, err := strconv.ParseInt(paidDigits[0], 10, 64)
	if err != nil {
		h.errorResponse(context, BadRequestError)
		return
	}

	payload, err := h.paymentService.ProcessCoursePayment(context.Request.Context(), key, paidInt)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	err = h.courseService.AddCourseStudent(context.Request.Context(), payload.UserID, payload.CourseID)
	if err != nil {
		h.errorResponse(context, err)
		return
	}

	h.successResponse(context, "successfully paid")
}
