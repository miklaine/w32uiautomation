package w32uiautomation

import (
	"github.com/go-ole/go-ole"
	"math"
	"syscall"
	"unsafe"
)

type IUIAutomationTransformPattern struct {
	ole.IUnknown
}

type IUIAutomationTransformPatternVtbl struct {
	ole.IUnknownVtbl
	Move                   			uintptr
	Resize          				uintptr
	Rotate      					uintptr
	Get_CurrentCanMove 				uintptr
	Get_CurrentCanResize 			uintptr
	Get_CurrentCanRotate 			uintptr
	Get_CachedCanMove 				uintptr
	Get_CachedCanResize 			uintptr
	Get_CachedCanRotate 			uintptr
}

var IID_IUIAutomationTransformPattern = &ole.GUID{0xa9b55844, 0xa55d, 0x4ef0, [8]byte{0x92, 0x6d, 0x56, 0x9c, 0x16, 0xff, 0x89, 0xbb}}

func (pat *IUIAutomationTransformPattern) VTable() *IUIAutomationTransformPatternVtbl {
	return (*IUIAutomationTransformPatternVtbl)(unsafe.Pointer(pat.RawVTable))
}

func (pat *IUIAutomationTransformPattern) Resize(width, height int) error {
	return resize(pat, float64(width), float64(height))
}

func resize(pat *IUIAutomationTransformPattern, width, height float64) error {

	hr, _, _ := syscall.Syscall6(
		pat.VTable().Resize,
		3,
		uintptr(unsafe.Pointer(pat)),
		uintptr(math.Float64bits(width)),
		uintptr(math.Float64bits(height)),
		0,
		0,
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}
	return nil
}
