// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package apm

import (
	"bytes"
	"context"
	"reflect"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

// Attributes:
//  - TraceId
//  - Name
//  - ID
//  - ParentId
//  - Annotations
//  - BinaryAnnotations
//  - Debug
//  - Timestamp
//  - Duration
type TSpan struct {
  TraceId int64 `thrift:"traceId,1" db:"traceId" json:"traceId"`
  // unused field # 2
  Name string `thrift:"name,3" db:"name" json:"name"`
  ID int64 `thrift:"id,4" db:"id" json:"id"`
  ParentId *int64 `thrift:"parentId,5" db:"parentId" json:"parentId,omitempty"`
  Annotations []*TAnnotation `thrift:"annotations,6" db:"annotations" json:"annotations"`
  // unused field # 7
  BinaryAnnotations []*TBinaryAnnotation `thrift:"binaryAnnotations,8" db:"binaryAnnotations" json:"binaryAnnotations"`
  Debug *bool `thrift:"debug,9" db:"debug" json:"debug,omitempty"`
  Timestamp *int64 `thrift:"timestamp,10" db:"timestamp" json:"timestamp,omitempty"`
  Duration *int64 `thrift:"duration,11" db:"duration" json:"duration,omitempty"`
}


func NewTSpan() *TSpan {
  return &TSpan{}
}


func (p *TSpan) GetTraceId() int64 {
  return p.TraceId
}

func (p *TSpan) GetName() string {
  return p.Name
}

func (p *TSpan) GetID() int64 {
  return p.ID
}
var TSpan_ParentId_DEFAULT int64
func (p *TSpan) GetParentId() int64 {
  if !p.IsSetParentId() {
    return TSpan_ParentId_DEFAULT
  }
return *p.ParentId
}

func (p *TSpan) GetAnnotations() []*TAnnotation {
  return p.Annotations
}

func (p *TSpan) GetBinaryAnnotations() []*TBinaryAnnotation {
  return p.BinaryAnnotations
}
var TSpan_Debug_DEFAULT bool
func (p *TSpan) GetDebug() bool {
  if !p.IsSetDebug() {
    return TSpan_Debug_DEFAULT
  }
return *p.Debug
}
var TSpan_Timestamp_DEFAULT int64
func (p *TSpan) GetTimestamp() int64 {
  if !p.IsSetTimestamp() {
    return TSpan_Timestamp_DEFAULT
  }
return *p.Timestamp
}
var TSpan_Duration_DEFAULT int64
func (p *TSpan) GetDuration() int64 {
  if !p.IsSetDuration() {
    return TSpan_Duration_DEFAULT
  }
return *p.Duration
}
func (p *TSpan) IsSetParentId() bool {
  return p.ParentId != nil
}

func (p *TSpan) IsSetDebug() bool {
  return p.Debug != nil
}

func (p *TSpan) IsSetTimestamp() bool {
  return p.Timestamp != nil
}

func (p *TSpan) IsSetDuration() bool {
  return p.Duration != nil
}

func (p *TSpan) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField4(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 5:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField5(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 6:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField6(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 8:
      if fieldTypeId == thrift.LIST {
        if err := p.ReadField8(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 9:
      if fieldTypeId == thrift.BOOL {
        if err := p.ReadField9(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 10:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField10(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 11:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField11(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TSpan)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.TraceId = v
}
  return nil
}

func (p *TSpan)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Name = v
}
  return nil
}

func (p *TSpan)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 4: ", err)
} else {
  p.ID = v
}
  return nil
}

func (p *TSpan)  ReadField5(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 5: ", err)
} else {
  p.ParentId = &v
}
  return nil
}

func (p *TSpan)  ReadField6(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*TAnnotation, 0, size)
  p.Annotations =  tSlice
  for i := 0; i < size; i ++ {
    _elem0 := &TAnnotation{}
    if err := _elem0.Read(iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
    }
    p.Annotations = append(p.Annotations, _elem0)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *TSpan)  ReadField8(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return thrift.PrependError("error reading list begin: ", err)
  }
  tSlice := make([]*TBinaryAnnotation, 0, size)
  p.BinaryAnnotations =  tSlice
  for i := 0; i < size; i ++ {
    _elem1 := &TBinaryAnnotation{}
    if err := _elem1.Read(iprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem1), err)
    }
    p.BinaryAnnotations = append(p.BinaryAnnotations, _elem1)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return thrift.PrependError("error reading list end: ", err)
  }
  return nil
}

