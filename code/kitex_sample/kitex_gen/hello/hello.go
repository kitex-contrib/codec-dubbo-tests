// Code generated by thriftgo (0.3.10). DO NOT EDIT.

package hello

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type GreetRequest struct {
	Req string `thrift:"req,1,required" frugal:"1,required,string" json:"req"`
}

func NewGreetRequest() *GreetRequest {
	return &GreetRequest{}
}

func (p *GreetRequest) InitDefault() {
	*p = GreetRequest{}
}

func (p *GreetRequest) GetReq() (v string) {
	return p.Req
}
func (p *GreetRequest) SetReq(val string) {
	p.Req = val
}

var fieldIDToName_GreetRequest = map[int16]string{
	1: "req",
}

func (p *GreetRequest) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetReq bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetReq = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetReq {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GreetRequest[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_GreetRequest[fieldId]))
}

func (p *GreetRequest) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Req = v
	}
	return nil
}

func (p *GreetRequest) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GreetRequest"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GreetRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Req); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GreetRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GreetRequest(%+v)", *p)

}

func (p *GreetRequest) DeepEqual(ano *GreetRequest) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *GreetRequest) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Req, src) != 0 {
		return false
	}
	return true
}

type GreetResponse struct {
	Resp string `thrift:"resp,1,required" frugal:"1,required,string" json:"resp"`
}

func NewGreetResponse() *GreetResponse {
	return &GreetResponse{}
}

func (p *GreetResponse) InitDefault() {
	*p = GreetResponse{}
}

func (p *GreetResponse) GetResp() (v string) {
	return p.Resp
}
func (p *GreetResponse) SetResp(val string) {
	p.Resp = val
}

var fieldIDToName_GreetResponse = map[int16]string{
	1: "resp",
}

func (p *GreetResponse) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16
	var issetResp bool = false

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
				issetResp = true
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	if !issetResp {
		fieldId = 1
		goto RequiredFieldNotSetError
	}
	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GreetResponse[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
RequiredFieldNotSetError:
	return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("required field %s is not set", fieldIDToName_GreetResponse[fieldId]))
}

func (p *GreetResponse) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Resp = v
	}
	return nil
}

func (p *GreetResponse) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GreetResponse"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GreetResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("resp", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Resp); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GreetResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GreetResponse(%+v)", *p)

}

func (p *GreetResponse) DeepEqual(ano *GreetResponse) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Resp) {
		return false
	}
	return true
}

func (p *GreetResponse) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Resp, src) != 0 {
		return false
	}
	return true
}

type GreetService interface {
	Greet(ctx context.Context, req string) (r string, err error)

	GreetWithStruct(ctx context.Context, req *GreetRequest) (r *GreetResponse, err error)
}

type GreetServiceClient struct {
	c thrift.TClient
}

func NewGreetServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *GreetServiceClient {
	return &GreetServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewGreetServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *GreetServiceClient {
	return &GreetServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewGreetServiceClient(c thrift.TClient) *GreetServiceClient {
	return &GreetServiceClient{
		c: c,
	}
}

func (p *GreetServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *GreetServiceClient) Greet(ctx context.Context, req string) (r string, err error) {
	var _args GreetServiceGreetArgs
	_args.Req = req
	var _result GreetServiceGreetResult
	if err = p.Client_().Call(ctx, "Greet", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
func (p *GreetServiceClient) GreetWithStruct(ctx context.Context, req *GreetRequest) (r *GreetResponse, err error) {
	var _args GreetServiceGreetWithStructArgs
	_args.Req = req
	var _result GreetServiceGreetWithStructResult
	if err = p.Client_().Call(ctx, "GreetWithStruct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type GreetServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      GreetService
}

func (p *GreetServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *GreetServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *GreetServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewGreetServiceProcessor(handler GreetService) *GreetServiceProcessor {
	self := &GreetServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("Greet", &greetServiceProcessorGreet{handler: handler})
	self.AddToProcessorMap("GreetWithStruct", &greetServiceProcessorGreetWithStruct{handler: handler})
	return self
}
func (p *GreetServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type greetServiceProcessorGreet struct {
	handler GreetService
}

func (p *greetServiceProcessorGreet) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GreetServiceGreetArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Greet", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := GreetServiceGreetResult{}
	var retval string
	if retval, err2 = p.handler.Greet(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Greet: "+err2.Error())
		oprot.WriteMessageBegin("Greet", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("Greet", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type greetServiceProcessorGreetWithStruct struct {
	handler GreetService
}

func (p *greetServiceProcessorGreetWithStruct) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := GreetServiceGreetWithStructArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GreetWithStruct", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := GreetServiceGreetWithStructResult{}
	var retval *GreetResponse
	if retval, err2 = p.handler.GreetWithStruct(ctx, args.Req); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GreetWithStruct: "+err2.Error())
		oprot.WriteMessageBegin("GreetWithStruct", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GreetWithStruct", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type GreetServiceGreetArgs struct {
	Req string `thrift:"req,1" frugal:"1,default,string" json:"req"`
}

func NewGreetServiceGreetArgs() *GreetServiceGreetArgs {
	return &GreetServiceGreetArgs{}
}

func (p *GreetServiceGreetArgs) InitDefault() {
	*p = GreetServiceGreetArgs{}
}

func (p *GreetServiceGreetArgs) GetReq() (v string) {
	return p.Req
}
func (p *GreetServiceGreetArgs) SetReq(val string) {
	p.Req = val
}

var fieldIDToName_GreetServiceGreetArgs = map[int16]string{
	1: "req",
}

func (p *GreetServiceGreetArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GreetServiceGreetArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GreetServiceGreetArgs) ReadField1(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Req = v
	}
	return nil
}

func (p *GreetServiceGreetArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Greet_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GreetServiceGreetArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRING, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Req); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GreetServiceGreetArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GreetServiceGreetArgs(%+v)", *p)

}

func (p *GreetServiceGreetArgs) DeepEqual(ano *GreetServiceGreetArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *GreetServiceGreetArgs) Field1DeepEqual(src string) bool {

	if strings.Compare(p.Req, src) != 0 {
		return false
	}
	return true
}

type GreetServiceGreetResult struct {
	Success *string `thrift:"success,0,optional" frugal:"0,optional,string" json:"success,omitempty"`
}

func NewGreetServiceGreetResult() *GreetServiceGreetResult {
	return &GreetServiceGreetResult{}
}

func (p *GreetServiceGreetResult) InitDefault() {
	*p = GreetServiceGreetResult{}
}

var GreetServiceGreetResult_Success_DEFAULT string

func (p *GreetServiceGreetResult) GetSuccess() (v string) {
	if !p.IsSetSuccess() {
		return GreetServiceGreetResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *GreetServiceGreetResult) SetSuccess(x interface{}) {
	p.Success = x.(*string)
}

var fieldIDToName_GreetServiceGreetResult = map[int16]string{
	0: "success",
}

func (p *GreetServiceGreetResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GreetServiceGreetResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GreetServiceGreetResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GreetServiceGreetResult) ReadField0(iprot thrift.TProtocol) error {

	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Success = &v
	}
	return nil
}

func (p *GreetServiceGreetResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("Greet_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GreetServiceGreetResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteString(*p.Success); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *GreetServiceGreetResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GreetServiceGreetResult(%+v)", *p)

}

func (p *GreetServiceGreetResult) DeepEqual(ano *GreetServiceGreetResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *GreetServiceGreetResult) Field0DeepEqual(src *string) bool {

	if p.Success == src {
		return true
	} else if p.Success == nil || src == nil {
		return false
	}
	if strings.Compare(*p.Success, *src) != 0 {
		return false
	}
	return true
}

type GreetServiceGreetWithStructArgs struct {
	Req *GreetRequest `thrift:"req,1" frugal:"1,default,GreetRequest" json:"req"`
}

func NewGreetServiceGreetWithStructArgs() *GreetServiceGreetWithStructArgs {
	return &GreetServiceGreetWithStructArgs{}
}

func (p *GreetServiceGreetWithStructArgs) InitDefault() {
	*p = GreetServiceGreetWithStructArgs{}
}

var GreetServiceGreetWithStructArgs_Req_DEFAULT *GreetRequest

func (p *GreetServiceGreetWithStructArgs) GetReq() (v *GreetRequest) {
	if !p.IsSetReq() {
		return GreetServiceGreetWithStructArgs_Req_DEFAULT
	}
	return p.Req
}
func (p *GreetServiceGreetWithStructArgs) SetReq(val *GreetRequest) {
	p.Req = val
}

var fieldIDToName_GreetServiceGreetWithStructArgs = map[int16]string{
	1: "req",
}

func (p *GreetServiceGreetWithStructArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GreetServiceGreetWithStructArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GreetServiceGreetWithStructArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GreetServiceGreetWithStructArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Req = NewGreetRequest()
	if err := p.Req.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *GreetServiceGreetWithStructArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GreetWithStruct_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GreetServiceGreetWithStructArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("req", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.Req.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *GreetServiceGreetWithStructArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GreetServiceGreetWithStructArgs(%+v)", *p)

}

func (p *GreetServiceGreetWithStructArgs) DeepEqual(ano *GreetServiceGreetWithStructArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Req) {
		return false
	}
	return true
}

func (p *GreetServiceGreetWithStructArgs) Field1DeepEqual(src *GreetRequest) bool {

	if !p.Req.DeepEqual(src) {
		return false
	}
	return true
}

type GreetServiceGreetWithStructResult struct {
	Success *GreetResponse `thrift:"success,0,optional" frugal:"0,optional,GreetResponse" json:"success,omitempty"`
}

func NewGreetServiceGreetWithStructResult() *GreetServiceGreetWithStructResult {
	return &GreetServiceGreetWithStructResult{}
}

func (p *GreetServiceGreetWithStructResult) InitDefault() {
	*p = GreetServiceGreetWithStructResult{}
}

var GreetServiceGreetWithStructResult_Success_DEFAULT *GreetResponse

func (p *GreetServiceGreetWithStructResult) GetSuccess() (v *GreetResponse) {
	if !p.IsSetSuccess() {
		return GreetServiceGreetWithStructResult_Success_DEFAULT
	}
	return p.Success
}
func (p *GreetServiceGreetWithStructResult) SetSuccess(x interface{}) {
	p.Success = x.(*GreetResponse)
}

var fieldIDToName_GreetServiceGreetWithStructResult = map[int16]string{
	0: "success",
}

func (p *GreetServiceGreetWithStructResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GreetServiceGreetWithStructResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}
		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_GreetServiceGreetWithStructResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *GreetServiceGreetWithStructResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = NewGreetResponse()
	if err := p.Success.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *GreetServiceGreetWithStructResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("GreetWithStruct_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}
	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *GreetServiceGreetWithStructResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := p.Success.Write(oprot); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *GreetServiceGreetWithStructResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GreetServiceGreetWithStructResult(%+v)", *p)

}

func (p *GreetServiceGreetWithStructResult) DeepEqual(ano *GreetServiceGreetWithStructResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *GreetServiceGreetWithStructResult) Field0DeepEqual(src *GreetResponse) bool {

	if !p.Success.DeepEqual(src) {
		return false
	}
	return true
}
