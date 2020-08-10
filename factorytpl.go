package fixtory

const factoryTpl = `
{{$lowerStructName := .StructName | ToLower }}

type Test{{ .StructName }}Factory interface {
	NewBuilder(bluePrint Test{{ .StructName }}BluePrintFunc, traits ...{{ .StructName }}) Test{{ .StructName }}Builder
	OnBuild(onBuild func(t *testing.T, {{ $lowerStructName }} *{{ .StructName }}))
	Reset()
}

type Test{{ .StructName }}Builder interface {
	Build() *{{ .StructName }}
	Build2() (*{{ .StructName }}, *{{ .StructName }})
	Build3() (*{{ .StructName }}, *{{ .StructName }}, *{{ .StructName }})
	BuildList(n int) []*{{ .StructName }}
	WithZero({{ $lowerStructName }}Fields ...Test{{ .StructName }}Field) Test{{ .StructName }}Builder
	WithReset() Test{{ .StructName }}Builder
	WithEachParams({{ $lowerStructName }}Traits ...{{ .StructName }}) Test{{ .StructName }}Builder
}

type Test{{ .StructName }}BluePrintFunc func(i int, last {{ .StructName }}) {{ .StructName }}

type Test{{ .StructName }}Field string

const (
{{- range .FieldNames }}
	Test{{ $.StructName }}{{ . }} Test{{ $.StructName }}Field = "{{ . }}"
{{- end}}
)

type test{{ .StructName }}Factory struct {
	t       *testing.T
	factory *fixtory.Factory
}

type test{{ .StructName }}Builder struct {
	t       *testing.T
	builder *fixtory.Builder
}

func TestNew{{ .StructName }}Factory(t *testing.T) Test{{ .StructName }}Factory {
	t.Helper()

	return &test{{ .StructName }}Factory{t: t, factory: fixtory.NewFactory(t, {{ .StructName }}{})}
}

func (uf *test{{ .StructName }}Factory) NewBuilder(bluePrint Test{{ .StructName }}BluePrintFunc, {{ $lowerStructName }}Traits ...{{ .StructName }}) Test{{ .StructName }}Builder {
	uf.t.Helper()

	var bp fixtory.BluePrintFunc
	if bluePrint != nil {
		bp = func(i int, last interface{}) interface{} { return bluePrint(i, last.({{ .StructName }})) }
	}
	builder := uf.factory.NewBuilder(bp, fixtory.ConvertToInterfaceArray({{ $lowerStructName }}Traits)...)

	return &test{{ .StructName }}Builder{t: uf.t, builder: builder}
}

func (uf *test{{ .StructName }}Factory) OnBuild(onBuild func(t *testing.T, {{ $lowerStructName }} *{{ .StructName }})) {
	uf.t.Helper()

	uf.factory.OnBuild = func(t *testing.T, v interface{}) { onBuild(t, v.(*{{ .StructName }})) }
}

func (uf *test{{ .StructName }}Factory) Reset() {
	uf.t.Helper()

	uf.factory.Reset()
}

func (ub *test{{ .StructName }}Builder) WithZero({{ $lowerStructName }}Fields ...Test{{ .StructName }}Field) Test{{ .StructName }}Builder {
	ub.t.Helper()

	fields := make([]string, 0, len({{ $lowerStructName }}Fields))
	for _, f := range {{ $lowerStructName }}Fields {
		fields = append(fields, string(f))
	}

	ub.builder = ub.builder.WithZero(fields...)
	return ub
}
func (ub *test{{ .StructName }}Builder) WithReset() Test{{ .StructName }}Builder {
	ub.t.Helper()

	ub.builder = ub.builder.WithReset()
	return ub
}

func (ub *test{{ .StructName }}Builder) WithEachParams({{ $lowerStructName }}Traits ...{{ .StructName }}) Test{{ .StructName }}Builder {
	ub.t.Helper()

	ub.builder = ub.builder.WithEachParams(fixtory.ConvertToInterfaceArray({{ $lowerStructName }}Traits)...)
	return ub
}

func (ub *test{{ .StructName }}Builder) Build() *{{ .StructName }} {
	ub.t.Helper()

	return ub.builder.Build().(*{{ .StructName }})
}

func (ub *test{{ .StructName }}Builder) Build2() (*{{ .StructName }}, *{{ .StructName }}) {
	ub.t.Helper()

	return ub.Build(), ub.Build()
}

func (ub *test{{ .StructName }}Builder) Build3() (*{{ .StructName }}, *{{ .StructName }}, *{{ .StructName }}) {
	ub.t.Helper()

	return ub.Build(), ub.Build(), ub.Build()
}

func (ub *test{{ .StructName }}Builder) BuildList(n int) []*{{ .StructName }} {
	ub.t.Helper()

	{{ $lowerStructName }}s := make([]*{{ .StructName }}, 0, n)
	for i := 0; i < n; i++ {
		{{ $lowerStructName }}s = append({{ $lowerStructName }}s, ub.builder.Build().(*{{ .StructName }}))
	}
	return {{ $lowerStructName }}s
}
`
