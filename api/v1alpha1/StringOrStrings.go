package v1alpha1

import "encoding/json"

// StringOrStrings type accepts one string or multiple strings
// +kubebuilder:object:generate=false
type StringOrStrings []string

// MarshalJSON function is a custom implementation of json.Marshal for StringOrStrings
func (s StringOrStrings) MarshalJSON() ([]byte, error) {
	//This is going to be tricky
	//if len(s) == 1 {
	//	return json.Marshal(s[0])
	//}
	//I need to convert it to string array
	// if i use json.Marshal(s) here it is going to go into infinite loop
	// since json.Marshal for type StringOrStrings are overwritten in this very own method
	var k []string
	for _, str := range s {
		k = append(k, str)
	}
	return json.Marshal(k)
}

// UnmarshalJson function is a custom implementation of json to unmarshal StringOrStrings
func (s *StringOrStrings) UnmarshalJSON(b []byte) error {
	//Try to convert to array
	var strings []string
	if err := json.Unmarshal(b, &strings); err != nil {
		//If err, convert it to string and add it to array
		var str string
		err = json.Unmarshal(b, &str)
		if err != nil {
			return err
		}
		strings = []string{str}

	}

	*s = strings
	return nil
}
