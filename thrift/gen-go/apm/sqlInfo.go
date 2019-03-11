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
//  - SqlId
//  - ResouceId
//  - Sql
type TSQLMeta struct {
  // unused fields # 1 to 9
  SqlId string `thrift:"sqlId,10" db:"sqlId" json:"sqlId"`
  // unused fields # 11 to 19
  ResouceId string `thrift:"resouceId,20" db:"resouceId" json:"resouceId"`
  // unused fields # 21 to 29
  Sql string `thrift:"sql,30" db:"sql" json:"sql"`
}

func NewTSQLMeta() *TSQLMeta {
  return &TSQLMeta{}
}


func (p *TSQLMeta) GetSqlId() string {
  return p.SqlId
}

func (p *TSQLMeta) GetResouceId() string {
  return p.ResouceId
}

func (p *TSQLMeta) GetSql() string {
  return p.Sql
}
func (p *TSQLMeta) Read(iprot thrift.TProtocol) error {
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
    case 10:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField10(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 20:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField20(iprot); err != nil {
          return err
        }
      } else {
        if err := iprot.Skip(fieldTypeId); err != nil {
          return err
        }
      }
    case 30:
      if fieldTypeId == thrift.STRING {
        if err := p.ReadField30(iprot); err != nil {
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

func (p *TSQLMeta)  ReadField10(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 10: ", err)
} else {
  p.SqlId = v
}
  return nil
}

func (p *TSQLMeta)  ReadField20(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 20: ", err)
} else {
  p.ResouceId = v
}
  return nil
}

func (p *TSQLMeta)  ReadField30(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 30: ", err)
} else {
  p.Sql = v
}
  return nil
}

func (p *TSQLMeta) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("TSQLMeta"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if p != nil {
    if err := p.writeField10(oprot); err != nil { return err }
    if err := p.writeField20(oprot); err != nil { return err }
    if err := p.writeField30(oprot); err != nil { return err }
  }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *TSQLMeta) writeField10(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("sqlId", thrift.STRING, 10); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 10:sqlId: ", p), err) }
  if err := oprot.WriteString(string(p.SqlId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.sqlId (10) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 10:sqlId: ", p), err) }
  return err
}

func (p *TSQLMeta) writeField20(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("resouceId", thrift.STRING, 20); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 20:resouceId: ", p), err) }
  if err := oprot.WriteString(string(p.ResouceId)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.resouceId (20) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 20:resouceId: ", p), err) }
  return err
}

func (p *TSQLMeta) writeField30(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("sql", thrift.STRING, 30); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 30:sql: ", p), err) }
  if err := oprot.WriteString(string(p.Sql)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.sql (30) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 30:sql: ", p), err) }
  return err
}

func (p *TSQLMeta) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("TSQLMeta(%+v)", *p)
}

