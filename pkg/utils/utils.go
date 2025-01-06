//When we will be righting for book controllers we will need unmarshal data i.e.
//data in go format , in struct but user will request it in JSON format
//hence we will unmarshal that data here.

package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	// interface is defined empty , hence it is flexiible and can unmarshal JSON data
	//into any provided variable , regardless of its type
	if body, err := io.ReadAll(r.Body); err == nil {
		//when calling parsebody function makke sure to pass x as address because json.unmarshall function requires
		//a pointer to the destianation variable.
		//for e.g. var user User
		//------->ParseBody(r,&user)<--------
		// without pointer it will show runtime error
		if err := json.Unmarshal(([]byte(body)), x); err != nil {
			return
		}
	}
}
