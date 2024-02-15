package clientgen

import (
	_ "embed" // blank import to embed the template
	"fmt"
	"sort"

	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
	gqlgencConfig "github.com/sonatard/gqlgenc/config"
)

//go:embed template.gotpl
var template string

func RenderTemplate(cfg *config.Config, query *Query, mutation *Mutation, fragments []*Fragment, operations []*Operation, operationResponses []*OperationResponse, generateCfg *gqlgencConfig.GenerateConfig, client config.PackageConfig) error {
	// sort to ensure a deterministic output
	sort.Slice(fragments, func(i, j int) bool { return fragments[i].Name < fragments[j].Name })
	sort.Slice(operations, func(i, j int) bool { return operations[i].Name < operations[j].Name })
	sort.Slice(operationResponses, func(i, j int) bool { return operationResponses[i].Name < operationResponses[j].Name })

	if err := templates.Render(templates.Options{
		PackageName: client.Package,
		Filename:    client.Filename,
		Template:    template,
		Data: map[string]interface{}{
			"Query":               query,
			"Mutation":            mutation,
			"Fragment":            fragments,
			"Operation":           operations,
			"OperationResponse":   operationResponses,
			"GenerateClient":      generateCfg.ShouldGenerateClient(),
			"ClientInterfaceName": generateCfg.GetClientInterfaceName(),
		},
		Packages:   cfg.Packages,
		PackageDoc: "// Code generated by github.com/sonatard/gqlgenc, DO NOT EDIT.\n",
	}); err != nil {
		return fmt.Errorf("%s generating failed: %w", client.Filename, err)
	}

	return nil
}
