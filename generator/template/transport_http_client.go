package template

import (
	"context"

	. "github.com/dave/jennifer/jen"
	"github.com/devimteam/microgen/generator/write_strategy"
	"github.com/vetcher/go-astra/types"
)

type httpClientTemplate struct {
	info *GenerationInfo
}

func NewHttpClientTemplate(info *GenerationInfo) Template {
	return &httpClientTemplate{
		info: info,
	}
}

func (t *httpClientTemplate) DefaultPath() string {
	return filenameBuilder(PathTransport, "http", "client")
}

func (t *httpClientTemplate) ChooseStrategy(ctx context.Context) (write_strategy.Strategy, error) {
	return write_strategy.NewCreateFileStrategy(t.info.AbsOutputFilePath, t.DefaultPath()), nil
}

func (t *httpClientTemplate) Prepare(ctx context.Context) error {
	return nil
}

// Render http client.
//		// This file was automatically generated by "microgen" utility.
//		// DO NOT EDIT.
//		package transporthttp
//
//		import (
//			svc "github.com/devimteam/microgen/examples/svc"
//			http1 "github.com/devimteam/microgen/examples/svc/transport/converter/http"
//			http "github.com/go-kit/kit/transport/http"
//			url "net/url"
//			strings "strings"
//		)
//
//		func NewHTTPClient(addr string, opts ...http.ClientOption) (svc.StringService, error) {
//			if !strings.HasPrefix(addr, "http") {
//				addr = "http://" + addr
//			}
//			u, err := url.Parse(addr)
//			if err != nil {
//				return nil, err
//			}
//			return &svc.Endpoints{
//				EmptyReqEndpoint: http.NewClient(
//					"POST",
//					u,
//					http1.EncodeHTTPEmptyReqRequest,
//					http1.DecodeHTTPEmptyReqResponse,
//					opts...,
//				).Endpoint(),
//				EmptyRespEndpoint: http.NewClient(
//					"POST",
//					u,
//					http1.EncodeHTTPEmptyRespRequest,
//					http1.DecodeHTTPEmptyRespResponse,
//					opts...,
//				).Endpoint(),
//				TestCaseEndpoint: http.NewClient(
//					"POST",
//					u,
//					http1.EncodeHTTPTestCaseRequest,
//					http1.DecodeHTTPTestCaseResponse,
//					opts...,
//				).Endpoint(),
//			}, nil
//		}
//
func (t *httpClientTemplate) Render(ctx context.Context) write_strategy.Renderer {
	f := NewFile("transporthttp")
	f.ImportAlias(t.info.SourcePackageImport, serviceAlias)
	f.ImportAlias(PackagePathGoKitTransportHTTP, "httpkit")
	f.HeaderComment(t.info.FileHeader)

	f.Func().Id("NewHTTPClient").ParamsFunc(func(p *Group) {
		p.Id("u").Op("*").Qual(PackagePathUrl, "URL")
		p.Id("opts").Op("...").Qual(PackagePathGoKitTransportHTTP, "ClientOption")
	}).Params(
		Qual(t.info.SourcePackageImport+"/transport", EndpointsSetName),
	).Block(
		t.clientBody(ctx),
	)

	if Tags(ctx).Has(TracingMiddlewareTag) {
		f.Line().Func().Id("TracingHTTPClientOptions").Params(
			Id("tracer").Qual(PackagePathOpenTracingGo, "Tracer"),
			Id("logger").Qual(PackagePathGoKitLog, "Logger"),
		).Params(
			Func().Params(Op("[]").Qual(PackagePathGoKitTransportHTTP, "ClientOption")).Params(Op("[]").Qual(PackagePathGoKitTransportHTTP, "ClientOption")),
		).Block(
			Return().Func().Params(Id("opts").Op("[]").Qual(PackagePathGoKitTransportHTTP, "ClientOption")).Params(Op("[]").Qual(PackagePathGoKitTransportHTTP, "ClientOption")).Block(
				Return().Append(Id("opts"), Qual(PackagePathGoKitTransportHTTP, "ClientBefore").Call(
					Line().Qual(PackagePathGoKitTracing, "ContextToHTTP").Call(Id("tracer"), Id("logger")).Op(",").Line(),
				)),
			),
		)
	}

	return f
}

// Render client body.
//		return &svc.Endpoints{
//			EmptyReqEndpoint: http.NewClient(
//				"POST",
//				u,
//				http1.EncodeHTTPEmptyReqRequest,
//				http1.DecodeHTTPEmptyReqResponse,
//				opts...,
//			).Endpoint(),
//			EmptyRespEndpoint: http.NewClient(
//				"POST",
//				u,
//				http1.EncodeHTTPEmptyRespRequest,
//				http1.DecodeHTTPEmptyRespResponse,
//				opts...,
//			).Endpoint(),
//			TestCaseEndpoint: http.NewClient(
//				"POST",
//				u,
//				http1.EncodeHTTPTestCaseRequest,
//				http1.DecodeHTTPTestCaseResponse,
//				opts...,
//			).Endpoint(),
//		}, nil
//
func (t *httpClientTemplate) clientBody(ctx context.Context) *Statement {
	g := &Statement{}
	g.Return(Qual(t.info.SourcePackageImport+"/transport", EndpointsSetName).Values(DictFunc(
		func(d Dict) {
			for _, fn := range t.info.Iface.Methods {
				method := FetchHttpMethodTag(fn.Docs)
				client := &Statement{}
				client.Qual(PackagePathGoKitTransportHTTP, "NewClient").Call(
					Line().Lit(method), Id("u"),
					Line().Id(encodeRequestName(fn)),
					Line().Id(decodeResponseName(fn)),
					Line().Add(t.clientOpts(fn)).Op("...").Line(),
				).Dot("Endpoint").Call()
				d[Id(endpointStructName(fn.Name))] = client
			}
		},
	)))
	return g
}

func (t *httpClientTemplate) clientOpts(fn *types.Function) *Statement {
	s := &Statement{}
	s.Id("opts")
	return s
}
