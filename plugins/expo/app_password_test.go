package expo

import (
	"testing"
	
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)
	
func TestAppPasswordProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, AppPassword().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{ // TODO: Check if this is correct
				fieldname.Password: "5nldlxt62example",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"EXPO_PASSWORD": "5nldlxt62example",
				},
			},
		},
	})
}

func TestAppPasswordImporter(t *testing.T) {
	plugintest.TestImporter(t, AppPassword().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{ // TODO: Check if this is correct
				"EXPO_PASSWORD": "5nldlxt62example",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Password: "5nldlxt62example",
					},
				},
			},
		},
		// TODO: If you implemented a config file importer, add a test file example in expo/test-fixtures
		// and fill the necessary details in the test template below.
		"config file": {
			Files: map[string]string{
				// "~/path/to/config.yml": plugintest.LoadFixture(t, "config.yml"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
			// 	{
			// 		Fields: map[sdk.FieldName]string{
			// 			fieldname.Token: "5nldlxt62example",
			// 		},
			// 	},
			},
		},
	})
}
