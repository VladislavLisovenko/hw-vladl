package printer

import (
	"fmt"

	"github.com/VladislavLisovenko/hw-vladl/hw06_testing/hw02fixapp/types"
)

func PrintStaff(staff []types.Employee) {
	for i := 0; i < len(staff); i++ {
		fmt.Println(staff[i])
	}
}
