package gapi

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func fieldViolation(field string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: err.Error(),
	}
}
func invalidArgumentError(violations []*errdetails.BadRequest_FieldViolation) error {
	badRequest := &errdetails.BadRequest{FieldViolations: violations}
	st := status.New(codes.InvalidArgument, "invalid argument")
	st, err := st.WithDetails(badRequest)
	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error attaching metadata: %v", err)
	}
	return st.Err()
}
