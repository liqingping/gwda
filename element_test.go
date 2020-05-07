package gwda

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func TestElement_Tap(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	// _ = c.Homescreen()
	element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeIcon' AND visible == true"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(element)

	rect, err := element.Rect()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rect)

	Debug = true
	err = element.Tap(rect.X+4, rect.Y+4)
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_DoubleTap(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	// _ = c.Homescreen()
	element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeIcon' AND visible == true"})
	// element, err := s.FindElement(WDALocator{PartialLinkText: NewWDAElementAttribute().SetLabel("文件夹")})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(element)

	rect, err := element.Rect()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rect)

	Debug = true
	err = element.DoubleTap()
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_TwoFingerTap(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	element, err := s.FindElement(WDALocator{Name: "自定手势作用区域"})
	if err != nil {
		t.Fatal(err)
	}

	err = element.TwoFingerTap()
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_TapWithNumberOfTaps(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	element, err := s.FindElement(WDALocator{Name: "自定手势作用区域"})
	if err != nil {
		t.Fatal(err)
	}

	err = element.TapWithNumberOfTaps(3, 3)
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_TouchAndHold(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	// _ = c.Homescreen()
	element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeIcon' AND visible == true"})
	if err != nil {
		t.Fatal(err)
	}

	Debug = true
	err = element.TouchAndHold(1)
	// err = element.TouchAndHoldFloat(2.5)
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_ForceTouch(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	// launchOpt := NewWDAAppLaunchOption().SetShouldWaitForQuiescence(true).SetArguments([]string{"-AppleLanguages", "(en-US)"})
	// s.AppLaunch(bundleId, launchOpt)
	// return
	Debug = true
	// element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeCell' AND name == '灵敏度评估'"})
	// element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeIcon' AND visible == true"})
	// element, err := s.FindElement(WDALocator{ClassName: WDAElementType{Cell: true}})
	element, err := s.FindElement(WDALocator{Name: "自定手势作用区域"})
	if err != nil {
		t.Fatal(err)
	}

	// element.TouchAndHold()
	// fmt.Println(element.GetAttribute(NewWDAElementAttribute().SetName("")))
	// err = element.ForceTouch(0.67)
	// err = element.ForceTouch(0.67, 3)
	// err = element.ForceTouch(2.2, 1.0)
	// err = element.ForceTouch(2.2, 0.5)
	// err = element.ForceTouch(1, 1.0)
	err = element.ForceTouch(1)
	// err = element.ForceTouchCoordinate(WDACoordinate{100, 100}, 2)

	// err = element.ForceTouchPeek()
	// err = element.ForceTouchPop()
	if err != nil {
		t.Fatal(err)
	}

	// body := newWdaBody().set("pressure", 2.0).set("duration", 1.0)
	// 弱
	// body := newWdaBody().set("pressure", 0.52).set("duration", 6.0) // 预览	pressure > 0.51 && duration < 6.0
	// body := newWdaBody().set("pressure", 0.72).set("duration", 1.97) // 打开?	pressure > 0.71 && duration > 1.98
	// body := newWdaBody().set("pressure", 2.0).set("duration", 1.0)
	// 中 > 0.66
	// body := newWdaBody().set("pressure", 0.66).set("duration", 1.0)
	// body := newWdaBody().set("pressure", 0.67).set("duration", 1.0) // 预览 pressure > 0.66
	// body := newWdaBody().set("pressure", 2.20).set("duration", 1.0) // 打开? pressure > 2.1
	// 强 > 0.66
	// body := newWdaBody().set("pressure", 0.67).set("duration", 1.0) // 预览 pressure > 0.66
	// body := newWdaBody().set("pressure", 2.20).set("duration", 1.0) // 打开? pressure > 2.1
}

func TestElement_Drag(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	// _ = c.Homescreen()
	element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeIcon' AND visible == true"})
	if err != nil {
		t.Fatal(err)
	}

	Debug = true
	// err = element.Drag(230, 130, 230, 480, 2)
	// err = element.Drag(230, 130, 230, 480)
	// err = element.Drag(230, 130, 230, 30)
	// err = element.Drag(230, 130, 130, 130)
	err = element.Drag(230, 130, 330, 130)
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_Swipe(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	// _ = c.Homescreen()
	// element, err := s.FindElement(WDALocator{ClassName: WDAElementType{Table: true}})
	// element, err := s.FindElement(WDALocator{ClassName: WDAElementType{PageIndicator: true}})
	// element, err := s.FindElement(WDALocator{ClassName: WDAElementType{Cell: true}})
	// element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeIcon' AND visible == true"})
	element, err := s.FindElement(WDALocator{Name: "自定手势作用区域"})
	// element, err := s.FindElement(WDALocator{Name: "辅助功能"})
	// element, err := s.FindElement(WDALocator{Name: "com.apple.mobilesafari"})
	// element, err := s.FindElement(WDALocator{ClassName: WDAElementType{Image: true}})
	if err != nil {
		t.Fatal(err)
	}

	// element.ScreenshotToDisk("/Users/hero/Desktop/e8.png")
	// return
	Debug = true
	// rect, _ := element.Rect()
	// 相对元素自身的坐标
	// err = element.Swipe(rect.X, rect.Y+rect.Height/2, rect.X+rect.Width, rect.Y+rect.Height/2)
	// err = element.SwipeDirection(WDASwipeDirectionRight)
	err = element.SwipeRight()
	// err = element.SwipeLeft()
	// err = element.SwipeUp()
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_Pinch(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	// _ = c.Homescreen()
	// element, err := s.FindElement(WDALocator{ClassName: WDAElementType{Image: true}})
	element, err := s.FindElement(WDALocator{Name: "自定手势作用区域"})
	if err != nil {
		t.Fatal(err)
	}

	Debug = true
	// - 放大的捏放在屏幕中心
	// - 缩小是 左上角+右下角 向中心靠拢

	// zoom in 放大
	// err = element.Pinch(49, 50)
	// err = element.Pinch(2, 32)
	err = element.Pinch(2, 8)
	// err = element.Pinch(1.9, 3.6)
	// err = element.Pinch(1.9, 3.6)
	// err = element.PinchToZoomIn()
	// zoom out 缩小
	// - iPhoneX 测试结果无法缩小（相册照片），开启 引导式访问 也无效
	// - iPad Pro 横屏会触发右上角的 控制中心（竖屏可成功触发缩小效果）
	// err = element.Pinch(0.9, -0.9)
	// err = element.Pinch(0.9, -3.6)
	// err = element.Pinch(0.9, -14.4)
	// err = element.PinchToZoomOut()

	err = element.PinchToZoomOutByActions()
	// err = element.PinchToZoomOutByActions(23)

	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_Rotate(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	element, err := s.FindElement(WDALocator{Name: "自定手势作用区域"})
	if err != nil {
		t.Fatal(err)
	}

	Debug = true
	// 顺时针 90度
	err = element.Rotate(math.Pi / 2)
	// 逆时针 180度
	// err = element.Rotate(math.Pi * -2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_Scroll(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	element, err := s.FindElement(WDALocator{ClassName: WDAElementType{Table: true}})
	// element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeCell' AND name == '手机淘宝'"})
	if err != nil {
		t.Fatal(err)
	}

	Debug = true
	// err = element.ScrollUp()
	// err = element.ScrollDown()
	// err = element.ScrollLeft()
	// err = element.ScrollRight()
	// err = element.ScrollToVisible()
	// err = element.ScrollElementByPredicate("type == 'XCUIElementTypeCell' AND name LIKE '*Keynote*'")
	err = element.ScrollElementByPredicate("type == 'XCUIElementTypeCell' AND name == '音乐'")
	// err = element.ScrollElementByName("其他")
	// err = element.ScrollElementByName("iCloud 帐户")

	// It's not working
	// err = element.ScrollElementByName("音乐")
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_PickerWheelSelect(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	element, err := s.FindElement(WDALocator{ClassName: WDAElementType{PickerWheel: true}})
	if err != nil {
		t.Fatal(err)
	}

	Debug = true
	err = element.PickerWheelSelect(WDAPickerWheelSelectOrderNext, 3)
	err = element.PickerWheelSelectNext()
	err = element.PickerWheelSelectPrevious(3)
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_Click(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	if err != nil {
		// staleElementReferenceErrorWithMessage
		t.Fatal(err)
	}
	t.Log(element)

	err = element.Click()
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_SendKeys(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{ClassName: WDAElementType{SearchField: true}})
	if err != nil {
		t.Fatal(err)
	}

	err = element.SendKeys(bundleId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_Clear(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	defer s.DeleteSession()
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{ClassName: WDAElementType{SearchField: true}})
	if err != nil {
		t.Fatal(err)
	}

	err = element.SendKeys(bundleId)
	if err != nil {
		t.Fatal(err)
	}

	err = element.Clear()
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_Rect(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(element)

	rect, err := element.Rect()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rect)
}

func TestElement_IsEnabled(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	bundleId := "com.apple.Preferences"
	_ = bundleId
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	if err != nil {
		t.Fatal(err)
	}

	isEnabled, err := element.IsEnabled()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isEnabled)
}

func TestElement_IsDisplayed(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	if err != nil {
		t.Fatal(err)
	}

	displayed, err := element.IsDisplayed()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(displayed)
}

func TestElement_IsSelected(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	if err != nil {
		t.Fatal(err)
	}

	isSelected, err := element.IsSelected()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isSelected)

	if isSelected {
		return
	}

	// iPad 左右分栏
	element, err = s.FindElement(WDALocator{Predicate: "selected == true AND label == '通用'"})
	if err != nil {
		t.Fatal(err)
	}
	isSelected, err = element.IsSelected()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isSelected)
}

func TestElement_GetAttribute(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetLabel("通用")})
	if err != nil {
		t.Fatal(err)
	}

	// attrName := "type"
	attr := NewWDAElementAttribute().SetUID("")
	// attr = NewWDAElementAttribute().SetAccessibilityContainer(false)
	value, err := element.GetAttribute(attr)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(attr.getAttributeName(), "=", value)
	t.Log("element.UID", "=", element.UID)
}

func TestElement_Text(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetLabel("通用")})
	if err != nil {
		t.Fatal(err)
	}

	// attrName := "type"
	// attrName = NewWDAElementAttribute().SetUID("").GetAttributeName()
	// attrName = NewWDAElementAttribute().SetAccessibilityContainer(false).GetAttributeName()
	text, err := element.Text()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(text)
}

func TestElement_Type(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetLabel("通用")})
	if err != nil {
		t.Fatal(err)
	}

	elemType, err := element.Type()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(elemType)
	t.Log(elemType == WDAElementType{Cell: true}.String())
	t.Log(elemType == fmt.Sprintf("%s", WDAElementType{StaticText: true}))
}

