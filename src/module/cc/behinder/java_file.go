package behinder

import "io/ioutil"

type JavaBehinder struct {
	*Behinder
}

func (cc *JavaBehinder) Generate(mode GenerateMode, arg string) (content []byte) {
	content, err := ioutil.ReadFile("external/module/cc/behinder/behinder.enc.jsp")
	if err != nil {
		return nil
	}

	if arg != "" {
		ioutil.WriteFile(arg, content, 0777)
	}
	return nil
}
