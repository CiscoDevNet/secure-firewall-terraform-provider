// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

//go:build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"gopkg.in/yaml.v3"
)

const (
	definitionsPath   = "./gen/definitions/"
	providerTemplate  = "./gen/templates/provider.go"
	providerLocation  = "./internal/provider/provider.go"
	changelogTemplate = "./gen/templates/changelog.md.tmpl"
	changelogLocation = "./templates/guides/changelog.md.tmpl"
	changelogOriginal = "./CHANGELOG.md"
)

type t struct {
	path   string
	prefix string
	suffix string
}

var templates = []t{
	{
		path:   "./gen/templates/model.go",
		prefix: "./internal/provider/model_fmc_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source.go",
		prefix: "./internal/provider/data_source_fmc_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/data_source_test.go",
		prefix: "./internal/provider/data_source_fmc_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/resource.go",
		prefix: "./internal/provider/resource_fmc_",
		suffix: ".go",
	},
	{
		path:   "./gen/templates/resource_test.go",
		prefix: "./internal/provider/resource_fmc_",
		suffix: "_test.go",
	},
	{
		path:   "./gen/templates/data-source.tf",
		prefix: "./examples/data-sources/fmc_",
		suffix: "/data-source.tf",
	},
	{
		path:   "./gen/templates/resource.tf",
		prefix: "./examples/resources/fmc_",
		suffix: "/resource.tf",
	},
	{
		path:   "./gen/templates/import.sh",
		prefix: "./examples/resources/fmc_",
		suffix: "/import.sh",
	},
}

type YamlConfig struct {
	Name                     string                `yaml:"name"`
	NoDataSource             bool                  `yaml:"no_data_source"`
	NoResource               bool                  `yaml:"no_resource"`
	NoImport                 bool                  `yaml:"no_import"`
	TfName                   string                `yaml:"tf_name"`
	RestEndpoint             string                `yaml:"rest_endpoint"`
	PutCreate                bool                  `yaml:"put_create"`
	NoUpdate                 bool                  `yaml:"no_update"`
	NoDelete                 bool                  `yaml:"no_delete"`
	MinimumVersion           string                `yaml:"minimum_version"`
	MinimumVersionCreate     string                `yaml:"minimum_version_create"`
	MinimumVersionBulkCreate string                `yaml:"minimum_version_bulk_create"`
	MinimumVersionBulkDelete string                `yaml:"minimum_version_bulk_delete"`
	DsDescription            string                `yaml:"ds_description"`
	ResDescription           string                `yaml:"res_description"`
	DocCategory              string                `yaml:"doc_category"`
	ExcludeTest              bool                  `yaml:"exclude_test"`
	SkipTest                 bool                  `yaml:"skip_test"`
	SkipMinimumTest          bool                  `yaml:"skip_minimum_test"`
	Attributes               []YamlConfigAttribute `yaml:"attributes"`
	TestTags                 []string              `yaml:"test_tags"`
	TestPrerequisites        string                `yaml:"test_prerequisites"`
	IsBulk                   bool                  `yaml:"is_bulk"`
	BulkSizeCreate           int                   `yaml:"bulk_size_create"`
	ImportNameQuery          bool                  `yaml:"import_name_query"`
}

type YamlConfigAttribute struct {
	ModelName            string                `yaml:"model_name"`
	TfName               string                `yaml:"tf_name"`
	Type                 string                `yaml:"type"`
	ElementType          string                `yaml:"element_type"`
	DataPath             []string              `yaml:"data_path"`
	Id                   bool                  `yaml:"id"`
	ResourceId           bool                  `yaml:"resource_id"`
	Reference            bool                  `yaml:"reference"`
	RequiresReplace      bool                  `yaml:"requires_replace"`
	Mandatory            bool                  `yaml:"mandatory"`
	WriteOnly            bool                  `yaml:"write_only"`
	WriteChangesOnly     bool                  `yaml:"write_changes_only"`
	ExcludeTest          bool                  `yaml:"exclude_test"`
	ExcludeExample       bool                  `yaml:"exclude_example"`
	Description          string                `yaml:"description"`
	Example              string                `yaml:"example"`
	EnumValues           []string              `yaml:"enum_values"`
	MinList              int64                 `yaml:"min_list"`
	MaxList              int64                 `yaml:"max_list"`
	MinInt               int64                 `yaml:"min_int"`
	MaxInt               int64                 `yaml:"max_int"`
	MinFloat             float64               `yaml:"min_float"`
	MaxFloat             float64               `yaml:"max_float"`
	MapKeyExample        string                `yaml:"map_key_example"`
	OrderedList          bool                  `yaml:"ordered_list"`
	StringPatterns       []string              `yaml:"string_patterns"`
	StringMinLength      int64                 `yaml:"string_min_length"`
	StringMaxLength      int64                 `yaml:"string_max_length"`
	Computed             string                `yaml:"computed"`
	ComputedRefreshValue string                `yaml:"computed_refresh_value"`
	DefaultValue         string                `yaml:"default_value"`
	Value                string                `yaml:"value"`
	TestValue            string                `yaml:"test_value"`
	MinimumTestValue     string                `yaml:"minimum_test_value"`
	TestTags             []string              `yaml:"test_tags"`
	DataSourceQuery      bool                  `yaml:"data_source_query"`
	Attributes           []YamlConfigAttribute `yaml:"attributes"`
	GoTypeName           string
}

