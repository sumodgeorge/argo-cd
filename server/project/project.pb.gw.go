// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: server/project/project.proto

/*
Package project is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package project

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_ProjectService_Create_0(ctx context.Context, marshaler runtime.Marshaler, client ProjectServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ProjectCreateRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Create(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_ProjectService_List_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_ProjectService_List_0(ctx context.Context, marshaler runtime.Marshaler, client ProjectServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ProjectQuery
	var metadata runtime.ServerMetadata

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_ProjectService_List_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.List(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_ProjectService_Get_0(ctx context.Context, marshaler runtime.Marshaler, client ProjectServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ProjectQuery
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "name")
	}

	protoReq.Name, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "name", err)
	}

	msg, err := client.Get(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_ProjectService_Update_0(ctx context.Context, marshaler runtime.Marshaler, client ProjectServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ProjectUpdateRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["project.metadata.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "project.metadata.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "project.metadata.name", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "project.metadata.name", err)
	}

	msg, err := client.Update(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_ProjectService_Delete_0(ctx context.Context, marshaler runtime.Marshaler, client ProjectServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ProjectQuery
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "name")
	}

	protoReq.Name, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "name", err)
	}

	msg, err := client.Delete(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_ProjectService_ListEvents_0(ctx context.Context, marshaler runtime.Marshaler, client ProjectServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ProjectQuery
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "name")
	}

	protoReq.Name, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "name", err)
	}

	msg, err := client.ListEvents(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterProjectServiceHandlerFromEndpoint is same as RegisterProjectServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterProjectServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterProjectServiceHandler(ctx, mux, conn)
}

// RegisterProjectServiceHandler registers the http handlers for service ProjectService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterProjectServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterProjectServiceHandlerClient(ctx, mux, NewProjectServiceClient(conn))
}

// RegisterProjectServiceHandler registers the http handlers for service ProjectService to "mux".
// The handlers forward requests to the grpc endpoint over the given implementation of "ProjectServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ProjectServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ProjectServiceClient" to call the correct interceptors.
func RegisterProjectServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ProjectServiceClient) error {

	mux.Handle("POST", pattern_ProjectService_Create_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProjectService_Create_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProjectService_Create_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ProjectService_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProjectService_List_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProjectService_List_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ProjectService_Get_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProjectService_Get_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProjectService_Get_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_ProjectService_Update_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProjectService_Update_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProjectService_Update_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_ProjectService_Delete_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProjectService_Delete_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProjectService_Delete_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ProjectService_ListEvents_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ProjectService_ListEvents_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ProjectService_ListEvents_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ProjectService_Create_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "v1", "projects"}, ""))

	pattern_ProjectService_List_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"api", "v1", "projects"}, ""))

	pattern_ProjectService_Get_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"api", "v1", "projects", "name"}, ""))

	pattern_ProjectService_Update_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"api", "v1", "projects", "project.metadata.name"}, ""))

	pattern_ProjectService_Delete_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"api", "v1", "projects", "name"}, ""))

	pattern_ProjectService_ListEvents_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3, 2, 4}, []string{"api", "v1", "projects", "name", "events"}, ""))
)

var (
	forward_ProjectService_Create_0 = runtime.ForwardResponseMessage

	forward_ProjectService_List_0 = runtime.ForwardResponseMessage

	forward_ProjectService_Get_0 = runtime.ForwardResponseMessage

	forward_ProjectService_Update_0 = runtime.ForwardResponseMessage

	forward_ProjectService_Delete_0 = runtime.ForwardResponseMessage

	forward_ProjectService_ListEvents_0 = runtime.ForwardResponseMessage
)
