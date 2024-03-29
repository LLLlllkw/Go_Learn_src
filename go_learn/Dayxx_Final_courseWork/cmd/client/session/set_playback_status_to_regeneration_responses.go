// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"courseWork/cmd/models"
)

// SetPlaybackStatusToRegenerationReader is a Reader for the SetPlaybackStatusToRegeneration structure.
type SetPlaybackStatusToRegenerationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetPlaybackStatusToRegenerationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSetPlaybackStatusToRegenerationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSetPlaybackStatusToRegenerationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetPlaybackStatusToRegenerationCreated creates a SetPlaybackStatusToRegenerationCreated with default headers values
func NewSetPlaybackStatusToRegenerationCreated() *SetPlaybackStatusToRegenerationCreated {
	return &SetPlaybackStatusToRegenerationCreated{}
}

/* SetPlaybackStatusToRegenerationCreated describes a response with status code 201, with default header values.

请求成功
*/
type SetPlaybackStatusToRegenerationCreated struct {
	Payload *SetPlaybackStatusToRegenerationCreatedBody
}

func (o *SetPlaybackStatusToRegenerationCreated) Error() string {
	return fmt.Sprintf("[PUT /{sessionId}/playbackStatus/regeneration][%d] setPlaybackStatusToRegenerationCreated  %+v", 201, o.Payload)
}
func (o *SetPlaybackStatusToRegenerationCreated) GetPayload() *SetPlaybackStatusToRegenerationCreatedBody {
	return o.Payload
}

func (o *SetPlaybackStatusToRegenerationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SetPlaybackStatusToRegenerationCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetPlaybackStatusToRegenerationInternalServerError creates a SetPlaybackStatusToRegenerationInternalServerError with default headers values
func NewSetPlaybackStatusToRegenerationInternalServerError() *SetPlaybackStatusToRegenerationInternalServerError {
	return &SetPlaybackStatusToRegenerationInternalServerError{}
}

/* SetPlaybackStatusToRegenerationInternalServerError describes a response with status code 500, with default header values.

服务器内部异常
*/
type SetPlaybackStatusToRegenerationInternalServerError struct {
	Payload *models.Error
}

func (o *SetPlaybackStatusToRegenerationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /{sessionId}/playbackStatus/regeneration][%d] setPlaybackStatusToRegenerationInternalServerError  %+v", 500, o.Payload)
}
func (o *SetPlaybackStatusToRegenerationInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetPlaybackStatusToRegenerationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SetPlaybackStatusToRegenerationCreatedBody set playback status to regeneration created body
swagger:model SetPlaybackStatusToRegenerationCreatedBody
*/
type SetPlaybackStatusToRegenerationCreatedBody struct {

	// msg
	// Example: 重新生成成功
	Msg string `json:"msg,omitempty"`

	// 结果情况, success 代表成功, 其他情况 代表失败
	// Example: success
	// Enum: [success sessionNoFound repeatSubmit generating]
	Result string `json:"result,omitempty"`
}

// Validate validates this set playback status to regeneration created body
func (o *SetPlaybackStatusToRegenerationCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var setPlaybackStatusToRegenerationCreatedBodyTypeResultPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["success","sessionNoFound","repeatSubmit","generating"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		setPlaybackStatusToRegenerationCreatedBodyTypeResultPropEnum = append(setPlaybackStatusToRegenerationCreatedBodyTypeResultPropEnum, v)
	}
}

const (

	// SetPlaybackStatusToRegenerationCreatedBodyResultSuccess captures enum value "success"
	SetPlaybackStatusToRegenerationCreatedBodyResultSuccess string = "success"

	// SetPlaybackStatusToRegenerationCreatedBodyResultSessionNoFound captures enum value "sessionNoFound"
	SetPlaybackStatusToRegenerationCreatedBodyResultSessionNoFound string = "sessionNoFound"

	// SetPlaybackStatusToRegenerationCreatedBodyResultRepeatSubmit captures enum value "repeatSubmit"
	SetPlaybackStatusToRegenerationCreatedBodyResultRepeatSubmit string = "repeatSubmit"

	// SetPlaybackStatusToRegenerationCreatedBodyResultGenerating captures enum value "generating"
	SetPlaybackStatusToRegenerationCreatedBodyResultGenerating string = "generating"
)

// prop value enum
func (o *SetPlaybackStatusToRegenerationCreatedBody) validateResultEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, setPlaybackStatusToRegenerationCreatedBodyTypeResultPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *SetPlaybackStatusToRegenerationCreatedBody) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(o.Result) { // not required
		return nil
	}

	// value enum
	if err := o.validateResultEnum("setPlaybackStatusToRegenerationCreated"+"."+"result", "body", o.Result); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this set playback status to regeneration created body based on context it is used
func (o *SetPlaybackStatusToRegenerationCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SetPlaybackStatusToRegenerationCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SetPlaybackStatusToRegenerationCreatedBody) UnmarshalBinary(b []byte) error {
	var res SetPlaybackStatusToRegenerationCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
