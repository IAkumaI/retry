package retry

import (
	"errors"
	"testing"
)

func TestDo(t *testing.T) {
	t.Run("max retry count", func(t *testing.T) {
		counter := 0
		err := Do(0, func(_ int) error {
			counter++
			return nil
		})

		if err != nil {
			t.Errorf("Must do not return error")
		}

		counter = 0
		err = Do(10, func(_ int) error {
			counter++
			for counter < 5 {
				return errors.New("it is normal")
			}

			return nil
		})

		if err != nil {
			t.Errorf("Must do not return error")
		}
		if counter != 5 {
			t.Errorf("Counter must be 5, but == %d", counter)
		}

		counter = 0
		err = Do(4, func(_ int) error {
			counter++
			for counter < 5 {
				return errors.New("it is normal")
			}

			return nil
		})
		if err == nil {
			t.Errorf("Must return error")
		}
		if counter != 4 {
			t.Errorf("Counter must be 5, but == %d", counter)
		}
	})
}
