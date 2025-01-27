// feature_toggle/feature_toggle.go
package featuretoggle

// FeatureToggle allows checking feature availability based on context
type FeatureToggle struct {
	UserRole         string
	Environment      string
	UserSubscription string
}

// IsFeatureAEnabled checks if Feature A is enabled for the user role
func (ft *FeatureToggle) IsFeatureAEnabled() bool {
	return ft.UserRole == "admin"
}

// IsFeatureBEnabled checks if Feature B is enabled in the current environment
func (ft *FeatureToggle) IsFeatureBEnabled() bool {
	return ft.Environment == "staging"
}

// IsFeatureCEnabled checks if Feature C is enabled for the user subscription
func (ft *FeatureToggle) IsFeatureCEnabled() bool {
	return ft.UserSubscription == "premium"
}
