package path

import "os"

// ExistOrNot
/**
 * @Description: path or file exist or not
 * @param path
 * @return bool
 * @return error
 */
func ExistOrNot(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Delete
/**
 * @Description: del path or file
 * @param path
 * @return error
 */
func Delete(path string) error {
	return os.RemoveAll(path)
}
