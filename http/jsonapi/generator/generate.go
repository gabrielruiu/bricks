// Copyright © 2018 by PACE Telematics GmbH. All rights reserved.
// Created at 2018/08/24 by Vincent Landgraf

package generator

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/getkin/kin-openapi/openapi3"
)

type buildFunc func(schema *openapi3.Swagger) error

// Generator for go types, requests handler and simple validators
// for the given OpenAPIv3. The OpenAPIv3 schema is expected to comply
// with the json-api specification.
// Everything that doesn't comply to the json-api specification will
// be ignored during generation.
// The Generator doesn't validate necessarily.
type Generator struct {
	goSource            *jen.File
	serviceName         string
	generatedTypes      map[string]bool
	generatedArrayTypes map[string]bool
}

func loadSwaggerFromURI(loader *openapi3.SwaggerLoader, url *url.URL) (*openapi3.Swagger, error) { // nolint: interfacer
	var schema *openapi3.Swagger

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // nolint: errcheck

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	schema, err = loader.LoadSwaggerFromData(body)
	if err != nil {
		return nil, err
	}

	return schema, nil
}

// BuildSource generates the go code in the specified path with specified package name
// based on the passed schema source (url or file path)
func (g *Generator) BuildSource(source, packagePath, packageName string) (string, error) {
	loader := openapi3.NewSwaggerLoader()
	var schema *openapi3.Swagger

	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		loc, err := url.Parse(source)
		if err != nil {
			return "", err
		}

		schema, err = loadSwaggerFromURI(loader, loc)
		if err != nil {
			return "", err
		}
	} else {
		// read spec
		data, err := ioutil.ReadFile(source) // nolint: gosec
		if err != nil {
			return "", err
		}

		// parse spec
		schema, err = loader.LoadSwaggerFromData(data)
		if err != nil {
			return "", err
		}
	}

	return g.BuildSchema(schema, packagePath, packageName)
}

// BuildSchema generates the go code in the specified path with specified package name
// based on the passed schema
func (g *Generator) BuildSchema(schema *openapi3.Swagger, packagePath, packageName string) (string, error) {
	g.generatedTypes = make(map[string]bool)
	g.generatedArrayTypes = make(map[string]bool)

	g.goSource = jen.NewFilePathName(packagePath, packageName)
	g.goSource.PackageComment("// Code generated by github.com/pace/bricks DO NOT EDIT.")
	g.goSource.ImportAlias(pkgJSONAPIMetrics, "metrics")
	g.goSource.ImportAlias(pkgOpentracing, "opentracing")

	g.serviceName = packageName

	buildFuncs := []buildFunc{
		g.BuildTypes,
		g.BuildHandler,
	}

	for _, bf := range buildFuncs {
		err := bf(schema)
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("%#v", g.goSource), nil
}
