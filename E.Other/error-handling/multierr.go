// You can edit this code!
// Click here and start typing.
package error_handling

import (
	"errors"
	"fmt"
)

type EngineError struct{}
type joinError struct {
	errs []error
}

func (e *EngineError) Error() string { return "catastrophic engine failure" }

func main() {
	var (
		ErrRelayOrientation = errors.New("bad relay orientation")
		ErrCosmicRayBitflip = errors.New("cosmic ray bitflip")
		ErrStickyPlunger    = errors.New("sticky sensor plunger")
	)

	err1 := fmt.Errorf("G-switch failed: %w %w %w", ErrRelayOrientation, ErrCosmicRayBitflip, ErrStickyPlunger)
	// 2009/11/10 23:00:00 G-switch failed: bad relay orientation cosmic ray bitflip sticky sensor plunger
	// log.Fatal(err1)

	err2 := errors.Join(
		ErrRelayOrientation,
		ErrCosmicRayBitflip,
		ErrStickyPlunger,
	)

	// 2009/11/10 23:00:00 bad relay orientation
	// cosmic ray bitflip
	// sticky sensor plunger
	// log.Fatal(err2)

	ok := errors.Is(err1, ErrStickyPlunger)
	fmt.Println(ok) // true

	var engineErr *EngineError
	ok = errors.As(err2, &engineErr)
	fmt.Println(ok) // false

	fmt.Println(errors.Unwrap(err1)) // nil
	fmt.Println(errors.Unwrap(err2)) // nil

	var joinedErrors interface{ Unwrap() []error }

	// You can use errors.As to make sure that the alternate Unwrap() implementation is available
	if errors.As(err1, &joinedErrors) {
		for _, e := range joinedErrors.Unwrap() {
			fmt.Println("-", e)
		}
	}

	// Or do it more directly with an inline cast
	if uw, ok := err2.(interface{ Unwrap() []error }); ok {
		for _, e := range uw.Unwrap() {
			fmt.Println("~", e)
		}
	}

	//err1 := fmt.Errorf("G-switch failed: %w\n%w\n%w", ErrRelayOrientation, ErrCosmicRayBitflip, ErrStickyPlunger)

	//err2 := fmt.Errorf("G-switch failed: %w", errors.Join(
	//	ErrRelayOrientation,
	//	ErrCosmicRayBitflip,
	//	ErrStickyPlunger,
	//))
	//
	//err3 := errors.Join(
	//	ErrRelayOrientation,
	//	ErrCosmicRayBitflip,
	//	ErrStickyPlunger,
	//)

	// &fmt.wrapErrors{msg:"bad relay orientation\ncosmic ray bitflip\nsticky sensor plunger", errs:[]error{(*errors.errorString)(0xc00009c050), (*errors.errorString)(0xc00009c060), (*errors.errorString)(0xc00009c070)}}
	// &fmt.wrapError{msg:"G-switch failed: bad relay orientation\ncosmic ray bitflip\nsticky sensor plunger", err:(*errors.joinError)(0xc0000be000)}
	// &errors.joinError{errs:[]error{(*errors.errorString)(0xc0000140a0), (*errors.errorString)(0xc0000140b0), (*errors.errorString)(0xc0000140c0)}}
}
