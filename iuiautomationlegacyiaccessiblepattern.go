package w32uiautomation

import (
	"github.com/go-ole/go-ole"
	"syscall"
	"unsafe"
)

const (
	SELFLAG_NONE 			= 0x0
	SELFLAG_TAKEFOCUS 		= 0x1
	SELFLAG_TAKESELECTION 	= 0x2
	SELFLAG_EXTENDSELECTION = 0x4
	SELFLAG_ADDSELECTION 	= 0x8
	SELFLAG_REMOVESELECTION = 0x10
)

type IUIAutomationLegacyIAccessiblePattern struct {
	ole.IUnknown
}

type IUIAutomationLegacyIAccessiblePatternVtbl struct {
	ole.IUnknownVtbl
	Select uintptr
	DoDefaultAction uintptr
	SetValue uintptr
	Get_CurrentChildId uintptr
	Get_CurrentName uintptr
	Get_CurrentValue uintptr
	get_CurrentDescription uintptr
	Get_CurrentRole uintptr
	Get_CurrentState uintptr
	Get_CurrentHelp uintptr
	Get_CurrentKeyboardShortcut uintptr
	Get_CurrentSelection uintptr
	Get_CurrentDefaultAction uintptr
	Get_CachedChildId uintptr
	Get_CachedName uintptr
	Get_CachedValue uintptr
	get_CachedDescription uintptr
	Get_CachedRole uintptr
	Get_CachedState uintptr
	Get_CachedHelp uintptr
	Get_CachedKeyboardShortcut uintptr
	Get_CachedSelection uintptr
	Get_CachedDefaultAction uintptr
	GetIAccessible uintptr
}

var IID_IUIAutomationLegacyIAccessiblePattern = &ole.GUID{0x828055ad, 0x355b, 0x4435, [8]byte{0x86, 0xd5, 0x3b, 0x51, 0xc1, 0x4a, 0x9b, 0x1b}}

func (pat *IUIAutomationLegacyIAccessiblePattern) VTable() *IUIAutomationLegacyIAccessiblePatternVtbl {
	return (*IUIAutomationLegacyIAccessiblePatternVtbl)(unsafe.Pointer(pat.RawVTable))
}

func (pat *IUIAutomationLegacyIAccessiblePattern) Select() error {
	return uiSelect(pat, SELFLAG_TAKEFOCUS | SELFLAG_TAKESELECTION)
}

func (pat *IUIAutomationLegacyIAccessiblePattern) DoDefaultAction() error {
	return doDefaultAction(pat)
}

func uiSelect(pat *IUIAutomationLegacyIAccessiblePattern, flags int32) error {
	hr, _, _ := syscall.Syscall(
		pat.VTable().Select,
		2,
		uintptr(unsafe.Pointer(pat)),
		uintptr(flags),
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}
	return nil
}

func doDefaultAction(pat *IUIAutomationLegacyIAccessiblePattern) error {
	hr, _, _ := syscall.Syscall(
		pat.VTable().DoDefaultAction,
		1,
		uintptr(unsafe.Pointer(pat)),
		0,
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}
	return nil
}
