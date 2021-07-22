package w32uiautomation

import (
	"errors"
	"unsafe"
)

func CloseWindow(element *IUIAutomationElement) error {
	unknown, err := element.GetCurrentPattern(UIA_WindowPatternId)
	if err != nil {
		return err
	}
	if unknown == nil {
		return errors.New("this pattern is not available for the element")
	}
	defer unknown.Release()

	disp, err := unknown.QueryInterface(IID_IUIAutomationWindowPattern)
	if err != nil {
		return err
	}

	pattern := (*IUIAutomationWindowPattern)(unsafe.Pointer(disp))
	defer pattern.Release()
	err = pattern.Close()
	if err != nil {
		return err
	}
	return nil
}
