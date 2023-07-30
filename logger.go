package core

type Logger interface {
	// LogInfo logs request-scoped information.
	LogInfo(s string, userInfo UserInfo)
	// LogWarning logs request-scoped warnings.
	LogWarning(s string, userInfo UserInfo)
	// LogError logs request-scoped errors.
	LogError(s string, userInfo UserInfo)
	// LogSystemInfo logs system-level information.
	LogSystemInfo(s string)
	// LogSystemInfo logs system-level warnings.
	LogSystemWarning(s string)
	// LogSystemInfo logs system-level errors.
	LogSystemError(s string)
}
