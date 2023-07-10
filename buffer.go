// Package bufrw provides functionality for reading and writing
// binary data with minimal memory allocations.

package bufrw

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
)

// Buffer provides utility methods for reading and writing binary data
// with minimal memory allocations.
type Buffer struct {
	b       []byte
	maxSize int
}

// NewBuffer creates a new buffer with an internal byte buffer of the
// specified size
func NewBuffer(size int) *Buffer {
	return &Buffer{b: make([]byte, size), maxSize: size}
}

// NewBufferSize creates a new buffer with an internal byte buffer of the
// specified size
// func NewBufferSize(size, maxSize int) *Buffer {
// 	return &Buffer{b: make([]byte, size), maxSize: maxSize}
// }

// borrow returns a byte slice of length n, borrowing from the internal
// byte slice of the buffer. If n is larger than the internal buffer,
// the internal buffer is grown to fit n bytes, unless n is larger than
// the specified max size, in which case the internal buffer is only
// grown to the max size and a temporary byte slice is returned.
//
// If n is less than 0, the underlying byte slice is returned directly
// in its full length.
//
// Consumers must not retain the returned slice or use multiple borrowed
// buffers at the same time, since they all point to the same underlying
// slice.
func (buf *Buffer) borrow(n int) []byte {
	if n < 0 {
		return buf.b
	}
	if n > len(buf.b) {
		// TODO: respect maxSize
		buf.b = make([]byte, n)
	}
	return buf.b[:n]
}

// Read reads n bytes from r.
func (buf *Buffer) Read(r io.Reader, n int) ([]byte, error) {
	b := buf.borrow(n)
	if _, err := io.ReadFull(r, b); err != nil {
		return nil, err
	}
	return b, nil
}

// WriteBool writes a boolean value to w.
func (buf *Buffer) WriteBool(w io.Writer, val bool) error {
	if val {
		return buf.WriteByteValue(w, 1)
	}
	return buf.WriteByteValue(w, 0)
}

// ReadBool reads a boolean value from r, where r reads from a source
// that has used WriteBool to write a boolean value.
func (buf *Buffer) ReadBool(r io.Reader) (bool, error) {
	b, err := buf.ReadByteValue(r)
	return b == 1, err
}

// WriteBools writes zero or more boolean values to w.
func (buf *Buffer) WriteBools(w io.Writer, val ...bool) error {
	if err := buf.WriteInt(w, len(val)); err != nil {
		return err
	}
	for _, v := range val {
		if err := buf.WriteBool(w, v); err != nil {
			return err
		}
	}
	return nil
}

// ReadBools reads zero or more boolean values from r, where r reads
// from a source that has used WriteBools to write the boolean values.
func (buf *Buffer) ReadBools(r io.Reader) ([]bool, error) {
	n, err := buf.ReadInt(r)
	if err != nil {
		return nil, err
	}
	values := make([]bool, n)
	for i := 0; i < n; i++ {
		if values[i], err = buf.ReadBool(r); err != nil {
			return nil, err
		}
	}
	return values, nil
}

// WriteByteValue writes a single byte to w.
func (buf *Buffer) WriteByteValue(w io.Writer, val byte) error {
	b := buf.borrow(1)
	b[0] = val
	_, err := w.Write(b)
	return err
}