func (p *TSpan)  ReadField9(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return thrift.PrependError("error reading field 9: ", err)
} else {
  p.Debug = &v
}
  return nil
}

func (p *TSpan)  ReadField10(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 10: ", err)
} else {
  p.Timestamp = &v
}
  return nil
}

func (p *TSpan)  ReadField11(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 11: ", err)
} else {
  p.Duration = &v
}
  return nil
}

func (p *TSpan) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("TSpan"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
    if err := p.writeField5(oprot); err != nil { return err }
    if err := p.writeField6(oprot); err != nil { return err }
    if err := p.writeField8(oprot); err != nil { return err }
    if err := p.writeField9(oprot); err != nil { return err }
    if err := p.writeField10(oprot); err != nil { return err }
    if err := p.writeField11(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TSpan) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("traceId", thrift.I64, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:traceId: ", p), err) }
  if err := oprot.WriteI64(int64(p.TraceId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.traceId (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:traceId: ", p), err) }
  return err
}

func (p *TSpan) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("name", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:name: ", p), err) }
  if err := oprot.WriteString(string(p.Name)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.name (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:name: ", p), err) }
  return err
}

func (p *TSpan) writeField4(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("id", thrift.I64, 4); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:id: ", p), err) }
  if err := oprot.WriteI64(int64(p.ID)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.id (4) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 4:id: ", p), err) }
  return err
}

func (p *TSpan) writeField5(oprot thrift.TProtocol) (err error) {
  if p.IsSetParentId() {
    if err := oprot.WriteFieldBegin("parentId", thrift.I64, 5); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:parentId: ", p), err) }
    if err := oprot.WriteI64(int64(*p.ParentId)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.parentId (5) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 5:parentId: ", p), err) }
  }
  return err
}

func (p *TSpan) writeField6(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("annotations", thrift.LIST, 6); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:annotations: ", p), err) }
  if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Annotations)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.Annotations {
    if err := v.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
    }
  }
  if err := oprot.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 6:annotations: ", p), err) }
  return err
}

func (p *TSpan) writeField8(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("binaryAnnotations", thrift.LIST, 8); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 8:binaryAnnotations: ", p), err) }
  if err := oprot.WriteListBegin(thrift.STRUCT, len(p.BinaryAnnotations)); err != nil {
    return thrift.PrependError("error writing list begin: ", err)
  }
  for _, v := range p.BinaryAnnotations {
    if err := v.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
    }
  }
  if err := oprot.WriteListEnd(); err != nil {
    return thrift.PrependError("error writing list end: ", err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 8:binaryAnnotations: ", p), err) }
  return err
}

func (p *TSpan) writeField9(oprot thrift.TProtocol) (err error) {
  if p.IsSetDebug() {
    if err := oprot.WriteFieldBegin("debug", thrift.BOOL, 9); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 9:debug: ", p), err) }
    if err := oprot.WriteBool(bool(*p.Debug)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.debug (9) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 9:debug: ", p), err) }
  }
  return err
}

func (p *TSpan) writeField10(oprot thrift.TProtocol) (err error) {
  if p.IsSetTimestamp() {
    if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 10); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:timestamp: ", p), err) }
    if err := oprot.WriteI64(int64(*p.Timestamp)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.timestamp (10) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 10:timestamp: ", p), err) }
  }
  return err
}

func (p *TSpan) writeField11(oprot thrift.TProtocol) (err error) {
  if p.IsSetDuration() {
    if err := oprot.WriteFieldBegin("duration", thrift.I64, 11); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 11:duration: ", p), err) }
    if err := oprot.WriteI64(int64(*p.Duration)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.duration (11) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 11:duration: ", p), err) }
  }
  return err
}

func (p *TSpan) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TSpan(%+v)", *p)
}

// Attributes:
//  - Timestamp
//  - Value
//  - Endpoint
type TAnnotation struct {
  Timestamp int64 `thrift:"timestamp,1" db:"timestamp" json:"timestamp"`
  Value *string `thrift:"value,2" db:"value" json:"value,omitempty"`
  Endpoint *TEndpoint `thrift:"endpoint,3" db:"endpoint" json:"endpoint,omitempty"`
}

func NewTAnnotation() *TAnnotation {
  return &TAnnotation{}
}


func (p *TAnnotation) GetTimestamp() int64 {
  return p.Timestamp
}
var TAnnotation_Value_DEFAULT string
func (p *TAnnotation) GetValue() string {
  if !p.IsSetValue() {
    return TAnnotation_Value_DEFAULT
  }
return *p.Value
}
var TAnnotation_Endpoint_DEFAULT *TEndpoint
func (p *TAnnotation) GetEndpoint() *TEndpoint {
  if !p.IsSetEndpoint() {
    return TAnnotation_Endpoint_DEFAULT
  }
return p.Endpoint
}
func (p *TAnnotation) IsSetValue() bool {
  return p.Value != nil
}

func (p *TAnnotation) IsSetEndpoint() bool {
  return p.Endpoint != nil
}

func (p *TAnnotation) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I64 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TAnnotation)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Timestamp = v
}
  return nil
}

func (p *TAnnotation)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Value = &v
}
  return nil
}

func (p *TAnnotation)  ReadField3(iprot thrift.TProtocol) error {
  p.Endpoint = &TEndpoint{}
  if err := p.Endpoint.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Endpoint), err)
  }
  return nil
}

