package tp2

import (
	"math"
	TDAPost "tp2/post"
)

func CmpAfinidad(id int) func(p1, p2 *TDAPost.Post) int {
	return func(p1, p2 *TDAPost.Post) int {
		_, idUser1 := (*p1).Creador()
		_, idUser2 := (*p2).Creador()
		if math.Abs(float64(idUser1-id)) < math.Abs(float64(idUser2-id)) {
			return 1
		}
		if math.Abs(float64(idUser1-id)) > math.Abs(float64(idUser2-id)) {
			return -1
		}
		if (*p1).IDPost() < (*p2).IDPost() {
			return 1
		}
		return -1
	}
}
