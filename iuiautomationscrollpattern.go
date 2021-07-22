package w32uiautomation

import (
	"github.com/go-ole/go-ole"
	"math"
	"syscall"
	"unsafe"
)

type IUIAutomationScrollPattern struct {
	ole.IUnknown
}

type IUIAutomationScrollPatternVtbl struct {
	ole.IUnknownVtbl
	Scroll                   			uintptr
	SetScrollPercent 					uintptr
	Get_CurrentHorizontalScrollPercent	uintptr
	Get_CurrentVerticalScrollPercent 	uintptr
	Get_CurrentHorizontalViewSize 		uintptr
	Get_CurrentVerticalViewSize 		uintptr
	Get_CurrentHorizontallyScrollable 	uintptr
	Get_CurrentVerticallyScrollable 	uintptr
	Get_CachedHorizontalScrollPercent 	uintptr
	Get_CachedVerticalScrollPercent 	uintptr
	Get_CachedHorizontalViewSize 		uintptr
	Get_CachedVerticalViewSize 			uintptr
	Get_CachedHorizontallyScrollable 	uintptr
	Get_CachedVerticallyScrollable 		uintptr
}

var IID_IUIAutomationScrollPattern = &ole.GUID{0x88f4d42a, 0xe881, 0x459d, [8]byte{0xa7, 0x7c, 0x73, 0xbb, 0xbb, 0x7e, 0x02, 0xdc}}

func (pat *IUIAutomationScrollPattern) VTable() *IUIAutomationScrollPatternVtbl {
	return (*IUIAutomationScrollPatternVtbl)(unsafe.Pointer(pat.RawVTable))
}

func (pat *IUIAutomationScrollPattern) SetScrollPercent(hScroll, vScroll float64) error {
	return setScrollPercent(pat, hScroll, vScroll)
}

func (pat *IUIAutomationScrollPattern) Get_CurrentVerticalScrollPercent() (float64, error) {
	return getCurrentVerticalScrollPercent(pat)
}

func (pat *IUIAutomationScrollPattern) Get_CurrentVerticalViewSize() (float64, error) {
	return getCurrentVerticalViewSize(pat)
}

func setScrollPercent(pat *IUIAutomationScrollPattern, hScroll, vScroll float64) error {
	hr, _, _ := syscall.Syscall(
		pat.VTable().SetScrollPercent,
		3,
		uintptr(unsafe.Pointer(pat)),
		uintptr(math.Float64bits(hScroll)),
		uintptr(math.Float64bits(vScroll)))

	if hr != 0 {
		return  ole.NewError(hr)
	}
	return nil
}

func getCurrentVerticalScrollPercent(pat *IUIAutomationScrollPattern) (ret float64, err error) {
	hr, _, _ := syscall.Syscall(
		pat.VTable().Get_CurrentVerticalScrollPercent,
		2,
		uintptr(unsafe.Pointer(pat)),
		uintptr(unsafe.Pointer(&ret)),
		0)
	_ = hr
	err = nil
	return
}

func getCurrentVerticalViewSize(pat *IUIAutomationScrollPattern) (ret float64, err error) {
	hr, _, _ := syscall.Syscall(
		pat.VTable().Get_CurrentVerticalViewSize,
		2,
		uintptr(unsafe.Pointer(pat)),
		uintptr(unsafe.Pointer(&ret)),
		0)
	_ = hr
	err = nil
	return
}
