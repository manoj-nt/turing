// feature_toggle/feature_toggle_test.go
package featuretoggle

import "testing"

// TestFeatureAForAdmin tests if Feature A is accessible to admin users.
func TestFeatureAForAdmin(t *testing.T) {
	ft := FeatureToggle{UserRole: "admin"}

	if !ft.IsFeatureAEnabled() {
		t.Errorf("Expected Feature A to be enabled for admin users")
	}
}

// TestFeatureBInStaging tests if Feature B is accessible in the staging environment.
func TestFeatureBInStaging(t *testing.T) {
	ft := FeatureToggle{Environment: "staging"}

	if !ft.IsFeatureBEnabled() {
		t.Errorf("Expected Feature B to be enabled in the staging environment")
	}
}

// TestFeatureCForPremiumUser tests if Feature C is accessible to premium users.
func TestFeatureCForPremiumUser(t *testing.T) {
	ft := FeatureToggle{UserSubscription: "premium"}

	if !ft.IsFeatureCEnabled() {
		t.Errorf("Expected Feature C to be enabled for premium users")
	}
}

// TestFeatureAccessFailure tests failure case for accessing a feature.
func TestFeatureAccessFailure(t *testing.T) {
	ft := FeatureToggle{UserRole: "user", Environment: "production", UserSubscription: "free"}

	if ft.IsFeatureAEnabled() {
		t.Errorf("Expected Feature A to be disabled for non-admin users")
	}
	if ft.IsFeatureBEnabled() {
		t.Errorf("Expected Feature B to be disabled outside the staging environment")
	}
	if ft.IsFeatureCEnabled() {
		t.Errorf("Expected Feature C to be disabled for free users")
	}
}
