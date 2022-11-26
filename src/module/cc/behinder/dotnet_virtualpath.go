package behinder

import "io/ioutil"

type DotnetVirtualPathBehinder struct {
}

func (cc *DotnetVirtualPathBehinder) Generate() (content []byte) {
	content, err := ioutil.ReadFile("external/module/cc/behinder/dotnet_virtualpath.enc.cs")
	if err != nil {
		return nil
	}
	return content
}
