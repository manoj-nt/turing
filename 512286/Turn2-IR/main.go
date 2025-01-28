package main

// Context struct holds the current configuration
type Context struct {
	Role        string
	Environment string
	IsPremium   bool
}

// FeatureToggle manages the logic for enabling/disabling features
type FeatureToggle struct {
	Context
}

// NewFeatureToggle creates a new FeatureToggle instance with a given context
func NewFeatureToggle(ctx Context) *FeatureToggle {
	return &FeatureToggle{Context: ctx}
}

// IsFeatureAEnabled checks if Feature A is enabled based on the user's role
func (ft *FeatureToggle) IsFeatureAEnabled() bool {
	return ft.Role == "admin"
}

// IsFeatureBEnabled checks if Feature B is enabled based on the environment
func (ft *FeatureToggle) IsFeatureBEnabled() bool {
	return ft.Environment == "staging"
}

// IsFeatureCEnabled checks if Feature C is enabled based on the subscription status
func (ft *FeatureToggle) IsFeatureCEnabled() bool {
	return ft.IsPremium
}
