package core

import (
	"time"

	"golang.org/x/text/language"
)

type Organization struct {
	Id       string
	OwnerId  string
	Name     string
	Domain   string
	Created  time.Time
	Status   AccountStatus
	Verified bool
}

type User struct {
	Id         string
	OrgId      string
	FirstName  string
	MiddleName string
	LastName   string
	Username   string
	Password   string
	Email      string
	Access     Access
	Enabled    bool
	Verified   bool
}

type Product struct {
	Id          string
	Name        string
	Description string
	Uri         string
	Created     time.Time
	LastUpdated time.Time
	Stage       ProductStage
}

type Plan struct {
	Id           string
	ProductId    string
	Name         string
	Description  string
	PayType      PayType
	PayAmount    float64
	PayFrequency PayFrequency
	Created      time.Time
	LastUpdated  time.Time
	Stage        ProductStage
}

type RefreshToken struct {
	Id      string
	UserId  string
	Issued  time.Time
	Expires int64
}

type AuthTokens struct {
	Type         string
	IdToken      string
	AccessToken  string
	RefreshToken RefreshToken
	ExpiresAt    int64
}

type CodeType int

const (
	CodeTypeRegistration    CodeType = 1
	CodeTypePasswordRecover CodeType = 2
	CodeTypeEmailUpdate     CodeType = 3
	CodeTypeUserAccount     CodeType = 4
	CodeTypeUnknown         CodeType = 100
)

type VerificationCode struct {
	Code    string
	OrgId   string
	UserId  string
	Type    CodeType
	Expires int64
}

// Constants & Enumerations

type ProductStage int

// ProductStageMax is the max number of PayType options.
const ProductStageMax = 3

const (
	ProductStageDefined    ProductStage = 1
	ProductStageProduction ProductStage = 2
	ProductStageRetired    ProductStage = 3
)

type PayType int

// PayTypeMax is the max number of PayType options.
const PayTypeMax = 2

const (
	PayTypePerUser PayType = 1
	PayTypePerOrg  PayType = 2
)

type PayFrequency int

// PayFrequencyMax is the max number of PayType options.
const PayFrequencyMax = 2

const (
	PayFrequencyMonth PayFrequency = 1
	PayFrequencyYear  PayFrequency = 2
)

// Methods

func (a PayFrequency) ToInt() int {
	return int(a)
}

func (a ProductStage) ToInt() int {
	return int(a)
}

// ProductEntity ******************************************************************

type ProductEntity struct {
	Id         string `json:"id"`
	ProductId  string `json:"productId"`
	EntityCode int    `json:"entityCode"`
	Name       string `json:"name"`
}

type ProductEntityRequest struct {
	ProductId  string `json:"productId"`
	EntityCode int    `json:"entityCode"`
	Name       string `json:"name"`
}

// TODO: implement Validate method
// func (r *ProductEntityRequest) Validate(info core.UserInfo) *Error {
// 	if len(r.Name) == 0 || len(r.ProductId) == 0 {
// 		return Error{
// 			Code:        code,
// 			StatusCode:  errorHttpStatusCode(code),
// 			Description: errorText(userInfo, code),
// 			Info:        info,
// 		}
// 	}
// 	return nil
// }

// User Product Access ******************************************************************

type UserAccessInfo struct {
	Product []UserProductAccess
	Entity  []UserProductEntity
}

type UserProductAccess struct {
	Id          string
	OrgId       string
	UserId      string
	ProductId   string
	ProductCode string
	Access      int
}

type UserProductEntity struct {
	Id         string
	EntityId   string
	EntityCode int
	Access     int
}

// Access ******************************************************************

type Access int

const (
	AccessOwner Access = 1
	AccessAdmin Access = 2
	AccessUser  Access = 3
	AccessGuest Access = 4
	AccessOpen  Access = 5
)

func (a Access) ToInt() int {
	return int(a)
}

func (a Access) ToString() string {
	if a == 1 {
		return "owner"
	} else if a == 2 {
		return "admin"
	} else if a == 3 {
		return "user"
	} else if a == 4 {
		return "guest"
	} else {
		return "open"
	}
}

// FIXME: this name is really confusing for what this does.
// TODO: remove and use "HasEqualOrMoreAccessThan" instead.
func (a Access) HasEqualOrMoreAccess(other Access) bool {
	return other <= a
}

func (a Access) HasEqualOrMoreAccessThan(other Access) bool {
	return a <= other
}

// AccountStatus ******************************************************************

type AccountStatus int

const (
	AccountStatusActive    AccountStatus = 1
	AccountStatusInactive  AccountStatus = 2
	AccountStatusFreeTrial AccountStatus = 3
)

func (a AccountStatus) ToString() string {
	if a == 1 {
		return "active"
	} else if a == 2 {
		return "inactive"
	} else {
		return "freeTrial"
	}
}

// UserInfo ******************************************************************

type UserInfo struct {
	UserId              string
	OrgId               string
	Access              Access
	Language            string
	ProductAccess       []string
	ProductEntityAccess []string
}

func (u UserInfo) LanguageTag() language.Tag {
	switch u.Language {
	case "es-419":
		return language.MustParse("es-419")
	default:
		return language.MustParse("en-US")
	}
}

// AccessClaims ******************************************************************

type AccessClaims struct {
	userId                  string
	access                  Access
	orgId                   string
	issuer                  string
	productUserAccess       []string
	productEntityUserAccess []string
}
