package fixtory

const fixtoryFileTpl = `
// Code generated by {{ .GeneratorName }}; DO NOT EDIT.

package {{ .PackageName }}

import (
{{- range .ImportPackages }}
	{{ . }}
{{- end}}
	"github.com/ikozinov/fixtory"
	"testing"
)

{{ .Body }}
`

const factoryTpl = `
{{ $lowerStructName := .StructName | ToLower }}
{{ $factoryInterface := printf "%s%s" .StructName "Factory" }}
{{ $builderInterface := printf "%s%s" .StructName "Builder" }}
{{ $factory := printf "%s%s" $lowerStructName "Factory" }}
{{ $builder := printf "%s%s" $lowerStructName "Builder" }}
{{ $fieldType := printf "%s%s" .StructName "Field" }}

type {{ $factoryInterface }} interface {
	NewBuilder(bluePrint {{ .StructName }}BluePrintFunc, traits ...{{ .Struct }}) {{ $builderInterface }}
	OnBuild(onBuild func(t *testing.T, {{ $lowerStructName }} *{{ .Struct }}))
	Reset()
}

type {{ $builderInterface }} interface {
	EachParam({{ $lowerStructName }}Params ...{{ .Struct }}) {{ $builderInterface }}
	Zero({{ $lowerStructName }}Fields ...{{ $fieldType }}) {{ $builderInterface }}
	ResetAfter() {{ $builderInterface }}

	Build() *{{ .Struct }}
	Build2() (*{{ .Struct }}, *{{ .Struct }})
	Build3() (*{{ .Struct }}, *{{ .Struct }}, *{{ .Struct }})
	BuildList(n int) []*{{ .Struct }}
}

type {{ .StructName }}BluePrintFunc func(i int, last {{ .Struct }}) {{ .Struct }}

type {{ $fieldType }} string

const (
{{- range .FieldNames }}
	{{ $.StructName }}{{ . }}Field {{ $fieldType }} = "{{ . }}"
{{- end}}
)

type {{ $factory }} struct {
	t       *testing.T
	factory *fixtory.Factory
}

type {{ $builder }} struct {
	t       *testing.T
	builder *fixtory.Builder
}

func New{{ .StructName }}Factory(t *testing.T) {{ $factoryInterface }} {
	t.Helper()

	return &{{ $factory }}{t: t, factory: fixtory.NewFactory(t, {{ .Struct }}{})}
}

func (uf *{{ $factory }}) NewBuilder(bluePrint {{ .StructName }}BluePrintFunc, {{ $lowerStructName }}Traits ...{{ .Struct }}) {{ $builderInterface }} {
	uf.t.Helper()

	var bp fixtory.BluePrintFunc
	if bluePrint != nil {
		bp = func(i int, last interface{}) interface{} { return bluePrint(i, last.({{ .Struct }})) }
	}
	builder := uf.factory.NewBuilder(bp, fixtory.ConvertToInterfaceArray({{ $lowerStructName }}Traits)...)

	return &{{ $builder }}{t: uf.t, builder: builder}
}

func (uf *{{ $factory }}) OnBuild(onBuild func(t *testing.T, {{ $lowerStructName }} *{{ .Struct }})) {
	uf.t.Helper()

	uf.factory.OnBuild = func(t *testing.T, v interface{}) { onBuild(t, v.(*{{ .Struct }})) }
}

func (uf *{{ $factory }}) Reset() {
	uf.t.Helper()

	uf.factory.Reset()
}

func (ub *{{ $builder }}) Zero({{ $lowerStructName }}Fields ...{{ $fieldType }}) {{ $builderInterface }} {
	ub.t.Helper()

	fields := make([]string, 0, len({{ $lowerStructName }}Fields))
	for _, f := range {{ $lowerStructName }}Fields {
		fields = append(fields, string(f))
	}

	ub.builder = ub.builder.Zero(fields...)
	return ub
}
func (ub *{{ $builder }}) ResetAfter() {{ $builderInterface }} {
	ub.t.Helper()

	ub.builder = ub.builder.ResetAfter()
	return ub
}

func (ub *{{ $builder }}) EachParam({{ $lowerStructName }}Params ...{{ .Struct }}) {{ $builderInterface }} {
	ub.t.Helper()

	ub.builder = ub.builder.EachParam(fixtory.ConvertToInterfaceArray({{ $lowerStructName }}Params)...)
	return ub
}

func (ub *{{ $builder }}) Build() *{{ .Struct }} {
	ub.t.Helper()

	return ub.builder.Build().(*{{ .Struct }})
}

func (ub *{{ $builder }}) Build2() (*{{ .Struct }}, *{{ .Struct }}) {
	ub.t.Helper()

	list := ub.BuildList(2)
	return list[0], list[1]
}

func (ub *{{ $builder }}) Build3() (*{{ .Struct }}, *{{ .Struct }}, *{{ .Struct }}) {
	ub.t.Helper()

	list := ub.BuildList(3)
	return list[0], list[1], list[2]
}

func (ub *{{ $builder }}) BuildList(n int) []*{{ .Struct }} {
	ub.t.Helper()

	{{ $lowerStructName }}s := make([]*{{ .Struct }}, 0, n)
	for _, {{ $lowerStructName }} := range ub.builder.BuildList(n) {
		{{ $lowerStructName }}s = append({{ $lowerStructName }}s, {{ $lowerStructName }}.(*{{ .Struct }}))
	}
	return {{ $lowerStructName }}s
}
`
