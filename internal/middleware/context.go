package middleware

import (
	"context"

	"ledger-service/internal/pkg/app"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Context struct{}

// Unary interceptor
func (a *Context) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		var err error

		ctx, err = a.context(ctx)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

// Stream interceptor
func (a *Context) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		_, err := a.context(stream.Context())
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func (a Context) context(ctx context.Context) (context.Context, error) {
	// Get token from incoming metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	userId := md["user_id"]
	if len(userId) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "user_id metadata is not provided")
	}

	universityId := md["university_id"]
	if len(universityId) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "university_id metadata is not provided")
	}

	programStudiId := md["program_studi_id"]
	if len(programStudiId) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "program_studi_id metadata is not provided")
	}

	ctx = context.WithValue(ctx, app.Ctx("user_id"), userId[0])
	ctx = context.WithValue(ctx, app.Ctx("university_id"), universityId[0])
	ctx = context.WithValue(ctx, app.Ctx("program_studi_id"), programStudiId[0])

	// Set token to outgoing metadata
	mdOutgoing := metadata.New(map[string]string{
		"user_id":          userId[0],
		"university_id":    universityId[0],
		"program_studi_id": programStudiId[0],
	})

	ctx = metadata.NewOutgoingContext(ctx, mdOutgoing)
	return ctx, nil
}
