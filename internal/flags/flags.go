package flags

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// DockerAPIMinVersion is the minimum version of the docker api required to
// use watchtower
const DockerAPIMinVersion string = "1.24"

// RegisterDockerFlags that are used directly by the docker api client
func RegisterDockerFlags(rootCmd *cobra.Command) {
	flags := rootCmd.PersistentFlags()
	flags.StringP("host", "H", viper.GetString("DOCKER_HOST"), "daemon socket to connect to")
	flags.BoolP("tlsverify", "v", viper.GetBool("DOCKER_TLS_VERIFY"), "use TLS and verify the remote")
	flags.StringP("api-version", "a", viper.GetString("DOCKER_API_VERSION"), "api version to use by docker client")
}

// RegisterSystemFlags that are used by watchtower to modify the program flow
func RegisterSystemFlags(rootCmd *cobra.Command) {
	flags := rootCmd.PersistentFlags()
	flags.IntP(
		"interval",
		"i",
		viper.GetInt("WATCHTOWER_POLL_INTERVAL"),
		"poll interval (in seconds)")

	flags.StringP("schedule",
		"s",
		viper.GetString("WATCHTOWER_SCHEDULE"),
		"the cron expression which defines when to update")

	flags.DurationP("stop-timeout",
		"t",
		viper.GetDuration("WATCHTOWER_TIMEOUT"),
		"timeout before a container is forcefully stopped")

	flags.BoolP(
		"no-pull",
		"",
		viper.GetBool("WATCHTOWER_NO_PULL"),
		"do not pull any new images")

	flags.StringToStringP(
		"pull-tags",
		"",
		viper.GetStringMapString("WATCHTOWER_PULL_TAGS"),
		"pull a different tag instead of the used container tag")

	flags.BoolP(
		"no-restart",
		"",
		viper.GetBool("WATCHTOWER_NO_RESTART"),
		"do not restart any containers")

	flags.BoolP(
		"cleanup",
		"c",
		viper.GetBool("WATCHTOWER_CLEANUP"),
		"remove previously used images after updating")

	flags.BoolP(
		"remove-volumes",
		"",
		viper.GetBool("WATCHTOWER_REMOVE_VOLUMES"),
		"remove attached volumes before updating")

	flags.BoolP(
		"label-enable",
		"e",
		viper.GetBool("WATCHTOWER_LABEL_ENABLE"),
		"watch containers where the com.centurylinklabs.watchtower.enable label is true")

	flags.BoolP(
		"debug",
		"d",
		viper.GetBool("WATCHTOWER_DEBUG"),
		"enable debug mode with verbose logging")

	flags.BoolP(
		"monitor-only",
		"m",
		viper.GetBool("WATCHTOWER_MONITOR_ONLY"),
		"Will only monitor for new images, not update the containers")

	flags.BoolP(
		"run-once",
		"R",
		viper.GetBool("WATCHTOWER_RUN_ONCE"),
		"Run once now and exit")

	flags.BoolP(
		"include-stopped",
		"S",
		viper.GetBool("WATCHTOWER_INCLUDE_STOPPED"),
		"Will also include created and exited containers")

	flags.BoolP(
		"revive-stopped",
		"",
		viper.GetBool("WATCHTOWER_REVIVE_STOPPED"),
		"Will also start stopped containers that were updated, if include-stopped is active")

	flags.BoolP(
		"enable-lifecycle-hooks",
		"",
		viper.GetBool("WATCHTOWER_LIFECYCLE_HOOKS"),
		"Enable the execution of commands triggered by pre- and post-update lifecycle hooks")
}

