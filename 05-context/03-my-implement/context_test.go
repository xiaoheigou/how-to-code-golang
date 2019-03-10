package contextimpl

import (
	"fmt"
	"testing"
	"time"
)

type empty struct {
}

func TestBackgroundNotTODO(t *testing.T) {

	todo := fmt.Sprint(TODO())
	bg := fmt.Sprint(Background())
	// prin1()
	// prin2()
	// fmt.Fprintln(os.Stderr, "------", todo, bg)
	if todo == bg {
		t.Errorf("TODO and Background are equal: %q vs %q", todo, bg)
	}

}

func TestWithCancel(t *testing.T) {
	ctx, cancel := WithCancel(Background())

	if err := ctx.Err(); err != nil {
		t.Errorf("error should be nil first, got %v", err)
	}
	cancel()

	<-ctx.Done()
	// fmt.Println(s)
	if err := ctx.Err(); err != Canceled {
		t.Errorf("error should be canceled now, got %v", err)
	}
}

func TestWithCancelConcurrent(t *testing.T) {
	ctx, cancel := WithCancel(Background())

	time.AfterFunc(1*time.Second, cancel)

	if err := ctx.Err(); err != nil {
		t.Errorf("error should be nil first, got %v", err)
	}
	<-ctx.Done()
	if err := ctx.Err(); err != Canceled {
		t.Errorf("error should be canceled now, got %v", err)
	}
}
