package expo

import (
	"context"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func AppPassword() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.AppPassword,
		DocsURL:       sdk.URL("https://docs.expo.dev/more/expo-cli/"),
		ManagementURL: sdk.URL("https://expo.dev/accounts/"),
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Email,
				MarkdownDescription: "Email or username used to authenticate Expo.",
				Secret:              false,
				Composition: &schema.ValueComposition{
					Length: 16, // TODO Identify the actual length
					Charset: schema.Charset{
						Lowercase: true, // TODO Update the charset because this was copied from `fieldname.Password`
						Digits:    true,
					},
				},
			},
			{
				Name:                fieldname.Password,
				MarkdownDescription: "Password used to authenticate to Expo.",
				Secret:              true,
				Composition: &schema.ValueComposition{
					Length: 16,
					Charset: schema.Charset{
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			{
				Name:                fieldname.OneTimePassword,
				MarkdownDescription: "One-Time Password used to authenticate to Expo (if account has 2FA).",
				Secret:              false,
				Composition: &schema.ValueComposition{
					Length: 6,
					Charset: schema.Charset{
						Digits:    true, // TODO Verify charset is correct
					},
				},
			},
		},
		// TODO Verify no EnvVars or config files are used
		// DefaultProvisioner: provision.EnvVars(defaultEnvVarMapping),
		// Importer: importer.TryAll(
		// 	importer.TryEnvVarPair(defaultEnvVarMapping),
		// 	TryExpoConfigFile(),
		// )
	}
}

var defaultEnvVarMapping = map[string]sdk.FieldName{
	"EXPO_PASSWORD": fieldname.Password, // TODO: Check if this is correct (I'm unsure. I don't use this)
}

// TODO: I've never seen documentation for storing Expo secrets in config files.
// func TryExpoConfigFile() sdk.Importer {
// 	return importer.TryFile("~/path/to/config/file.yml", func(ctx context.Context, contents importer.FileContents, in sdk.ImportInput, out *sdk.ImportAttempt) {
		// var config Config
		// if err := contents.ToYAML(&config); err != nil {
		// 	out.AddError(err)
		// 	return
		// }

		// if config.Password == "" {
		// 	return
		// }

		// out.AddCandidate(sdk.ImportCandidate{
		// 	Fields: map[sdk.FieldName]string{
		// 		fieldname.Password: config.Password,
		// 	},
		// })
// 	})
// }

// TODO: Implement the config file schema
// type Config struct {
//	Password string
// }
