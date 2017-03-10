package mxj

import "strings"

// Removes the path.
func (mv Map) Remove(path string) error {
	m := map[string]interface{}(mv)
	return remove(m, path)
}

//Created by rgonzalez
//Useful when eliminating from large structs
//RemoveByKey is a Middle Function between Remove and Path
func (mv Map) RemoveByKey(key string) error {
	path := mv.PathForKeyShortest(key)
	err := mv.Remove(path)
	return err
}

func remove(m interface{}, path string) error {
	val, err := prevValueByPath(m, path)
	if err != nil {
		return err
	}

	lastKey := lastKey(path)
	delete(val, lastKey)

	return nil
}

// returns the last key of the path.
// lastKey("a.b.c") would had returned "c"
func lastKey(path string) string {
	keys := strings.Split(path, ".")
	key := keys[len(keys)-1]
	return key
}

// returns the path without the last key
// parentPath("a.b.c") whould had returned "a.b"
func parentPath(path string) string {
	keys := strings.Split(path, ".")
	parentPath := strings.Join(keys[0:len(keys)-1], ".")
	return parentPath
}
