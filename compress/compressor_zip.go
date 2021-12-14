package compress

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"strings"
)

var _ = Compressor(&ZipCompress{})

// ZipCompress "Zip Compress"
type ZipCompress struct {
	Type Kind
}

// Compress
/**
 * @Description:
 * @receiver c
 * @param output
 * @param files
 * @return error
 */
func (c *ZipCompress) Compress(output string, files ...*os.File) error {
	compressionFile, err := os.Create(output)
	defer compressionFile.Close() // nolint:gosec
	if err != nil {
		return err
	}
	w := zip.NewWriter(compressionFile)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

// compress
/**
 * @Description:
 * @param file
 * @param zw
 * @return error
 */
func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		if prefix == "" {
			prefix = info.Name()
		} else {
			prefix = prefix + "/" + info.Name()
		}
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// UnCompress
/**
 * @Description:
 * @receiver c
 * @param zipFile
 * @param output
 * @return error
 */
func (c *ZipCompress) UnCompress(zipFile, output string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		err := c.unCompress(file, output)
		if err != nil {
			return err
		}
	}
	return nil
}

// unCompress
/**
 * @Description:
 * @receiver c
 * @param file
 * @param output
 * @return error
 */
func (c *ZipCompress) unCompress(file *zip.File, output string) error {
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	filename := path.Join(output, file.Name)
	err = os.MkdirAll(getDir(filename), 0755)
	if err != nil {
		return err
	}
	w, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = io.Copy(w, rc)
	if err != nil {
		return err
	}
	w.Close()
	rc.Close()
	return nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}
