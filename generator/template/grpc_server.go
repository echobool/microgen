package template

import (
	. "github.com/dave/jennifer/jen"
	"github.com/devimteam/microgen/parser"
	"github.com/devimteam/microgen/util"
)

type GRPCServerTemplate struct {
}

// Render whole grpc server file.
//
//		// This file was automatically generated by "microgen" utility.
//		// Please, do not edit.
//		package transportgrpc
//
//		import (
//			transportlayer "github.com/devimteam/go-kit/transportlayer"
//			stringsvc "gitlab.devim.team/protobuf/stringsvc"
//			context "golang.org/x/net/context"
//		)
//
//		type server struct {
//			ts transportlayer.Server
//		}
//
//		func NewServer(endpoints []transportlayer.Endpoint) stringsvc.StringServiceServer {
//			return &server{transportlayer.NewServer(endpoints)}
//		}
//
//		func (s *server) Count(ctx context.Context, req *stringsvc.CountRequest) (*stringsvc.CountResponse, error) {
//			_, resp, err := s.ts.Serve(ctx, req)
//			if err != nil {
//				return nil, err
//			}
//			return resp.(*stringsvc.CountResponse), nil
//		}
//
func (GRPCServerTemplate) Render(i *parser.Interface) *File {
	f := NewFile("transportgrpc")

	f.Type().Id("server").Struct(
		Id("ts").Qual(PackagePathTransportLayer, "Server"),
	)

	f.Func().Id("NewServer").
		Call(Id("endpoints").
			Index().Qual(PackagePathTransportLayer, "Endpoint")).Qual(protobufPath(i), serverStructName(i)).
		Block(
			Return().Op("&").Id("server").Values(
				Qual(PackagePathTransportLayer, "NewServer").Call(Id("endpoints")),
			),
		)
	f.Line()

	for _, signature := range i.FuncSignatures {
		f.Line()
		f.Add(grpcServerFunc(signature, i))
	}

	return f
}

func (GRPCServerTemplate) Path() string {
	return "./transport/grpc/server.go"
}

// Render service interface method for grpc server.
//
//		func (s *server) Count(ctx context.Context, req *stringsvc.CountRequest) (*stringsvc.CountResponse, error) {
//			_, resp, err := s.ts.Serve(ctx, req)
//			if err != nil {
//				return nil, err
//			}
//			return resp.(*stringsvc.CountResponse), nil
//		}
//
func grpcServerFunc(signature *parser.FuncSignature, i *parser.Interface) *Statement {
	return Func().
		Params(Id(util.FirstLowerChar("server")).Op("*").Id("server")).
		Id(signature.Name).
		Call(Id("ctx").Qual(PackagePathNetContext, "Context"), Id("req").Op("*").Qual(protobufPath(i), requestStructName(signature))).
		Params(Op("*").Qual(protobufPath(i), responseStructName(signature)), Error()).
		BlockFunc(grpcServerFuncBody(signature, i))
}

// Render service method body for grpc server.
//
//		_, resp, err := s.ts.Serve(ctx, req)
//		if err != nil {
//			return nil, err
//		}
//		return resp.(*stringsvc.CountResponse), nil
//
func grpcServerFuncBody(signature *parser.FuncSignature, i *parser.Interface) func(g *Group) {
	return func(g *Group) {
		g.List(Id("_"), Id("resp"), Err()).
			Op(":=").
			Id(util.FirstLowerChar("server")).Dot("ts").Dot("Serve").Call(Id("ctx"), Id("req"))

		g.If(Err().Op("!=").Nil()).Block(
			Return().List(Nil(), Err()),
		)

		g.Return().List(Id("resp").Assert(Op("*").Qual(protobufPath(i), responseStructName(signature))), Nil())
	}
}

func protobufPath(iface *parser.Interface) string {
	return "gitlab.devim.team/protobuf/" + iface.PackageName
}

func serverStructName(iface *parser.Interface) string {
	return iface.Name + "Server"
}
