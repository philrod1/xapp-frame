// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/philrod1/xapp-frame/pkg/clientmodel"
)

// GetAllSubscriptionsReader is a Reader for the GetAllSubscriptions structure.
type GetAllSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAllSubscriptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllSubscriptionsOK creates a GetAllSubscriptionsOK with default headers values
func NewGetAllSubscriptionsOK() *GetAllSubscriptionsOK {
	return &GetAllSubscriptionsOK{}
}

/*GetAllSubscriptionsOK handles this case with default header values.

successful query of subscriptions
*/
type GetAllSubscriptionsOK struct {
	Payload clientmodel.SubscriptionList
}

func (o *GetAllSubscriptionsOK) Error() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getAllSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *GetAllSubscriptionsOK) GetPayload() clientmodel.SubscriptionList {
	return o.Payload
}

func (o *GetAllSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllSubscriptionsInternalServerError creates a GetAllSubscriptionsInternalServerError with default headers values
func NewGetAllSubscriptionsInternalServerError() *GetAllSubscriptionsInternalServerError {
	return &GetAllSubscriptionsInternalServerError{}
}

/*GetAllSubscriptionsInternalServerError handles this case with default header values.

Internal error
*/
type GetAllSubscriptionsInternalServerError struct {
}

func (o *GetAllSubscriptionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /subscriptions][%d] getAllSubscriptionsInternalServerError ", 500)
}

func (o *GetAllSubscriptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
