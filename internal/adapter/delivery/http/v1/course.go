package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/internal/adapter/delivery/http/v1/dto"
	"github.com/paw1a/eschool/internal/core/port"
)

func (h *Handler) initCourseRoutes(api *gin.RouterGroup) {
	courses := api.Group("/courses")
	{
		courses.GET("/", h.findAllCourses)
		courses.GET("/:id", h.findCourseByID)
		authenticated := courses.Group("/", h.verifyToken)
		{
			authenticated.GET("/:id/reviews", h.findCourseReviews)
			authenticated.POST("/:id/reviews", h.addCourseReview)
		}
	}
}

func (h *Handler) findAllCourses(context *gin.Context) {
	courses, err := h.courseService.FindAll(context.Request.Context())
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	courseDTOs := make([]dto.CourseDTO, len(courses))
	for i, course := range courses {
		courseDTOs[i] = dto.NewCourseDTO(course)
	}

	SuccessResponse(context, courseDTOs)
}

func (h *Handler) findCourseByID(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, PathIdParamError)
		return
	}

	course, err := h.courseService.FindByID(context.Request.Context(), courseID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	courseDTO := dto.NewCourseDTO(course)
	SuccessResponse(context, courseDTO)
}

func (h *Handler) addCourseReview(context *gin.Context) {
	userID, err := getIdFromRequestContext(context)
	if err != nil {
		ErrorResponse(context, UnauthorizedError)
		return
	}

	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, PathIdParamError)
		return
	}

	var createReviewDTO dto.CreateReviewDTO
	err = context.BindJSON(&createReviewDTO)
	if err != nil {
		ErrorResponse(context, UnmarshalError)
		return
	}

	review, err := h.reviewService.CreateCourseReview(context, courseID, userID,
		port.CreateReviewParam{Text: createReviewDTO.Text})

	reviewDTO := dto.NewReviewDTO(review)
	SuccessResponse(context, reviewDTO)
}

func (h *Handler) findCourseReviews(context *gin.Context) {
	courseID, err := getIdFromPath(context, "id")
	if err != nil {
		ErrorResponse(context, PathIdParamError)
		return
	}

	reviews, err := h.reviewService.FindCourseReviews(context.Request.Context(), courseID)
	if err != nil {
		ErrorResponse(context, err)
		return
	}

	reviewDTOs := make([]dto.ReviewDTO, len(reviews))
	for i, review := range reviews {
		reviewDTOs[i] = dto.NewReviewDTO(review)
	}

	SuccessResponse(context, reviewDTOs)
}
