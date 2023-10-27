package lang

import (
	"github.com/murakmii/gj/vm"
	"runtime"
)

func RuntimeGetAvailableProcessors(thread *vm.Thread, _ []interface{}) error {
	thread.CurrentFrame().PushOperand(int32(runtime.NumCPU()))
	return nil
}
