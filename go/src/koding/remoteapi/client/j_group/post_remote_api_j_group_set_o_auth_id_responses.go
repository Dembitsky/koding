package j_group

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

// PostRemoteAPIJGroupSetOAuthIDReader is a Reader for the PostRemoteAPIJGroupSetOAuthID structure.
type PostRemoteAPIJGroupSetOAuthIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJGroupSetOAuthIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJGroupSetOAuthIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJGroupSetOAuthIDOK creates a PostRemoteAPIJGroupSetOAuthIDOK with default headers values
func NewPostRemoteAPIJGroupSetOAuthIDOK() *PostRemoteAPIJGroupSetOAuthIDOK {
	return &PostRemoteAPIJGroupSetOAuthIDOK{}
}

/*PostRemoteAPIJGroupSetOAuthIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJGroupSetOAuthIDOK struct {
	Payload PostRemoteAPIJGroupSetOAuthIDOKBody
}

func (o *PostRemoteAPIJGroupSetOAuthIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JGroup.setOAuth/{id}][%d] postRemoteApiJGroupSetOAuthIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJGroupSetOAuthIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJGroupSetOAuthIDOKBody post remote API j group set o auth ID o k body
swagger:model PostRemoteAPIJGroupSetOAuthIDOKBody
*/
type PostRemoteAPIJGroupSetOAuthIDOKBody struct {
	models.JGroup

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJGroupSetOAuthIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJGroupSetOAuthIDOKBodyAO0 models.JGroup
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupSetOAuthIDOKBodyAO0); err != nil {
		return err
	}
	o.JGroup = postRemoteAPIJGroupSetOAuthIDOKBodyAO0

	var postRemoteAPIJGroupSetOAuthIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJGroupSetOAuthIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJGroupSetOAuthIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJGroupSetOAuthIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJGroupSetOAuthIDOKBodyAO0, err := swag.WriteJSON(o.JGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupSetOAuthIDOKBodyAO0)

	postRemoteAPIJGroupSetOAuthIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJGroupSetOAuthIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j group set o auth ID o k body
func (o *PostRemoteAPIJGroupSetOAuthIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JGroup.Validate(formats); err != nil {
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
