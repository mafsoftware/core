package core

import (
	"golang.org/x/text/language"
)

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
	UserId                  string
	Access                  Access
	OrgId                   string
	Issuer                  string
	IssuesAt                int64
	ExpiresAt               int64
	ProductUserAccess       []string
	ProductEntityUserAccess []string
}
