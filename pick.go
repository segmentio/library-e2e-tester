package tester

func pickID(msg map[string]interface{}, key string) (string, bool) {
	a, ok := msg[key].(map[string]interface{})
	if !ok {
		return "", false
	}
	id, ok := a["id"]
	if !ok {
		return "", false
	}
	v, ok := id.(string)
	return v, ok
}
