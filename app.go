package core

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

const (
	tokenType   = "bearer"
	tokenIssuer = "mafsoftware.com"
)

type AccessClaimsProvider func(token string, options TaskOptions, lang string) (*AccessClaims, *Error)

type App struct {
	Logger
	Build           string
	Router          *httprouter.Router
	claimsProvider  AccessClaimsProvider
	appSecretKey    string
	serverSecretKey string
}

func NewApp(router *httprouter.Router, logger Logger, claimsProvider AccessClaimsProvider, appSecretKey, serverSecretKey, build string) *App {
	return &App{
		Logger:          logger,
		Build:           build,
		Router:          router,
		claimsProvider:  claimsProvider,
		appSecretKey:    appSecretKey,
		serverSecretKey: serverSecretKey,
	}
}

// Methods

func (app *App) validateRequestWithOptions(r *http.Request, options TaskOptions) (UserInfo, *Error) {
	if options.RequiresSecretKey {
		err := app.checkServerSecretKey(r)
		if err != nil {
			return UserInfo{}, err
		}
	}
	lang := app.extractLanguageTag(r)
	err := app.checkAppSecretKey(r)
	if err != nil {
		return UserInfo{Language: lang}, err
	}
	if !options.RefreshingToken && options.Access == AccessOpen {
		return UserInfo{Language: lang}, nil
	}
	token, err := app.extractToken(r)
	if err != nil {
		return UserInfo{Language: lang}, err
	}
	claims, err := app.claimsProvider(token, options, lang)
	if err != nil {
		return UserInfo{Language: lang}, err
	}
	if claims.Issuer != tokenIssuer {
		return UserInfo{Language: lang}, NewError(ErrorCodeErrorBadAccessToken)
	}
	// FIXME: need to test replacing this call with a call to options.Access.HasEqualOrMoreAccessThan
	if !options.Access.HasEqualOrMoreAccess(claims.Access) {
		return UserInfo{Language: lang}, NewError(ErrorCodeErrorBadAccessToken)
	}
	userInfo := UserInfo{
		UserId:              claims.UserId,
		OrgId:               claims.OrgId,
		Access:              claims.Access,
		Language:            lang,
		ProductAccess:       claims.ProductUserAccess,
		ProductEntityAccess: claims.ProductEntityUserAccess,
	}
	return userInfo, nil
}

func (app *App) checkAppSecretKey(r *http.Request) *Error {
	appKey := r.Header.Get("AppSecretKey")
	if appKey == "" {
		return NewError(ErrorCodeAppSecretKeyNotFound)
	}
	if appKey != app.appSecretKey {
		return NewError(ErrorCodeInvalidAppSecretKey)
	}
	return nil
}

func (app *App) checkServerSecretKey(r *http.Request) *Error {
	secret := r.Header.Get("ServerSecretKey")
	if secret == "" {
		return NewError(ErrorCodeEndpointSecretKeyNotFound)
	}
	if secret != app.serverSecretKey {
		return NewError(ErrorCodeEndpointForbidden)
	}
	return nil
}

func (app *App) extractLanguageTag(r *http.Request) string {
	languages := r.Header["Accept-Language"]
	if len(languages) < 1 {
		return "en-US"
	}
	locale := r.Header["Accept-Language"][0]
	switch locale {
	case "es-419":
		return "es-419"
	default:
		return "en-US"
	}
}

func (app *App) extractToken(r *http.Request) (string, *Error) {
	authorization := r.Header.Get("Authorization")
	if authorization == "" {
		return "", NewError(ErrorCodeAccessTokenNotFound)
	}
	items := strings.Split(authorization, " ")
	if len(items) != 2 {
		return "", NewError(ErrorCodeAccessTokenMalformed)
	}
	item := strings.ToLower(items[0])
	if item != tokenType {
		return "", NewError(ErrorCodeAccessTokenMalformed)
	}
	return items[1], nil
}

func (a App) logAppServe(method, path string) {
	a.Logger.LogSystemInfo(a.Build + " : " + method + path)
}
