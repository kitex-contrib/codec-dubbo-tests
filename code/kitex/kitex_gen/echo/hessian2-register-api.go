package echo

import (
	"fmt"

	"github.com/kitex-contrib/codec-dubbo/pkg/hessian2"
	codec "github.com/kitex-contrib/codec-dubbo/pkg/iface"
	"github.com/pkg/errors"
)

var objectsApi = []interface{}{
	&EchoRequest{},
	&EchoResponse{},
	&EchoMultiBoolResponse{},
	&EchoMultiByteResponse{},
	&EchoMultiInt16Response{},
	&EchoMultiInt32Response{},
	&EchoMultiInt64Response{},
	&EchoMultiFloatResponse{},
	&EchoMultiDoubleResponse{},
	&EchoMultiStringResponse{},
	&EchoOptionalStructRequest{},
	&EchoOptionalStructResponse{},
	&EchoOptionalMultiBoolRequest{},
	&EchoOptionalMultiInt32Request{},
	&EchoOptionalMultiStringRequest{},
	&EchoOptionalMultiBoolResponse{},
	&EchoOptionalMultiInt32Response{},
	&EchoOptionalMultiStringResponse{},
	&EchoCustomizedException{},
}

func init() {
	hessian2.Register(objectsApi)
}

func GetTestServiceIDLAnnotations() map[string][]string {
	return map[string][]string{
		"JavaClassName": {"org.apache.dubbo.tests.api.UserProvider"},
	}
}

func (p *EchoRequest) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Int32)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoRequest) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Int32)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoRequest) JavaClassName() string {
	return "kitex.echo.EchoRequest"
}

func (p *EchoResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Int32)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Int32)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoResponse) JavaClassName() string {
	return "kitex.echo.EchoResponse"
}

func (p *EchoMultiBoolResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiBoolResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiBoolResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiBoolResponse"
}

func (p *EchoMultiByteResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiByteResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiByteResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiByteResponse"
}

func (p *EchoMultiInt16Response) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiInt16Response) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiInt16Response) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiInt16Response"
}

func (p *EchoMultiInt32Response) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiInt32Response) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiInt32Response) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiInt32Response"
}

func (p *EchoMultiInt64Response) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiInt64Response) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiInt64Response) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiInt64Response"
}

func (p *EchoMultiFloatResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiFloatResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiFloatResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiFloatResponse"
}

func (p *EchoMultiDoubleResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiDoubleResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiDoubleResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiDoubleResponse"
}

func (p *EchoMultiStringResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoMultiStringResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoMultiStringResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoMultiStringResponse"
}

func (p *EchoOptionalStructRequest) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoOptionalStructRequest) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoOptionalStructRequest) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoOptionalStructRequest"
}

func (p *EchoOptionalStructResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Resp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoOptionalStructResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Resp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoOptionalStructResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoOptionalStructResponse"
}

func (p *EchoOptionalMultiBoolRequest) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BasicReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.PackReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoOptionalMultiBoolRequest) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BasicReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.PackReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoOptionalMultiBoolRequest) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoOptionalMultiBoolRequest"
}

func (p *EchoOptionalMultiInt32Request) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BasicReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.PackReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoOptionalMultiInt32Request) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BasicReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.PackReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoOptionalMultiInt32Request) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoOptionalMultiInt32Request"
}

func (p *EchoOptionalMultiStringRequest) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoOptionalMultiStringRequest) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoOptionalMultiStringRequest) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoOptionalMultiStringRequest"
}

func (p *EchoOptionalMultiBoolResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BasicResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.PackResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoOptionalMultiBoolResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BasicResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.PackResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoOptionalMultiBoolResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoOptionalMultiBoolResponse"
}

func (p *EchoOptionalMultiInt32Response) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BasicResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.PackResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoOptionalMultiInt32Response) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BasicResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.PackResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoOptionalMultiInt32Response) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoOptionalMultiInt32Response"
}

func (p *EchoOptionalMultiStringResponse) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListResp)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapResp)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoOptionalMultiStringResponse) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapResp)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoOptionalMultiStringResponse) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoOptionalMultiStringResponse"
}

func (p *EchoCustomizedException) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Exception)
	if err != nil {
		return err
	}

	err = e.Encode(p.CustomizedMessage)
	if err != nil {
		return err
	}

	return nil
}

func (p *EchoCustomizedException) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Exception)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.CustomizedMessage)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *EchoCustomizedException) JavaClassName() string {
	return "org.apache.dubbo.tests.api.EchoCustomizedException"
}

func (p *TestServiceEchoRetByteArgs) Encode(e codec.Encoder) error {
	return nil
}

