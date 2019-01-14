package syserror

import ("testing"
	"fmt")

func TestErrorCodes (t *testing.T) {
	fmt.Println ("Testing NoErr --> ");
	fmt.Printf ("\t%s\n", NoErr.Error());
	fmt.Printf ("\t%s\n", NoErr.getMsg());
	fmt.Printf ("\t%d\n", NoErr.getCode());

	fmt.Println ("Testing ErrFatal --> ");
	fmt.Printf ("\t%s\n", ErrFatal.Error());
	fmt.Printf ("\t%s\n", ErrFatal.getMsg());
	fmt.Printf ("\t%d\n", ErrFatal.getCode());

	fmt.Println ("Testing ErrOutOfRes --> ");
	fmt.Printf ("\t%s\n", ErrOutOfRes.Error());
	fmt.Printf ("\t%s\n", ErrOutOfRes.getMsg());
	fmt.Printf ("\t%d\n", ErrOutOfRes.getCode());

	fmt.Println ("Testing ErrNotAvailable --> ");
	fmt.Printf ("\t%s\n", ErrNotAvailable.Error());
	fmt.Printf ("\t%s\n", ErrNotAvailable.getMsg());
	fmt.Printf ("\t%d\n", ErrNotAvailable.getCode());
}
