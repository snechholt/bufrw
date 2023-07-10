package bufrw

import "io"

type Writer struct {
	buf         *Buffer
	w           io.Writer
	stopOnError bool
	err         error
}

func (w *Writer) do(fn func() error) error {
	if w.err != nil && w.stopOnError {
		return w.err
	}
	w.err = fn()
	return w.err
}

func (w *Writer) WriteBool(val bool) error {
	return w.do(func() error { return w.buf.WriteBool(w.w, val) })
}

func (w *Writer) WriteBools(val ...bool) error {
	return w.do(func() error { return w.buf.WriteBools(w.w, val...) })
}

func (w *Writer) WriteByteValue(val byte) error {
	return w.do(func() error { return w.buf.WriteByteValue(w.w, val) })
}

func (w *Writer) WriteByteValues(val ...byte) error {
	return w.do(func() error { return w.buf.WriteByteValues(w.w, val...) })
}

func (w *Writer) WriteInt(val int) error {
	return w.do(func() error { return w.buf.WriteInt(w.w, val) })
}

func (w *Writer) WriteInts(val ...int) error {
	return w.do(func() error { return w.buf.WriteInts(w.w, val...) })
}

func (w *Writer) WriteInt64(val int64) error {
	return w.do(func() error { return w.buf.WriteInt64(w.w, val) })
}

func (w *Writer) WriteInt64s(val ...int64) error {
	return w.do(func() error { return w.buf.WriteInt64s(w.w, val...) })
}

func (w *Writer) WriteFloat64(val float64) error {
	return w.do(func() error { return w.buf.WriteFloat64(w.w, val) })
}

func (w *Writer) WriteFloat64s(val ...float64) error {
	return w.do(func() error { return w.buf.WriteFloat64s(w.w, val...) })
}

func (w *Writer) WriteString(val string) error {
	return w.do(func() error { return w.buf.WriteString(w.w, val) })
}

func (w *Writer) WriteStrings(val ...string) error {
	return w.do(func() error { return w.buf.WriteStrings(w.w, val...) })
}

func (w *Writer) WriteSerializable(val Serializable) error {
	return w.do(func() error { return w.buf.WriteSerializable(w.w, val) })
}

func (w *Writer) Err() error {
	return w.err
}

func (buf *Buffer) Writer(w io.Writer, stopOnError ...bool) *Writer {
	return &Writer{buf: buf, w: w, stopOnError: len(stopOnError) > 0 && stopOnError[0]}
}

type Reader struct {
	buf *Buffer
	r   io.Reader
}

func (r *Reader) Read(n int) ([]byte, error)              { return r.buf.Read(r.r, n) }
func (r *Reader) ReadBool() (bool, error)                 { return r.buf.ReadBool(r.r) }
func (r *Reader) ReadBools() ([]bool, error)              { return r.buf.ReadBools(r.r) }
func (r *Reader) ReadByteValue() (byte, error)            { return r.buf.ReadByteValue(r.r) }
func (r *Reader) ReadByteValues() ([]byte, error)         { return r.buf.ReadByteValues(r.r) }
func (r *Reader) ReadInt() (int, error)                   { return r.buf.ReadInt(r.r) }
func (r *Reader) ReadInts() ([]int, error)                { return r.buf.ReadInts(r.r) }
func (r *Reader) ReadInt64() (int64, error)               { return r.buf.ReadInt64(r.r) }
func (r *Reader) ReadInt64s() ([]int64, error)            { return r.buf.ReadInt64s(r.r) }
func (r *Reader) ReadFloat64() (float64, error)           { return r.buf.ReadFloat64(r.r) }
func (r *Reader) ReadFloat64s() ([]float64, error)        { return r.buf.ReadFloat64s(r.r) }
func (r *Reader) ReadString() (string, error)             { return r.buf.ReadString(r.r) }
func (r *Reader) ReadStrings() ([]string, error)          { return r.buf.ReadStrings(r.r) }
func (r *Reader) ReadSerializable(val Serializable) error { return r.buf.ReadSerializable(r.r, val) }

func (buf *Buffer) Reader(r io.Reader) *Reader {
	return &Reader{buf: buf, r: r}
}