func (p *TestServiceEchoRetByteArgs) Decode(d codec.Decoder) error {
	return nil
}

func (p *TestServiceEchoRetByteResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoRetByteResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoRetBoolArgs) Encode(e codec.Encoder) error {
	return nil
}

func (p *TestServiceEchoRetBoolArgs) Decode(d codec.Decoder) error {
	return nil
}

func (p *TestServiceEchoRetBoolResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoRetBoolResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoRetInt16Args) Encode(e codec.Encoder) error {
	return nil
}

func (p *TestServiceEchoRetInt16Args) Decode(d codec.Decoder) error {
	return nil
}

func (p *TestServiceEchoRetInt16Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoRetInt16Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoRetInt32Args) Encode(e codec.Encoder) error {
	return nil
}

func (p *TestServiceEchoRetInt32Args) Decode(d codec.Decoder) error {
	return nil
}

func (p *TestServiceEchoRetInt32Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoRetInt32Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoRetInt64Args) Encode(e codec.Encoder) error {
	return nil
}

func (p *TestServiceEchoRetInt64Args) Decode(d codec.Decoder) error {
	return nil
}

func (p *TestServiceEchoRetInt64Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoRetInt64Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoRetFloatArgs) Encode(e codec.Encoder) error {
	return nil
}

func (p *TestServiceEchoRetFloatArgs) Decode(d codec.Decoder) error {
	return nil
}

func (p *TestServiceEchoRetFloatResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoRetFloatResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoRetDoubleArgs) Encode(e codec.Encoder) error {
	return nil
}

func (p *TestServiceEchoRetDoubleArgs) Decode(d codec.Decoder) error {
	return nil
}

func (p *TestServiceEchoRetDoubleResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoRetDoubleResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoRetStringArgs) Encode(e codec.Encoder) error {
	return nil
}

func (p *TestServiceEchoRetStringArgs) Decode(d codec.Decoder) error {
	return nil
}

func (p *TestServiceEchoRetStringResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoRetStringResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoIntArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoIntArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoIntResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoIntResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBoolArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBoolArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBoolResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBoolResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoByteArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoByteArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoByteResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoByteResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoInt16Args) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoInt16Args) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoInt16Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoInt16Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoInt32Args) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoInt32Args) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoInt32Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoInt32Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoInt64Args) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoInt64Args) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoInt64Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoInt64Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoFloatArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoFloatArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoFloatResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoFloatResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoDoubleArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoDoubleArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoDoubleResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoDoubleResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoStringArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoStringArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoStringResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoStringResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBinaryArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBinaryArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBinaryResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBinaryResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBoolListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBoolListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBoolListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBoolListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoByteListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoByteListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoByteListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoByteListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoInt16ListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoInt16ListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoInt16ListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoInt16ListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoInt32ListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoInt32ListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoInt32ListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoInt32ListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoInt64ListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoInt64ListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoInt64ListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoInt64ListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoFloatListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoFloatListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoFloatListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoFloatListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoDoubleListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoDoubleListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoDoubleListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoDoubleListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoStringListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoStringListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoStringListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoStringListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBinaryListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBinaryListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBinaryListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBinaryListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2BoolMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2BoolMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2BoolMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2BoolMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2ByteMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2ByteMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2ByteMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2ByteMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int16MapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int16MapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int16MapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int16MapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int32MapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int32MapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int32MapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int32MapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int64MapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int64MapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int64MapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int64MapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2FloatMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2FloatMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2FloatMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2FloatMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2DoubleMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2DoubleMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2DoubleMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2DoubleMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2StringMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2StringMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2StringMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2StringMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2BinaryMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2BinaryMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2BinaryMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2BinaryMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2BoolListMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2BoolListMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2BoolListMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2BoolListMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2ByteListMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2ByteListMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2ByteListMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2ByteListMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int16ListMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int16ListMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int16ListMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int16ListMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int32ListMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int32ListMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int32ListMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int32ListMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int64ListMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int64ListMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int64ListMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int64ListMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2FloatListMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2FloatListMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2FloatListMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2FloatListMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2DoubleListMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2DoubleListMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2DoubleListMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2DoubleListMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2StringListMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2StringListMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2StringListMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2StringListMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2BinaryListMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2BinaryListMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2BinaryListMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2BinaryListMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBoolArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBoolArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBoolResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBoolResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiByteArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiByteArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiByteResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiByteResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiInt16Args) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiInt16Args) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiInt16Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiInt16Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiInt32Args) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiInt32Args) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiInt32Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiInt32Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiInt64Args) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiInt64Args) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiInt64Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiInt64Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiFloatArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiFloatArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiFloatResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiFloatResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiDoubleArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiDoubleArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiDoubleResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiDoubleResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiStringArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiStringArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiStringResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiStringResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseBoolArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseBoolArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseBoolResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseBoolResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseByteArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseByteArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseByteResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseByteResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseInt16Args) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseInt16Args) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseInt16Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseInt16Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseInt32Args) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseInt32Args) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseInt32Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseInt32Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseInt64Args) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseInt64Args) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseInt64Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseInt64Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseFloatArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseFloatArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseFloatResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseFloatResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseDoubleArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseDoubleArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseDoubleResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseDoubleResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseBoolListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseBoolListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseBoolListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseBoolListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseByteListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseByteListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseByteListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseByteListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseInt16ListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseInt16ListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseInt16ListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseInt16ListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseInt32ListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseInt32ListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseInt32ListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseInt32ListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseInt64ListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseInt64ListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseInt64ListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseInt64ListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseFloatListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseFloatListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseFloatListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseFloatListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseDoubleListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseDoubleListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBaseDoubleListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBaseDoubleListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2BoolBaseMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2BoolBaseMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2BoolBaseMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2BoolBaseMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2ByteBaseMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2ByteBaseMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2ByteBaseMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2ByteBaseMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int16BaseMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int16BaseMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int16BaseMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int16BaseMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int32BaseMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int32BaseMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int32BaseMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int32BaseMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int64BaseMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int64BaseMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2Int64BaseMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2Int64BaseMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2FloatBaseMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2FloatBaseMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2FloatBaseMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2FloatBaseMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2DoubleBaseMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2DoubleBaseMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoBool2DoubleBaseMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoBool2DoubleBaseMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseBoolArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseBoolArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseBoolResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseBoolResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseByteArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseByteArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseByteResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseByteResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseInt16Args) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseInt16Args) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseInt16Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseInt16Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseInt32Args) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseInt32Args) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseInt32Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseInt32Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseInt64Args) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseInt64Args) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseInt64Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseInt64Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseFloatArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseFloatArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseFloatResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseFloatResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseDoubleArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.BaseReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.ListReq)
	if err != nil {
		return err
	}

	err = e.Encode(p.MapReq)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseDoubleArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.BaseReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.ListReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.MapReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMultiBaseDoubleResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMultiBaseDoubleResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMethodAArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMethodAArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMethodAResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMethodAResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMethodBArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMethodBArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMethodBResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMethodBResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMethodCArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMethodCArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMethodCResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMethodCResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMethodDArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req1)
	if err != nil {
		return err
	}

	err = e.Encode(p.Req2)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMethodDArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req1)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req2)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoMethodDResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoMethodDResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalBoolArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalBoolArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalBoolResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalBoolResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalInt32Args) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalInt32Args) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalInt32Result) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalInt32Result) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalStringArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalStringArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalStringResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalStringResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalBoolListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalBoolListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalBoolListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalBoolListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalInt32ListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalInt32ListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalInt32ListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalInt32ListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalStringListArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalStringListArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalStringListResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalStringListResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalBool2BoolMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalBool2BoolMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalBool2BoolMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalBool2BoolMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalBool2Int32MapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalBool2Int32MapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalBool2Int32MapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalBool2Int32MapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalBool2StringMapArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalBool2StringMapArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalBool2StringMapResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalBool2StringMapResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalStructArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalStructArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalStructResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalStructResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiBoolRequestArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiBoolRequestArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiBoolRequestResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiBoolRequestResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiInt32RequestArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiInt32RequestArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiInt32RequestResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiInt32RequestResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiStringRequestArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiStringRequestArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiStringRequestResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiStringRequestResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiBoolResponseArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiBoolResponseArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiBoolResponseResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiBoolResponseResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiInt32ResponseArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiInt32ResponseArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiInt32ResponseResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiInt32ResponseResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiStringResponseArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiStringResponseArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiStringResponseResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoOptionalMultiStringResponseResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoExceptionArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoExceptionArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoExceptionResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoExceptionResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoCustomizedExceptionArgs) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Req)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoCustomizedExceptionArgs) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Req)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}

func (p *TestServiceEchoCustomizedExceptionResult) Encode(e codec.Encoder) error {
	var err error
	err = e.Encode(p.Success)
	if err != nil {
		return err
	}

	return nil
}

func (p *TestServiceEchoCustomizedExceptionResult) Decode(d codec.Decoder) error {
	var (
		err error
		v   interface{}
	)
	v, err = d.Decode()
	if err != nil {
		return err
	}
	err = hessian2.ReflectResponse(v, &p.Success)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("invalid data type: %T", v))
	}

	return nil
}
