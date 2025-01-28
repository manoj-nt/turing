package main

import (
	"testing"
)

// TestFeatureA_AdminUser tests if Feature A is enabled for admin users
func TestFeatureA_AdminUser(t *testing.T) {
	ctx := Context{
		Role: "admin",
	}
	ft := NewFeatureToggle(ctx)

	if !ft.IsFeatureAEnabled() {
		t.Errorf("Expected Feature A to be enabled for admin user")
	}
}

// TestFeatureA_StandardUser tests if Feature A is NOT enabled for standard users
func TestFeatureA_StandardUser(t *testing.T) {
	ctx := Context{
		Role: "standard",
	}
	ft := NewFeatureToggle(ctx)

	if ft.IsFeatureAEnabled() {
		t.Errorf("Expected Feature A to be disabled for standard user")
	}
}

// TestFeatureB_Staging tests if Feature B is enabled in the staging environment
func TestFeatureB_Staging(t *testing.T) {
	ctx := Context{
		Environment: "staging",
	}
	ft := NewFeatureToggle(ctx)

	if !ft.IsFeatureBEnabled() {
		t.Errorf("Expected Feature B to be enabled in staging environment")
	}
}

// TestFeatureB_Production tests if Feature B is NOT enabled in the production environment
func TestFeatureB_Production(t *testing.T) {
	ctx := Context{
		Environment: "production",
	}
	ft := NewFeatureToggle(ctx)

	if ft.IsFeatureBEnabled() {
		t.Errorf("Expected Feature B to be disabled in production environment")
	}
}

// TestFeatureC_PremiumUser tests if Feature C is enabled for premium users
func TestFeatureC_PremiumUser(t *testing.T) {
	ctx := Context{
		IsPremium: true,
	}
	ft := NewFeatureToggle(ctx)

	if !ft.IsFeatureCEnabled() {
		t.Errorf("Expected Feature C to be enabled for premium user")
	}
}

// TestFeatureC_StandardUser tests if Feature C is NOT enabled for standard users
func TestFeatureC_StandardUser(t *testing.T) {
	ctx := Context{
		IsPremium: false,
	}
	ft := NewFeatureToggle(ctx)

	if ft.IsFeatureCEnabled() {
		t.Errorf("Expected Feature C to be disabled for standard user")
	}
}

// TestFailure tests the failure case where the user does not meet the feature conditions
func TestFailure(t *testing.T) {
	ctx := Context{
		Role:        "standard",   // Non-admin user
		Environment: "production", // Not staging
		IsPremium:   false,        // Not a premium user
	}
	ft := NewFeatureToggle(ctx)

	// Expecting all features to be disabled
	if ft.IsFeatureAEnabled() || ft.IsFeatureBEnabled() || ft.IsFeatureCEnabled() {
		t.Errorf("Expected all features to be disabled for this user")
	}
}
