package map_utils

func GetKeys[KT comparable, T any](m map[KT]T) []KT {
	var keys []KT
	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

func GetValues[KT comparable, T any](m map[KT]T) []T {
	var values []T
	for _, v := range m {
		values = append(values, v)
	}

	return values
}

func FilterFunc[KT comparable, T any](m map[KT]T, f func(KT, T) bool) map[KT]T {
	nm := make(map[KT]T, len(m))
	for kt, t := range m {
		if f(kt, t) {
			nm[kt] = t
		}
	}
	return nm
}

func Merge[KT comparable, T any](m map[KT]T, m2 ...map[KT]T) map[KT]T {
	n := Copy(m)
	for _, m3 := range m2 {
		for kt, t := range m3 {
			if _, ok := n[kt]; !ok {
				n[kt] = t
			}
		}
	}
	return n
}

func Copy[KT comparable, T any](m map[KT]T) map[KT]T {
	n := map[KT]T{}
	for kt, t := range m {
		n[kt] = t
	}
	return n
}
