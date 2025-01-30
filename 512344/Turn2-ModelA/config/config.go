package config

type Preferences struct {
	Language string `json:"language"`
	Timezone string `json:"timezone"`
}

type UserPreferences struct {
	Username             string      `json:"username"`
	Theme                string      `json:"theme"`
	NotificationsEnabled bool        `json:"notifications_enabled"`
	Preferences          Preferences `json:"preferences"`
}
