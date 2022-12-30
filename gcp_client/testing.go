package gcp_client

import (
	"context"
	"testing"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/test_helper"
	"github.com/spf13/viper"
)

type TestOptions struct {
	SkipEmptyJsonB bool
}

func MockTestHelper(t *testing.T, table *schema.Table, createService func() (*GcpServices, error), _ TestOptions) {
	testProvider := newTestProvider(table, createService)
	config := "test : test"
	test_helper.RunProviderPullTables(testProvider, config, "./", "*")
}

func newTestProvider(table *schema.Table, builder func() (*GcpServices, error)) *provider.Provider {
	return &provider.Provider{
		Name:      "gcp",
		Version:   "v0.0.1",
		TableList: []*schema.Table{table},
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				services, err := builder()
				if err != nil {
					return nil, schema.NewDiagnostics().AddError(err)
				}
				client := &Client{
					projects:    []string{"testProject"},
					GcpServices: services,
					ProjectId:   "testProject",
				}
				return []any{client}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return ``
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var gcpConfig Config
				err := config.Unmarshal(&gcpConfig)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
