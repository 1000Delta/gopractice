package main

import (
	"bytes"
	"log"
	"text/template"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

type netrpcPlugin struct {
	*generator.Generator
}

func (p *netrpcPlugin) Name() string { return "netrpc" }

func (p *netrpcPlugin) Init(g *generator.Generator) { p.Generator = g }

func (p *netrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}

func (p *netrpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P(`import "net/rpc"`)
}

// ServiceSpec 服务规范
type ServiceSpec struct {
	ServiceName string
	MethodList  []ServiceMethodSpec
}

// ServiceMethodSpec 服务内方法规范
type ServiceMethodSpec struct {
	MethodName     string
	InputTypeName  string
	OutputTypeName string
}

// buildServiceSpec 构建服务规范
func (p *netrpcPlugin) buildServiceSpec(
	svc *descriptor.ServiceDescriptorProto,
) *ServiceSpec {
	spec := &ServiceSpec{
		ServiceName: generator.CamelCase(svc.GetName()),
	}

	for _, m := range svc.Method {
		spec.MethodList = append(spec.MethodList, ServiceMethodSpec{
			MethodName:     generator.CamelCase(m.GetName()),
			InputTypeName:  p.TypeName(p.ObjectNamed(m.GetInputType())),
			OutputTypeName: p.TypeName(p.ObjectNamed(m.GetOutputType())),
		})
	}

	return spec
}

func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	spec := p.buildServiceSpec(svc)

	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplService))
	err := t.Execute(&buf, spec)
	if err != nil {
		log.Fatal(err)
	}

	p.P(buf.String())
}

const tmplService = `
type {{ .ServiceName }}Interface interface {
  {{- range .MethodList }}
  {{ .MethodName }}(req {{ .InputTypeName }}, rsp *{{ .OutputTypeName }}) error
  {{- end }}
}
`

func init() {
	generator.RegisterPlugin(&netrpcPlugin{})
}