// Templating helper function to convert TF name to GO name
func ToGoName(s string) string {
	var g []string

	p := strings.Split(s, "_")

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	s = strings.Join(g, "")
	return s
}

// Templating helper function to convert string to camel case
func CamelCase(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "-", " ")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.Title(value))
	}
	return strings.Join(g, "")
}

// Templating helper function to convert string to snake case
func SnakeCase(s string) string {
	var g []string

	s = strings.ReplaceAll(s, "-", " ")
	p := strings.Fields(s)

	for _, value := range p {
		g = append(g, strings.ToLower(value))
	}
	return strings.Join(g, "_")
}

// Templating helper function to fail a template mid-way
func Errorf(s string, args ...any) (struct{}, error) {
	return struct{}{}, fmt.Errorf(s, args...)
}

// Templating helper function to build a SJSON path
func BuildPath(s []string) string {
	return strings.Join(s, ".")
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// Templating helper function to check if any of the attributes is a data source query
func HasDataSourceQuery(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.DataSourceQuery {
			return true
		}
	}
	return false
}

// Templating helper function to return Data Source Query Attribute
func GetDataSourceQueryAttribute(config YamlConfig) YamlConfigAttribute {
	for _, attr := range config.Attributes {
		if attr.DataSourceQuery {
			return attr
		}
	}
	return YamlConfigAttribute{}
}

// Templating helper function to return Attribute by TfName
func GetAttributeByTfName(attributes []YamlConfigAttribute, tfName string) YamlConfigAttribute {
	for _, attr := range attributes {
		if attr.TfName == tfName {
			return attr
		}
	}
	return YamlConfigAttribute{}
}

// Templating helper function to return true if id included in attributes
func HasId(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Id {
			return true
		}
	}
	return false
}

// Templating helper function to return true if reference included in attributes
func HasReference(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.Reference {
			return true
		}
	}
	return false
}

// Templating helper function to return true if reference included in attributes
func HasResourceId(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.ResourceId {
			return true
		}
		if len(attr.Attributes) > 0 {
			if HasResourceId(attr.Attributes) {
				return true
			}
		}
	}
	return false
}

// Templating helper function to return true if any of the attributes has `requires_replace: true`
func HasRequiresReplace(attributes []YamlConfigAttribute) bool {
	for _, attr := range attributes {
		if attr.RequiresReplace {
			return true
		}
		if len(attr.Attributes) > 0 {
			if HasRequiresReplace(attr.Attributes) {
				return true
			}
		}
	}
	return false
}

