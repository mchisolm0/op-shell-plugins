package expo

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func ExpoCLI() schema.Executable {
	return schema.Executable{
		Name:      "Expo CLI",
		Runs:      []string{"npx expo login"},
		DocsURL:   sdk.URL("https://docs.expo.dev/more/expo-cli/"),
		NeedsAuth: needsauth.IfAll(
			needsauth.NotForHelpOrVersion(),
		),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.AppPassword, // TODO Should this have email and OTP, also?
			},
		},
	}
}