// ReadByteValue reads a single byte from r.
func (buf *Buffer) ReadByteValue(r io.Reader) (byte, error) {
	b, err := buf.Read(r, 1)
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

// WriteByteValues writes zero or more single byte values to w.
func (buf *Buffer) WriteByteValues(w io.Writer, val ...byte) error {
	if err := buf.WriteInt(w, len(val)); err != nil {
		return err
	}
	_, err := w.Write(val)
	return err
}

// ReadByteValues reads zero or more single byte values from r, where r reads
// from a source that has used WriteByteValues to write the byte values.
func (buf *Buffer) ReadByteValues(r io.Reader) ([]byte, error) {
	n, err := buf.ReadInt(r)
	if err != nil {
		return nil, err
	}
	values := make([]byte, n)
	for i := 0; i < n; i++ {
		if values[i], err = buf.ReadByteValue(r); err != nil {
			return nil, err
		}
	}
	return values, nil
}

// WriteInt writes an int to w. The value must be within the range of
// a +/- 32 bit integer.
func (buf *Buffer) WriteInt(w io.Writer, val int) error {
	if val > math.MaxInt32 || val < math.MinInt32 {
		return errors.New("util/binary/Buffer.WriteInt(): value must be in int32 range")
	}
	b := buf.borrow(4)
	if val < 0 {
		val = -val + math.MaxInt32
	}
	binary.BigEndian.PutUint32(b, uint32(val))
	_, err := w.Write(b)
	return err
}

// ReadInt reads an integer from r, where r reads from a source
// that has used WriteInt to write an int value.
func (buf *Buffer) ReadInt(r io.Reader) (int, error) {
	b, err := buf.Read(r, 4)
	if err != nil {
		return 0, err
	}
	val := binary.BigEndian.Uint32(b)
	if val <= math.MaxInt32 {
		return int(val), nil
	}
	v := val - math.MaxInt32
	return -int(v), nil
}

// WriteInts writes zero or more int values to w.
func (buf *Buffer) WriteInts(w io.Writer, val ...int) error {
	if err := buf.WriteInt(w, len(val)); err != nil {
		return err
	}
	for _, v := range val {
		if err := buf.WriteInt(w, v); err != nil {
			return err
		}
	}
	return nil
}

// ReadInts reads zero or more inte values from r, where r reads
// from a source that has used WriteInts to write the inte values.
func (buf *Buffer) ReadInts(r io.Reader) ([]int, error) {
	n, err := buf.ReadInt(r)
	if err != nil {
		return nil, err
	}
	values := make([]int, n)
	for i := 0; i < n; i++ {
		if values[i], err = buf.ReadInt(r); err != nil {
			return nil, err
		}
	}
	return values, nil
}

// WriteInt64 write an int64 value to w.
func (buf *Buffer) WriteInt64(w io.Writer, val int64) error {
	b := buf.borrow(8)
	binary.BigEndian.PutUint64(b, uint64(val))
	_, err := w.Write(b)
	return err
}

// ReadInt64 reads a int64 value from r, where r reads from a source
// that has used WriteInt64 to write an int64 value.
func (buf *Buffer) ReadInt64(r io.Reader) (int64, error) {
	b, err := buf.Read(r, 8)
	if err != nil {
		return 0, err
	}
	return int64(binary.BigEndian.Uint64(b)), nil
}

// WriteInt64s writes zero or more int64 values to w.
func (buf *Buffer) WriteInt64s(w io.Writer, val ...int64) error {
	if err := buf.WriteInt(w, len(val)); err != nil {
		return err
	}
	for _, v := range val {
		if err := buf.WriteInt64(w, v); err != nil {
			return err
		}
	}
	return nil
}

// ReadInt64s reads zero or more int64 values from r, where r reads
// from a source that has used WriteInt64s to write the int64 values.
func (buf *Buffer) ReadInt64s(r io.Reader) ([]int64, error) {
	n, err := buf.ReadInt(r)
	if err != nil {
		return nil, err
	}
	values := make([]int64, n)
	for i := 0; i < n; i++ {
		if values[i], err = buf.ReadInt64(r); err != nil {
			return nil, err
		}
	}
	return values, nil
}

// WriteFloat64 write a float64 value to w.
func (buf *Buffer) WriteFloat64(w io.Writer, val float64) error {
	b := buf.borrow(8)
	binary.BigEndian.PutUint64(b, math.Float64bits(val))
	_, err := w.Write(b)
	return err
}

// ReadFloat64 reads a float64 value from r, where r reads from a source
// that has used WriteFloat64 to write a float64 value.
func (buf *Buffer) ReadFloat64(r io.Reader) (float64, error) {
	b, err := buf.Read(r, 8)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.BigEndian.Uint64(b)), nil
}

// WriteFloat64s writes zero or more float64 values to w.
func (buf *Buffer) WriteFloat64s(w io.Writer, val ...float64) error {
	if err := buf.WriteInt(w, len(val)); err != nil {
		return err
	}
	for _, v := range val {
		if err := buf.WriteFloat64(w, v); err != nil {
			return err
		}
	}
	return nil
}

