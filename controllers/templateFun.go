package controllers

func TemplateFunIsDisable(choice int, curPage int, allPage int) string {
	if choice == 0 && curPage == 1 {
		return "disabled"
	} else if choice == 1 && curPage == allPage {
		return "disabled"
	}
	return ""
}

func TemplateFunCheckString(str string) bool {
	if len(str) > 0 {
		return true
	}
	return false
}

func TemplateAdd(a int, b int) int {
	return a + b
}

func TemplateSub(a int, b int) int {
	return a - b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
