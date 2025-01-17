package w32uiautomation

import (
	"errors"
	"unsafe"
)

func Select(element *IUIAutomationElement) error {
	unknown, err := element.GetCurrentPattern(UIA_SelectionItemPatternId)
	if err != nil {
		return err
	}
	if unknown == nil {
		return errors.New("this pattern is not available for the element")
	}
	defer unknown.Release()

	disp, err := unknown.QueryInterface(IID_IUIAutomationSelectionItemPattern)
	if err != nil {
		return err
	}

	pattern := (*IUIAutomationSelectionItemPattern)(unsafe.Pointer(disp))
	defer pattern.Release()
	err = pattern.Select()
	if err != nil {
		return err
	}
	return nil
}
