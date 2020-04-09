package schreibvogel

import (
	"fmt"
)

type foundSlice struct {
	found bool
	item  string
}

func InSlice(needed []string, slice []string) error {
	ch := make(chan foundSlice)
	defer close(ch)

	for _, n := range needed {
		go func(i string, s []string) {
			for _, j := range s {
				if j == i {
					ch <- foundSlice{item: i, found: true}
					return
				}
			}
			ch <- foundSlice{item: i, found: false}
		}(n, slice)
	}

	for i := 1; i <= len(needed); i++ {
		val := <-ch
		if !val.found {
			return fmt.Errorf("%q item not in slice", val.item)
		}
	}

	return nil
}
