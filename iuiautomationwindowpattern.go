package w32uiautomation

import (
	"github.com/go-ole/go-ole"
	"syscall"
	"unsafe"
)

type IUIAutomationWindowPattern struct {
	ole.IUnknown
}

type IUIAutomationWindowPatternVtbl struct {
	ole.IUnknownVtbl
	Close                   			uintptr
	WaitForInputIdle          			uintptr
	SetWindowVisualState      			uintptr
	Get_CurrentCanMaximize 				uintptr
	Get_CurrentCanMinimize 				uintptr
	Get_CurrentIsModal 					uintptr
	Get_CurrentIsTopmost 				uintptr
	Get_CurrentWindowVisualState 		uintptr
	Get_CurrentWindowInteractionState 	uintptr
	Get_CachedCanMaximize 				uintptr
	Get_CachedCanMinimize 				uintptr
	Get_CachedIsModal 					uintptr
	Get_CachedIsTopmost 				uintptr
	Get_CachedWindowVisualState 		uintptr
	Get_CachedWindowInteractionState 	uintptr
}

var IID_IUIAutomationWindowPattern = &ole.GUID{0x0faef453, 0x9208, 0x43ef, [8]byte{0xbb, 0xb2, 0x3b, 0x48, 0x51, 0x77, 0x86, 0x4f}}

func (pat *IUIAutomationWindowPattern) VTable() *IUIAutomationWindowPatternVtbl {
	return (*IUIAutomationWindowPatternVtbl)(unsafe.Pointer(pat.RawVTable))
}

func (pat *IUIAutomationWindowPattern) Close() error {
	return closeWindow(pat)
}

func closeWindow(pat *IUIAutomationWindowPattern) error {
	hr, _, _ := syscall.Syscall(
		pat.VTable().Close,
		1,
		uintptr(unsafe.Pointer(pat)),
		0,
		0)
	if hr != 0 {
		return ole.NewError(hr)
	}
	return nil
}