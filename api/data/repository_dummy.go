package data

type RepositoryDummy struct {
}

func NewDummy() RepositoryDummy {
	return RepositoryDummy{}
}

func (r RepositoryDummy) CheckUser(id string, hash string) (bool, error) {
	// Test: pass1234
	// Hash generated using: https://emn178.github.io/online-tools/sha512.html
	exampleHashValue := "b66dd5a7a689f88e302ab2ae4a9567f9c7572c18e520b3bf712bb2630b3931a503d647baedf48df470006312d07984216578b60526e5ee6137ef1fd215190a0c"
	if id == "Junio" && hash == exampleHashValue {
		return true, nil
	} else {
		return false, nil
	}
}