// Templating helper function to return true if type is a list or set without nested elements
func IsListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list without nested elements
func IsList(attribute YamlConfigAttribute) bool {
	if attribute.Type == "List" && attribute.ElementType != "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a set without nested elements
func IsSet(attribute YamlConfigAttribute) bool {
	if attribute.Type == "Set" && (attribute.ElementType != "" || len(attribute.Attributes) > 0) {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set of strings without nested elements
func IsStringListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "String" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set of integers without nested elements
func IsInt64ListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "Int64" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or a map or a set, anyway with nested elements
func IsNestedListMapSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Map" || attribute.Type == "Set") && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list or set with nested elements
func IsNestedListSet(attribute YamlConfigAttribute) bool {
	if (attribute.Type == "List" || attribute.Type == "Set") && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a list with nested elements
func IsNestedList(attribute YamlConfigAttribute) bool {
	if attribute.Type == "List" && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a map with nested elements
func IsNestedMap(attribute YamlConfigAttribute) bool {
	if attribute.Type == "Map" && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if type is a set with nested elements
func IsNestedSet(attribute YamlConfigAttribute) bool {
	if attribute.Type == "Set" && attribute.ElementType == "" {
		return true
	}
	return false
}

// Templating helper function to return true if object is domain dependent (exists within a domain)
func IsDomainDependent(config YamlConfig) bool {
	return strings.Contains(config.RestEndpoint, "{DOMAIN_UUID}")
}

// Templating helper function to return number of import parts
func ImportParts(attributes []YamlConfigAttribute) int {
	parts := 1
	for _, attr := range attributes {
		if attr.Reference {
			parts += 1
		} else if attr.Id {
			parts += 1
		}
	}
	return parts
}

// Templating helper function to subtract one number from another
func Subtract(a, b int) int {
	return a - b
}

// Map of templating functions
var functions = template.FuncMap{
	"toGoName":                    ToGoName,
	"camelCase":                   CamelCase,
	"snakeCase":                   SnakeCase,
	"sprintf":                     fmt.Sprintf,
	"errorf":                      Errorf,
	"toLower":                     strings.ToLower,
	"path":                        BuildPath,
	"hasDataSourceQuery":          HasDataSourceQuery,
	"getDataSourceQueryAttribute": GetDataSourceQueryAttribute,
	"getAttributeByTfName":        GetAttributeByTfName,
	"hasId":                       HasId,
	"hasReference":                HasReference,
	"hasResourceId":               HasResourceId,
	"hasRequiresReplace":          HasRequiresReplace,
	"isListSet":                   IsListSet,
	"isList":                      IsList,
	"isSet":                       IsSet,
	"isStringListSet":             IsStringListSet,
	"isInt64ListSet":              IsInt64ListSet,
	"isNestedListMapSet":          IsNestedListMapSet,
	"isNestedListSet":             IsNestedListSet,
	"isNestedList":                IsNestedList,
	"isNestedMap":                 IsNestedMap,
	"isNestedSet":                 IsNestedSet,
	"isDomainDependent":           IsDomainDependent,
	"importParts":                 ImportParts,
	"subtract":                    Subtract,
}

func (attr *YamlConfigAttribute) init(parentGoTypeName string) error {
	// Augument
	if attr.TfName == "" {
		var words []string
		l := 0
		for s := attr.ModelName; s != ""; s = s[l:] {
			l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
			if l <= 0 {
				l = len(s)
			}
			words = append(words, strings.ToLower(s[:l]))
		}
		attr.TfName = strings.Join(words, "_")
	}

	attr.GoTypeName = parentGoTypeName + ToGoName(attr.TfName)

	// Validate
	if len(attr.Attributes) > 0 && attr.Type != "List" && attr.Type != "Map" && attr.Type != "Set" {
		return fmt.Errorf("%q has type %q which cannot have `attributes`: instead use type List, Map, Set",
			attr.TfName, attr.Type)
	}

	if len(attr.Attributes) > 0 && attr.ElementType != "" {
		return fmt.Errorf("%q: either `attributes` or `element_type` can be specified, but not both", attr.TfName)
	}

	if attr.Type == "Map" && attr.ElementType != "" {
		return fmt.Errorf("%q: the `element_type` is not yet implemented for type Map", attr.TfName)
	}

	if attr.OrderedList {
		if attr.Type != "List" {
			return fmt.Errorf("%q has type %q which cannot use `ordered_list`: instead use type List",
				attr.TfName, attr.Type)
		}
		if HasId(attr.Attributes) {
			return fmt.Errorf("%q: the `ordered_list: true` conflicts with sub-attributes having `id: true`, as it treats list index ([i]) as the only unique id",
				attr.TfName)
		}
	}

	if attr.Type == "Map" && HasId(attr.Attributes) {
		return fmt.Errorf("Map %q cannot contain sub-attributes with `id: true`, as it treats map key ([k]) as the only unique id",
			attr.TfName)
	}

	// Recurse
	for i := range attr.Attributes {
		if err := attr.Attributes[i].init(attr.GoTypeName); err != nil {
			return err
		}
	}

	return nil
}

func NewYamlConfig(bytes []byte) (YamlConfig, error) {
	var config YamlConfig
	var hasDataSourceQuery bool = false

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return config, err
	}

	for i := range config.Attributes {
		if err := config.Attributes[i].init(CamelCase(config.Name)); err != nil {
			return YamlConfig{}, err
		}
		if config.Attributes[i].DataSourceQuery {
			if hasDataSourceQuery {
				return YamlConfig{}, fmt.Errorf("Multiple `data_source_query` attributes found. Only one is allowed.")
			}
			hasDataSourceQuery = true
		}
	}
	if config.DsDescription == "" {
		config.DsDescription = fmt.Sprintf("This data source reads the %s.", config.Name)
	}
	if config.ResDescription == "" {
		name := strings.ToLower(config.Name)
		if config.IsBulk {
			config.ResDescription = fmt.Sprintf("This resource manages %s through bulk operations.", config.Name)
		} else if strings.HasPrefix(name, "a") || strings.HasPrefix(name, "e") || strings.HasPrefix(name, "i") || strings.HasPrefix(name, "o") || strings.HasPrefix(name, "u") {
			config.ResDescription = fmt.Sprintf("This resource manages an %s.", config.Name)
		} else {
			config.ResDescription = fmt.Sprintf("This resource manages a %s.", config.Name)
		}
	}
	if config.TfName == "" {
		config.TfName = strings.Replace(config.Name, " ", "_", -1)
	}

	return config, nil
}

func getTemplateSection(content, name string) string {
	scanner := bufio.NewScanner(strings.NewReader(content))
	result := ""
	foundSection := false
	beginRegex := regexp.MustCompile(`\/\/template:begin\s` + name + `$`)
	endRegex := regexp.MustCompile(`\/\/template:end\s` + name + `$`)
	for scanner.Scan() {
		line := scanner.Text()
		if !foundSection {
			match := beginRegex.MatchString(line)
			if match {
				foundSection = true
				result += line + "\n"
			}
		} else {
			result += line + "\n"
			match := endRegex.MatchString(line)
			if match {
				foundSection = false
			}
		}
	}
	return result
}

func renderTemplate(templatePath, outputPath string, config interface{}) {
	file, err := os.Open(templatePath)
	if err != nil {
		log.Fatalf("Error opening template: %v", err)
	}
	defer file.Close()

	// skip first line with 'build-ignore' directive for go files
	scanner := bufio.NewScanner(file)
	if strings.HasSuffix(templatePath, ".go") {
		scanner.Scan()
	}
	var temp string
	for scanner.Scan() {
		temp = temp + scanner.Text() + "\n"
	}

	template, err := template.New(path.Base(templatePath)).Funcs(functions).Parse(temp)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	output := new(bytes.Buffer)
	err = template.Execute(output, config)
	if err != nil {
		log.Fatalf("Error executing template for %s: %v", outputPath, err)
	}

	outputFile := filepath.Join(outputPath)
	existingFile, err := os.Open(outputPath)
	if err != nil {
		os.MkdirAll(filepath.Dir(outputFile), 0755)
	} else if strings.HasSuffix(templatePath, ".go") {
		existingScanner := bufio.NewScanner(existingFile)
		var newContent string
		currentSectionName := ""
		beginRegex := regexp.MustCompile(`\/\/template:begin\s(.*?)$`)
		endRegex := regexp.MustCompile(`\/\/template:end\s(.*?)$`)
		for existingScanner.Scan() {
			line := existingScanner.Text()
			if currentSectionName == "" {
				matches := beginRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] != "" {
					currentSectionName = matches[1]
				} else {
					newContent += line + "\n"
				}
			} else {
				matches := endRegex.FindStringSubmatch(line)
				if len(matches) > 1 && matches[1] == currentSectionName {
					currentSectionName = ""
					newSection := getTemplateSection(string(output.Bytes()), matches[1])
					newContent += newSection
				}
			}
		}
		output = bytes.NewBufferString(newContent)
	}
	// write to output file
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	f.Write(output.Bytes())
}

func main() {
	// Load configs
	var configs []YamlConfig
	files, _ := os.ReadDir(definitionsPath)

	for _, filename := range files {
		path := filepath.Join(definitionsPath, filename.Name())
		bytes, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Error reading file %q: %v", path, err)
		}

		config, err := NewYamlConfig(bytes)
		if err != nil {
			log.Fatalf("Error parsing %q: %v", path, err)
		}
		configs = append(configs, config)
	}

	for i := range configs {
		// Iterate over templates and render files
		for _, t := range templates {
			if configs[i].NoImport && t.path == "./gen/templates/import.sh" ||
				configs[i].NoDataSource && t.path == "./gen/templates/data_source.go" ||
				configs[i].NoDataSource && t.path == "./gen/templates/data_source_test.go" ||
				configs[i].SkipTest && t.path == "./gen/templates/data_source_test.go" ||
				configs[i].NoDataSource && t.path == "./gen/templates/data-source.tf" ||
				configs[i].NoResource && t.path == "./gen/templates/resource.go" ||
				configs[i].SkipTest && t.path == "./gen/templates/resource_test.go" ||
				configs[i].NoResource && t.path == "./gen/templates/resource_test.go" ||
				configs[i].NoResource && t.path == "./gen/templates/resource.tf" ||
				// Data source test cannot be generated if there is no corresponding resource
				configs[i].NoResource && t.path == "./gen/templates/data_source_test.go" {
				continue
			}
			renderTemplate(t.path, t.prefix+SnakeCase(configs[i].Name)+t.suffix, configs[i])
		}
	}

	// render provider.go
	renderTemplate(providerTemplate, providerLocation, configs)

	changelog, err := os.ReadFile(changelogOriginal)
	if err != nil {
		log.Fatalf("Error reading %q: %v", changelogOriginal, err)
	}
	renderTemplate(changelogTemplate, changelogLocation, string(changelog))
}
