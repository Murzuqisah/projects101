package functions

import (
	"fmt"
)

func CheckNilError(err error, customMsg string) error {
	if err != nil {
		return fmt.Errorf("error: %s\n%w", customMsg, err)
	}
	return nil
}
