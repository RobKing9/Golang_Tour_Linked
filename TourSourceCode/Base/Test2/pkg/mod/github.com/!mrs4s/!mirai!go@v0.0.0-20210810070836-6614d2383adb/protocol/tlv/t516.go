package tlv

import "github.com/Mrs4s/MiraiGo/binary"

func T516() []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x516)
		w.WriteBytesShort(binary.NewWriterF(func(w *binary.Writer) {
			w.WriteUInt32(0)
		}))
	})
}
