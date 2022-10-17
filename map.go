package conv

import (
	"strings"
)

func ToEnvMap(m map[string]any) map[string]any {
	res := make(map[string]any, len(m))
	for k, v := range m {
		if m1, ok := v.(map[string]any); ok {
			m2 := ToEnvMap(m1)
			for k2, v2 := range m2 {
				res[ToEnvKey(k+"."+k2)] = v2
			}
		} else {
			res[k] = v
		}
	}
	return res
}

func OSEnvsToEnvMap(envs []string) map[string]string {
	m := make(map[string]string, len(envs))
	for _, pair := range envs {
		for i, c := range pair {
			if c == '=' {
				m[ToEnvKey(pair[:i])] = pair[i+1:]
			}
		}
	}
	return m
}

func OSArgsToEnvMap(args []string) map[string]string {
	m := make(map[string]string, len(args))
	var key string
	for _, arg := range args {
		if arg[0] != '-' {
			if key != "" {
				m[ToEnvKey(key)] = arg
				key = ""
			} else {
				m[arg] = ""
			}
		} else {
			key = ""
			j := 0
			for j < len(arg) && arg[j] == '-' {
				j++
			}

			key = arg[j:]
			for k, c := range key {
				if c == '=' {
					m[ToEnvKey(key[:k])] = key[k+1:]
					key = ""
					break
				}
			}
		}
	}

	if key != "" {
		m[ToEnvKey(key)] = ""
	}

	return m
}

func ToEnvKey(k string) string {
	return strings.ReplaceAll(strings.ToLower(k), "_", ".")
}

func GetMapKeys[K comparable, V any](m map[K]V) []K {
	a := make([]K, 0, len(m))
	for k := range m {
		a = append(a, k)
	}
	return a
}

func GetMapValues[K comparable, V any](m map[K]V) []V {
	a := make([]V, 0, len(m))
	for _, v := range m {
		a = append(a, v)
	}
	return a
}

func GetKeysAndValues[K comparable, V any](m map[K]V) ([]K, []V) {
	kl := make([]K, 0, len(m))
	vl := make([]V, 0, len(m))
	for k, v := range m {
		kl = append(kl, k)
		vl = append(vl, v)
	}
	return kl, vl
}