func TestElement_FindElement(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetLabel("通用")})
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	subElement, err := element.FindElement(WDALocator{ClassName: WDAElementType{Image: true}})
	if err != nil {
		t.Fatal(err)
	}

	err = subElement.ScreenshotToDisk("/Users/hero/Desktop/e2.png")
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_FindElements(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetLabel("通用")})
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	subElements, err := element.FindElements(WDALocator{Predicate: "value != 'abc123'"})
	if err != nil {
		t.Fatal(err)
	}

	for i := range subElements {
		err = subElements[i].ScreenshotToDisk("/Users/hero/Desktop/es" + strconv.FormatInt(int64(i), 10) + ".png")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestElement_FindVisibleCells(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	// bundleId = "com.apple.mobilenotes"
	_ = s.AppLaunch(bundleId)
	element, err := s.FindElement(WDALocator{ClassName: WDAElementType{Table: true}})
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	elemCells, err := element.FindVisibleCells()
	if err != nil {
		t.Fatal(err)
	}

	for i := range elemCells {
		text, err := elemCells[i].Text()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(text)
	}
}

func TestElement_Screenshot(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeCell' AND name == '通知'"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(element)

	_, err = element.Screenshot()
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_ScreenshotToImage(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	// element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeCell' AND name == '通知'"})
	if err != nil {
		t.Fatal(err)
	}

	// toPng, err := element.ScreenshotToJpeg()
	img, format, err := element.ScreenshotToImage()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("元素图片的格式:", format)
	t.Log("元素图片的大小:", img.Bounds().Size())
	t.Log(element.Rect())
}

func TestElement_ScreenshotToDisk(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	// element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetValue("通知")})
	element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeCell' AND name == '通知'"})
	if err != nil {
		t.Fatal(err)
	}

	// err = element.ScreenshotToDiskAsJpeg("/Users/hero/Desktop/e1.png")
	err = element.ScreenshotToDisk("/Users/hero/Desktop/e1.png")
	if err != nil {
		t.Fatal(err)
	}
}

func TestElement_IsAccessible(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeCell' AND name == '通知'"})
	if err != nil {
		t.Fatal(err)
	}

	isAccessible, err := element.IsAccessible()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isAccessible)

	element, err = s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeStaticText' AND name == '通知'"})
	if err != nil {
		t.Fatal(err)
	}

	isAccessible, err = element.IsAccessible()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isAccessible)
}

