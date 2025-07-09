

package calc 

import "errors"

func Divide(a int , b int)  (int , error) {

	if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
}