func (p *TAnnotation) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("TAnnotation"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TAnnotation) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("timestamp", thrift.I64, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:timestamp: ", p), err) }
  if err := oprot.WriteI64(int64(p.Timestamp)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.timestamp (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:timestamp: ", p), err) }
  return err
}

func (p *TAnnotation) writeField2(oprot thrift.TProtocol) (err error) {
  if p.IsSetValue() {
    if err := oprot.WriteFieldBegin("value", thrift.STRING, 2); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:value: ", p), err) }
    if err := oprot.WriteString(string(*p.Value)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.value (2) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 2:value: ", p), err) }
  }
  return err
}

func (p *TAnnotation) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetEndpoint() {
    if err := oprot.WriteFieldBegin("endpoint", thrift.STRUCT, 3); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:endpoint: ", p), err) }
    if err := p.Endpoint.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Endpoint), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 3:endpoint: ", p), err) }
  }
  return err
}

func (p *TAnnotation) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TAnnotation(%+v)", *p)
}

// Attributes:
//  - Ipv4
//  - Port
//  - ServiceName
type TEndpoint struct {
  Ipv4 int32 `thrift:"ipv4,1" db:"ipv4" json:"ipv4"`
  Port int16 `thrift:"port,2" db:"port" json:"port"`
  ServiceName string `thrift:"serviceName,3" db:"serviceName" json:"serviceName"`
}

func NewTEndpoint() *TEndpoint {
  return &TEndpoint{}
}


func (p *TEndpoint) GetIpv4() int32 {
  return p.Ipv4
}

func (p *TEndpoint) GetPort() int16 {
  return p.Port
}

func (p *TEndpoint) GetServiceName() string {
  return p.ServiceName
}
func (p *TEndpoint) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.I16 {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TEndpoint)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Ipv4 = v
}
  return nil
}

func (p *TEndpoint)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI16(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Port = v
}
  return nil
}

func (p *TEndpoint)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.ServiceName = v
}
  return nil
}

