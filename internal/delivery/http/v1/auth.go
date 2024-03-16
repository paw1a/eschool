package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/paw1a/eschool/pkg/auth"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func extractAuthToken(context *gin.Context) (string, error) {
	authHeader := context.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	log.Printf("token = %s", headerParts[1])

	return headerParts[1], nil
}

func (h *Handler) refreshToken(context *gin.Context) {
	var input auth.RefreshInput

	err := context.BindJSON(&input)
	if err != nil {
		badRequestResponse(context, "can't parse request body", err)
		return
	}

	refreshCookie, err := context.Cookie("refreshToken")
	if err != nil {
		badRequestResponse(context, "refresh cookie not found", err)
		return
	}

	input.RefreshToken = refreshCookie

	authDetails, err := h.tokenProvider.Refresh(auth.RefreshInput{
		RefreshToken: input.RefreshToken,
		Fingerprint:  input.Fingerprint,
	})

	if err != nil {
		unauthorizedResponse(context, err.Error())
		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("refreshToken", authDetails.RefreshToken,
		int(h.config.JWT.RefreshTokenTime), "/", h.config.Server.Host, false, false)

	successResponse(context, authDetails.AccessToken)
}

func (h *Handler) verifyToken(context *gin.Context, idName string) {
	tokenString, err := extractAuthToken(context)
	if err != nil {
		unauthorizedResponse(context, err.Error())
		return
	}

	tokenClaims, err := h.tokenProvider.VerifyToken(tokenString)
	if err != nil {
		unauthorizedResponse(context, err.Error())
		return
	}

	id, ok := tokenClaims[idName]
	if !ok {
		errorResponse(context, http.StatusForbidden, "this endpoint is forbidden")
		return
	}

	context.Set(idName, id)
}

func (h *Handler) extractIdFromAuthHeader(context *gin.Context, idName string) (int64, error) {
	tokenString, err := extractAuthToken(context)
	if err != nil {
		return 0, err
	}

	tokenClaims, err := h.tokenProvider.VerifyToken(tokenString)
	if err != nil {
		return 0, err
	}

	id, ok := tokenClaims[idName]
	if !ok {
		return 0, fmt.Errorf("failed to extract %s from auth header", idName)
	}

	return id.(int64), nil
}
