package compress

import "os"

type Kind int32

const (
	Nil Kind = iota
	Zip
)

type Compressor interface {
	Compress(output string, files ...*os.File) error
	UnCompress(zipFile, output string) error
}

// NewCompressor
/**
 * @Description:
 * @param compressType
 * @return Compressor
 */
func NewCompressor(compressType Kind) Compressor {
	switch compressType {
	case Zip:
		return &ZipCompress{
			Type: compressType,
		}
	case Nil:
		return nil
	default:
		return nil
	}
}