func (p *TEndpoint) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("TEndpoint"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TEndpoint) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("ipv4", thrift.I32, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:ipv4: ", p), err) }
  if err := oprot.WriteI32(int32(p.Ipv4)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.ipv4 (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:ipv4: ", p), err) }
  return err
}

func (p *TEndpoint) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("port", thrift.I16, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:port: ", p), err) }
  if err := oprot.WriteI16(int16(p.Port)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.port (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:port: ", p), err) }
  return err
}

func (p *TEndpoint) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("serviceName", thrift.STRING, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:serviceName: ", p), err) }
  if err := oprot.WriteString(string(p.ServiceName)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.serviceName (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:serviceName: ", p), err) }
  return err
}

func (p *TEndpoint) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TEndpoint(%+v)", *p)
}

// Attributes:
//  - Key
//  - Value
//  - Type
//  - Endpoint
type TBinaryAnnotation struct {
  Key string `thrift:"key,1" db:"key" json:"key"`
  Value string `thrift:"value,2" db:"value" json:"value"`
  Type int32 `thrift:"type,3" db:"type" json:"type"`
  Endpoint *TEndpoint `thrift:"endpoint,4" db:"endpoint" json:"endpoint,omitempty"`
}

func NewTBinaryAnnotation() *TBinaryAnnotation {
  return &TBinaryAnnotation{}
}


func (p *TBinaryAnnotation) GetKey() string {
  return p.Key
}

func (p *TBinaryAnnotation) GetValue() string {
  return p.Value
}

func (p *TBinaryAnnotation) GetType() int32 {
  return p.Type
}
var TBinaryAnnotation_Endpoint_DEFAULT *TEndpoint
func (p *TBinaryAnnotation) GetEndpoint() *TEndpoint {
  if !p.IsSetEndpoint() {
    return TBinaryAnnotation_Endpoint_DEFAULT
  }
return p.Endpoint
}
func (p *TBinaryAnnotation) IsSetEndpoint() bool {
  return p.Endpoint != nil
}

func (p *TBinaryAnnotation) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField1(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 2:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField2(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 3:
      if fieldTypeId == thrift.I32 {
        if err := p.ReadField3(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 4:
      if fieldTypeId == thrift.STRUCT {
        if err := p.ReadField4(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *TBinaryAnnotation)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  p.Key = v
}
  return nil
}

func (p *TBinaryAnnotation)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 2: ", err)
} else {
  p.Value = v
}
  return nil
}

func (p *TBinaryAnnotation)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return thrift.PrependError("error reading field 3: ", err)
} else {
  p.Type = v
}
  return nil
}

func (p *TBinaryAnnotation)  ReadField4(iprot thrift.TProtocol) error {
  p.Endpoint = &TEndpoint{}
  if err := p.Endpoint.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Endpoint), err)
  }
  return nil
}

func (p *TBinaryAnnotation) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("TBinaryAnnotation"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField1(oprot); err != nil { return err }
    if err := p.writeField2(oprot); err != nil { return err }
    if err := p.writeField3(oprot); err != nil { return err }
    if err := p.writeField4(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TBinaryAnnotation) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("key", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:key: ", p), err) }
  if err := oprot.WriteString(string(p.Key)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.key (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:key: ", p), err) }
  return err
}

func (p *TBinaryAnnotation) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("value", thrift.STRING, 2); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:value: ", p), err) }
  if err := oprot.WriteString(string(p.Value)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.value (2) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 2:value: ", p), err) }
  return err
}

func (p *TBinaryAnnotation) writeField3(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("type", thrift.I32, 3); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:type: ", p), err) }
  if err := oprot.WriteI32(int32(p.Type)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.type (3) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 3:type: ", p), err) }
  return err
}

func (p *TBinaryAnnotation) writeField4(oprot thrift.TProtocol) (err error) {
  if p.IsSetEndpoint() {
    if err := oprot.WriteFieldBegin("endpoint", thrift.STRUCT, 4); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:endpoint: ", p), err) }
    if err := p.Endpoint.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Endpoint), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 4:endpoint: ", p), err) }
  }
  return err
}

func (p *TBinaryAnnotation) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TBinaryAnnotation(%+v)", *p)
}