// RegisterNotificationFlags that are used by watchtower to send notifications
func RegisterNotificationFlags(rootCmd *cobra.Command) {
	flags := rootCmd.PersistentFlags()

	flags.StringSliceP(
		"notifications",
		"n",
		viper.GetStringSlice("WATCHTOWER_NOTIFICATIONS"),
		" notification types to send (valid: email, slack, msteams, gotify)")

	flags.StringP(
		"notifications-level",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATIONS_LEVEL"),
		"The log level used for sending notifications. Possible values: panic, fatal, error, warn, info or debug")

	flags.StringP(
		"notification-email-from",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_EMAIL_FROM"),
		"Address to send notification emails from")

	flags.StringP(
		"notification-email-to",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_EMAIL_TO"),
		"Address to send notification emails to")

	flags.IntP(
		"notification-email-delay",
		"",
		viper.GetInt("WATCHTOWER_NOTIFICATION_EMAIL_DELAY"),
		"Delay before sending notifications, expressed in seconds")

	flags.StringP(
		"notification-email-server",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_EMAIL_SERVER"),
		"SMTP server to send notification emails through")

	flags.IntP(
		"notification-email-server-port",
		"",
		viper.GetInt("WATCHTOWER_NOTIFICATION_EMAIL_SERVER_PORT"),
		"SMTP server port to send notification emails through")

	flags.BoolP(
		"notification-email-server-tls-skip-verify",
		"",
		viper.GetBool("WATCHTOWER_NOTIFICATION_EMAIL_SERVER_TLS_SKIP_VERIFY"),
		`
Controls whether watchtower verifies the SMTP server's certificate chain and host name.
Should only be used for testing.
`)

	flags.StringP(
		"notification-email-server-user",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_EMAIL_SERVER_USER"),
		"SMTP server user for sending notifications")

	flags.StringP(
		"notification-email-server-password",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_EMAIL_SERVER_PASSWORD"),
		"SMTP server password for sending notifications")

	flags.StringP(
		"notification-email-subjecttag",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_EMAIL_SUBJECTTAG"),
		"Subject prefix tag for notifications via mail")

	flags.StringP(
		"notification-slack-hook-url",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_SLACK_HOOK_URL"),
		"The Slack Hook URL to send notifications to")

	flags.StringP(
		"notification-slack-identifier",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_SLACK_IDENTIFIER"),
		"A string which will be used to identify the messages coming from this watchtower instance")

	flags.StringP(
		"notification-slack-channel",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_SLACK_CHANNEL"),
		"A string which overrides the webhook's default channel. Example: #my-custom-channel")

	flags.StringP(
		"notification-slack-icon-emoji",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_SLACK_ICON_EMOJI"),
		"An emoji code string to use in place of the default icon")

	flags.StringP(
		"notification-slack-icon-url",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_SLACK_ICON_URL"),
		"An icon image URL string to use in place of the default icon")

	flags.StringP(
		"notification-msteams-hook",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_MSTEAMS_HOOK_URL"),
		"The MSTeams WebHook URL to send notifications to")

	flags.BoolP(
		"notification-msteams-data",
		"",
		viper.GetBool("WATCHTOWER_NOTIFICATION_MSTEAMS_USE_LOG_DATA"),
		"The MSTeams notifier will try to extract log entry fields as MSTeams message facts")

	flags.StringP(
		"notification-gotify-url",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_GOTIFY_URL"),
		"The Gotify URL to send notifications to")
	flags.StringP(
		"notification-gotify-token",
		"",
		viper.GetString("WATCHTOWER_NOTIFICATION_GOTIFY_TOKEN"),
		"The Gotify Application required to query the Gotify API")
}

// SetDefaults provides default values for environment variables
func SetDefaults() {
	viper.AutomaticEnv()
	viper.SetDefault("DOCKER_HOST", "unix:///var/run/docker.sock")
	viper.SetDefault("DOCKER_API_VERSION", DockerAPIMinVersion)
	viper.SetDefault("WATCHTOWER_POLL_INTERVAL", 300)
	viper.SetDefault("WATCHTOWER_TIMEOUT", time.Second*10)
	viper.SetDefault("WATCHTOWER_NOTIFICATIONS", []string{})
	viper.SetDefault("WATCHTOWER_NOTIFICATIONS_LEVEL", "info")
	viper.SetDefault("WATCHTOWER_NOTIFICATION_EMAIL_SERVER_PORT", 25)
	viper.SetDefault("WATCHTOWER_NOTIFICATION_EMAIL_SUBJECTTAG", "")
	viper.SetDefault("WATCHTOWER_NOTIFICATION_SLACK_IDENTIFIER", "watchtower")
}

// EnvConfig translates the command-line options into environment variables
// that will initialize the api client
func EnvConfig(cmd *cobra.Command) error {
	var err error
	var host string
	var tls bool
	var version string

	flags := cmd.PersistentFlags()

	if host, err = flags.GetString("host"); err != nil {
		return err
	}
	if tls, err = flags.GetBool("tlsverify"); err != nil {
		return err
	}
	if version, err = flags.GetString("api-version"); err != nil {
		return err
	}
	if err = setEnvOptStr("DOCKER_HOST", host); err != nil {
		return err
	}
	if err = setEnvOptBool("DOCKER_TLS_VERIFY", tls); err != nil {
		return err
	}
	if err = setEnvOptStr("DOCKER_API_VERSION", version); err != nil {
		return err
	}
	return nil
}

// ReadFlags reads common flags used in the main program flow of watchtower
func ReadFlags(cmd *cobra.Command) (bool, bool, bool, time.Duration) {
	flags := cmd.PersistentFlags()

	var err error
	var cleanup bool
	var noRestart bool
	var monitorOnly bool
	var timeout time.Duration

	if cleanup, err = flags.GetBool("cleanup"); err != nil {
		log.Fatal(err)
	}
	if noRestart, err = flags.GetBool("no-restart"); err != nil {
		log.Fatal(err)
	}
	if monitorOnly, err = flags.GetBool("monitor-only"); err != nil {
		log.Fatal(err)
	}
	if timeout, err = flags.GetDuration("stop-timeout"); err != nil {
		log.Fatal(err)
	}

	return cleanup, noRestart, monitorOnly, timeout
}

func setEnvOptStr(env string, opt string) error {
	if opt == "" || opt == os.Getenv(env) {
		return nil
	}
	err := os.Setenv(env, opt)
	if err != nil {
		return err
	}
	return nil
}

func setEnvOptBool(env string, opt bool) error {
	if opt {
		return setEnvOptStr(env, "1")
	}
	return nil
}
