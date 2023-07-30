package core

import (
	"testing"
)

func TestAccessHasEqualOrMoreAccess(t *testing.T) {
	if AccessOwner.HasEqualOrMoreAccess(AccessOwner) == false {
		t.Errorf("AccessOwner is same as owner")
	}
	if AccessOwner.HasEqualOrMoreAccess(AccessAdmin) {
		t.Errorf("admin doesn't have equal or more access than AccessOwner")
	}
	if AccessOwner.HasEqualOrMoreAccess(AccessUser) {
		t.Errorf("user doesn't have equal or more access than AccessOwner")
	}
	if AccessOwner.HasEqualOrMoreAccess(AccessGuest) {
		t.Errorf("guest doesn't have equal or more access than AccessOwner")
	}
	if AccessOwner.HasEqualOrMoreAccess(AccessOpen) {
		t.Errorf("open doesn't have equal or more access than AccessOwner")
	}

	if AccessAdmin.HasEqualOrMoreAccess(AccessOwner) == false {
		t.Errorf("owner does have more access than AccessAdmin")
	}
	if AccessAdmin.HasEqualOrMoreAccess(AccessAdmin) == false {
		t.Errorf("AccessAdmin is same as admin")
	}
	if AccessAdmin.HasEqualOrMoreAccess(AccessUser) {
		t.Errorf("user doesn't have equal or more access than AccessAdmin")
	}
	if AccessAdmin.HasEqualOrMoreAccess(AccessGuest) {
		t.Errorf("guest doesn't have equal or more access than AccessAdmin")
	}
	if AccessAdmin.HasEqualOrMoreAccess(AccessOpen) {
		t.Errorf("open doesn't have equal or more access than AccessAdmin")
	}

	if AccessUser.HasEqualOrMoreAccess(AccessOwner) == false {
		t.Errorf("owner does have more access than AccessUser")
	}
	if AccessUser.HasEqualOrMoreAccess(AccessAdmin) == false {
		t.Errorf("admin does have more access than AccessUser")
	}
	if AccessUser.HasEqualOrMoreAccess(AccessUser) == false {
		t.Errorf("AccessUser is same as user")
	}
	if AccessUser.HasEqualOrMoreAccess(AccessGuest) {
		t.Errorf("guest doesn't have equal or more access than AccessUser")
	}
	if AccessUser.HasEqualOrMoreAccess(AccessOpen) {
		t.Errorf("open doesn't have equal or more access than AccessUser")
	}

	if AccessGuest.HasEqualOrMoreAccess(AccessOwner) == false {
		t.Errorf("owner does have more access than AccessGuest")
	}
	if AccessGuest.HasEqualOrMoreAccess(AccessAdmin) == false {
		t.Errorf("admin does have more access than AccessGuest")
	}
	if AccessGuest.HasEqualOrMoreAccess(AccessUser) == false {
		t.Errorf("user does have more access than AccessGuest")
	}
	if AccessGuest.HasEqualOrMoreAccess(AccessGuest) == false {
		t.Errorf("AccessGuest is same as guest")
	}
	if AccessGuest.HasEqualOrMoreAccess(AccessOpen) {
		t.Errorf("open doesn't have equal or more access than AccessGuest")
	}

	if AccessOpen.HasEqualOrMoreAccess(AccessOwner) == false {
		t.Errorf("owner does have more access than AccessOpen")
	}
	if AccessOpen.HasEqualOrMoreAccess(AccessAdmin) == false {
		t.Errorf("admin does have more access than AccessOpen")
	}
	if AccessOpen.HasEqualOrMoreAccess(AccessUser) == false {
		t.Errorf("user does have more access than AccessOpen")
	}
	if AccessOpen.HasEqualOrMoreAccess(AccessGuest) == false {
		t.Errorf("guest does have more access than AccessOpen")
	}
	if AccessOpen.HasEqualOrMoreAccess(AccessOpen) == false {
		t.Errorf("AccessOpen is same as open")
	}
}
