
// Code generated by generateconfig; DO NOT EDIT.

package config

type UserPreferences struct {
    Username           string `json:"username"`
    Theme              string `json:"theme"`
    NotificationsEnabled bool `json:"notifications_enabled"`
    Preferences        struct {
        Language string `json:"language"`
        Timezone string `json:"timezone"`
    } `json:"preferences"`
}