func TestElement_IsAccessibilityContainer(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	_ = s.AppLaunch(bundleId)
	Debug = true
	element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeCell' AND name == '通知'"})
	if err != nil {
		t.Fatal(err)
	}

	isAccessibilityContainer, err := element.IsAccessibilityContainer()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isAccessibilityContainer)

	element, err = s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeStaticText' AND name == '通知'"})
	if err != nil {
		t.Fatal(err)
	}

	isAccessibilityContainer, err = element.IsAccessibilityContainer()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isAccessibilityContainer)
}

func TestElement_Tmp(t *testing.T) {
	c, err := NewClient(deviceURL)
	if err != nil {
		t.Fatal(err)
	}
	_ = c.Unlock()
	s, err := c.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	Debug = true
	element, err := s.FindElement(WDALocator{ClassName: WDAElementType{PickerWheel: true}})
	// element, err := s.FindElement(WDALocator{ClassName: WDAElementType{Table: true}})
	// element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeCell' AND name == '手机淘宝'"})
	// element, err := s.FindElement(WDALocator{Name: "自定手势作用区域"})
	// element, err := s.FindElement(WDALocator{ClassName: WDAElementType{Image: true}})
	// element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeIcon' AND visible == true"})
	// element, err := s.FindElement(WDALocator{Predicate: "type == 'XCUIElementTypeCell' AND name == '灵敏度评估'"})
	// element, err := s.FindElement(WDALocator{LinkText: NewWDAElementAttribute().SetLabel("通用")})
	// element, err := s.FindElement(WDALocator{ClassName: WDAElementType{Table: true}})
	// element, err := s.FindElement(WDALocator{ClassName: WDAElementType{StatusBar: true}})
	// element, err := s.FindElement(WDALocator{Predicate: "selected == true AND label == '通用'"})
	// element, err := s.FindElements(WDALocator{LinkText: NewWDAElementAttribute().SetLabel("通用")})
	if err != nil {
		t.Fatal(err)
	}

	// elemType, err := element.Type()
	// fmt.Println(err, elemType)
	// text, _ := element.Text()
	// t.Log(text)
	// element.ScrollUp(10)
	// return
	element.tttTmp()
	// text, _ = element.Text()
	// t.Log(text)
}
