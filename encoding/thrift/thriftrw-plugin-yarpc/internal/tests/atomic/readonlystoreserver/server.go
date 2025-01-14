// Code generated by thriftrw-plugin-yarpc
// @generated

package readonlystoreserver

import (
	context "context"
	stream "go.uber.org/thriftrw/protocol/stream"
	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	atomic "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/atomic"
	baseserviceserver "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/common/baseserviceserver"
	yarpcerrors "go.uber.org/yarpc/yarpcerrors"
)

// Interface is the server-side interface for the ReadOnlyStore service.
type Interface interface {
	baseserviceserver.Interface

	Integer(
		ctx context.Context,
		Key *string,
	) (int64, error)
}

// New prepares an implementation of the ReadOnlyStore service for
// registration.
//
// 	handler := ReadOnlyStoreHandler{}
// 	dispatcher.Register(readonlystoreserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "ReadOnlyStore",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "integer",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Unary,
					Unary:  thrift.UnaryHandler(h.Integer),
					NoWire: integer_NoWireHandler{impl},
				},
				Signature:    "Integer(Key *string) (int64)",
				ThriftModule: atomic.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 1)

	procedures = append(
		procedures,
		baseserviceserver.New(
			impl,
			append(
				opts,
				thrift.Named("ReadOnlyStore"),
			)...,
		)...,
	)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

type yarpcErrorNamer interface{ YARPCErrorName() string }

type yarpcErrorCoder interface{ YARPCErrorCode() *yarpcerrors.Code }

func (h handler) Integer(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args atomic.ReadOnlyStore_Integer_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'ReadOnlyStore' procedure 'Integer': %w", err)
	}

	success, appErr := h.impl.Integer(ctx, args.Key)

	hadError := appErr != nil
	result, err := atomic.ReadOnlyStore_Integer_Helper.WrapResponse(success, appErr)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}

	return response, err
}

type integer_NoWireHandler struct{ impl Interface }

func (h integer_NoWireHandler) HandleNoWire(ctx context.Context, nwc *thrift.NoWireCall) (thrift.NoWireResponse, error) {
	var (
		args atomic.ReadOnlyStore_Integer_Args
		rw   stream.ResponseWriter
		err  error
	)

	rw, err = nwc.RequestReader.ReadRequest(ctx, nwc.EnvelopeType, nwc.Reader, &args)
	if err != nil {
		return thrift.NoWireResponse{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode (via no wire) Thrift request for service 'ReadOnlyStore' procedure 'Integer': %w", err)
	}

	success, appErr := h.impl.Integer(ctx, args.Key)

	hadError := appErr != nil
	result, err := atomic.ReadOnlyStore_Integer_Helper.WrapResponse(success, appErr)
	response := thrift.NoWireResponse{ResponseWriter: rw}
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}
	return response, err

}
