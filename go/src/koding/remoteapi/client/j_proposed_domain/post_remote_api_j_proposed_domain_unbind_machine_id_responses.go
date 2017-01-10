package j_proposed_domain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJProposedDomainUnbindMachineIDReader is a Reader for the PostRemoteAPIJProposedDomainUnbindMachineID structure.
type PostRemoteAPIJProposedDomainUnbindMachineIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJProposedDomainUnbindMachineIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJProposedDomainUnbindMachineIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJProposedDomainUnbindMachineIDOK creates a PostRemoteAPIJProposedDomainUnbindMachineIDOK with default headers values
func NewPostRemoteAPIJProposedDomainUnbindMachineIDOK() *PostRemoteAPIJProposedDomainUnbindMachineIDOK {
	return &PostRemoteAPIJProposedDomainUnbindMachineIDOK{}
}

/*PostRemoteAPIJProposedDomainUnbindMachineIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJProposedDomainUnbindMachineIDOK struct {
	Payload PostRemoteAPIJProposedDomainUnbindMachineIDOKBody
}

func (o *PostRemoteAPIJProposedDomainUnbindMachineIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProposedDomain.unbindMachine/{id}][%d] postRemoteApiJProposedDomainUnbindMachineIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJProposedDomainUnbindMachineIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJProposedDomainUnbindMachineIDOKBody post remote API j proposed domain unbind machine ID o k body
swagger:model PostRemoteAPIJProposedDomainUnbindMachineIDOKBody
*/
type PostRemoteAPIJProposedDomainUnbindMachineIDOKBody struct {
	models.JProposedDomain

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJProposedDomainUnbindMachineIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJProposedDomainUnbindMachineIDOKBodyAO0 models.JProposedDomain
	if err := swag.ReadJSON(raw, &postRemoteAPIJProposedDomainUnbindMachineIDOKBodyAO0); err != nil {
		return err
	}
	o.JProposedDomain = postRemoteAPIJProposedDomainUnbindMachineIDOKBodyAO0

	var postRemoteAPIJProposedDomainUnbindMachineIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJProposedDomainUnbindMachineIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJProposedDomainUnbindMachineIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJProposedDomainUnbindMachineIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJProposedDomainUnbindMachineIDOKBodyAO0, err := swag.WriteJSON(o.JProposedDomain)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJProposedDomainUnbindMachineIDOKBodyAO0)

	postRemoteAPIJProposedDomainUnbindMachineIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJProposedDomainUnbindMachineIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j proposed domain unbind machine ID o k body
func (o *PostRemoteAPIJProposedDomainUnbindMachineIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JProposedDomain.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
