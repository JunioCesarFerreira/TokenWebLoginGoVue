package data

type RepositoryDummy struct {
}

func NewDummy() RepositoryDummy {
	return RepositoryDummy{}
}

func (r RepositoryDummy) CheckUser(id string, hash string) (bool, int, error) {
	// Test: pass1234
	// Hash generated using: https://emn178.github.io/online-tools/sha512.html
	exampleHashValue := "b66dd5a7a689f88e302ab2ae4a9567f9c7572c18e520b3bf712bb2630b3931a503d647baedf48df470006312d07984216578b60526e5ee6137ef1fd215190a0c"
	hash1234 := "d404559f602eab6fd602ac7680dacbfaadd13630335e951f097af3900e9de176b6db28512f2e000b9d04fba5133e8b1c6e8df59db3a8ab9d60be4b97cc9e81db"
	if id == "Junio" && hash == exampleHashValue {
		return true, 42, nil
	} else if id == "Carlos" && hash == hash1234 {
		return true, 24, nil
	} else {
		return false, 0, nil
	}
}