// ReadFloat64s reads zero or more float64 values from r, where r reads
// from a source that has used WriteFloat64s to write the float64 values.
func (buf *Buffer) ReadFloat64s(r io.Reader) ([]float64, error) {
	n, err := buf.ReadInt(r)
	if err != nil {
		return nil, err
	}
	values := make([]float64, n)
	for i := 0; i < n; i++ {
		if values[i], err = buf.ReadFloat64(r); err != nil {
			return nil, err
		}
	}
	return values, nil
}

// WriteString writes a string value to w.
func (buf *Buffer) WriteString(w io.Writer, val string) error {
	n := len(val)
	if err := buf.WriteInt(w, n); err != nil {
		return err
	}
	b := buf.borrow(n)
	copy(b, val)
	_, err := w.Write(b)
	return err
}

// ReadString reads a string value from r, where r reads from a source
// that has used WriteString to write a string value.
func (buf *Buffer) ReadString(r io.Reader) (string, error) {
	n, err := buf.ReadInt(r)
	if err != nil {
		return "", err
	}
	b, err := buf.Read(r, n)
	return string(b), err
}

// WriteStrings writes zero or more string values to w.
func (buf *Buffer) WriteStrings(w io.Writer, val ...string) error {
	if err := buf.WriteInt(w, len(val)); err != nil {
		return err
	}
	for _, v := range val {
		if err := buf.WriteString(w, v); err != nil {
			return err
		}
	}
	return nil
}

// ReadStrings reads zero or more string values from r, where r reads
// from a source that has used WriteStrings to write the string values.
func (buf *Buffer) ReadStrings(r io.Reader) ([]string, error) {
	n, err := buf.ReadInt(r)
	if err != nil {
		return nil, err
	}
	values := make([]string, n)
	for i := 0; i < n; i++ {
		if values[i], err = buf.ReadString(r); err != nil {
			return nil, err
		}
	}
	return values, nil
}

// Serializable describes an object that can serialize itself into a byte slice
// and deserialize itself from the byte slice.
type Serializable interface {
	Serialize() ([]byte, error)
	Deserialize(b []byte) error
}

// WriteSerializable writes a serializable object to w. If val implements
// SerializableToBufRW, WriteSerializableBufRW will be called instead, passing in
// buf as the buffer.
func (buf *Buffer) WriteSerializable(w io.Writer, val Serializable) error {
	if s, ok := val.(SerializableToBufRW); ok {
		return buf.WriteSerializableBufRW(w, s)
	}
	b, err := val.Serialize()
	if err != nil {
		return err
	}
	return buf.WriteByteValues(w, b...)
}

// ReadSerializable reads a serializable value from r, where r reads from a source
// produced by WriteSerializable. If val implements SerializableToBufRW,
// ReadSerializableBufRW will be called instead, passing in buf as the buffer.
// buf as the buffer.
func (buf *Buffer) ReadSerializable(r io.Reader, val Serializable) error {
	if s, ok := val.(SerializableToBufRW); ok {
		return buf.ReadSerializableBufRW(r, s)
	}
	b, err := buf.ReadByteValues(r)
	if err != nil {
		return err
	}
	return val.Deserialize(b)
}

// SerializableToBufRW describes an object that can serialize itself to and deserialize
// itself from a io.Writer/io.Reader with the help of a buffer from this package.
type SerializableToBufRW interface {
	SerializeToBufRW(w io.Writer, buf *Buffer) error
	DeserializeFromBufRW(r io.Reader, buf *Buffer) error
}

// WriteSerializableBufRW writes a serializable object to w.
func (buf *Buffer) WriteSerializableBufRW(w io.Writer, val SerializableToBufRW) error {
	return val.SerializeToBufRW(w, buf)
}

// ReadSerializableBufRW reads a serializable value from r, where r reads from a source
// that has used ReadSerializableBufRW to write a serialized value of val.
func (buf *Buffer) ReadSerializableBufRW(r io.Reader, val SerializableToBufRW) error {
	return val.DeserializeFromBufRW(r, buf)
}
