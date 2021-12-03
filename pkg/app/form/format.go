package form

import "fmt"

func FormatBindErr(errs ValidErrors) string {
	return fmt.Sprintf("BindAndValid err: %v", errs)
}

func GenerateTokenErr(err error) string {
	return fmt.Sprintf("Err GenerateToken err: %v", err)
}
