// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideowebrtcstorage

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	//
	// You do not have required permissions to perform this operation
	//
	// The AccessDeniedException can be thrown for operation access as well as tokens
	// or resource access
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeClientLimitExceededException for service response error code
	// "ClientLimitExceededException".
	//
	// Kinesis Video Streams has throttled the request because you have exceeded
	// the limit of allowed client calls. Try making the call later.
	ErrCodeClientLimitExceededException = "ClientLimitExceededException"

	// ErrCodeInvalidArgumentException for service response error code
	// "InvalidArgumentException".
	//
	//
	// The value for this input parameter is invalid.
	//
	// Additional details may notbe returned.
	ErrCodeInvalidArgumentException = "InvalidArgumentException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	//
	// The specified resource is not found
	//
	// You have not specified a channel in this API call.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":        newErrorAccessDeniedException,
	"ClientLimitExceededException": newErrorClientLimitExceededException,
	"InvalidArgumentException":     newErrorInvalidArgumentException,
	"ResourceNotFoundException":    newErrorResourceNotFoundException,
}
