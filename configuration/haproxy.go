package configuration

type HAProxy struct {
	TemplatePath          string
	OutputPath            string
	ReloadCommand         string
	ReloadValidateCommand string
	ReloadCleanupCommand  string
}
