package models_boil

import "strconv"

func GetUsers(users UserSlice) string {
	res := ""
	for _, v := range users {
		res += strconv.Itoa(v.ID)
		res += ","
		res += v.Name.String
		res += "\n"
	}
	return res
}

