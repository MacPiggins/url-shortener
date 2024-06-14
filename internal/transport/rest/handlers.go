package rest

import (
	"url-shortener/internal/database"
	"url-shortener/pkg/tokeniser"
)

func CreateToken(link string, db database.Wrapper) (string, error) {
	token := tokeniser.GenerateToken(link)
	err := db.Set(token, link)
	if err != nil {
		if _, ok := err.(database.UniqueError); ok {

			someLink, err := RetrieveLink(token, db)
			if err != nil {
				return "", err
			}

			if someLink == link {
				return token, nil
			} else {
				return CreateToken(token, db)
			}

		} else {
			return "", err
		}
	}
	return token, err
}

func RetrieveLink(token string, db database.Wrapper) (string, error) {
	link, err := db.Get(token)
	if err != nil {
		if _, ok := err.(database.NotFoundError); ok {
			// TODO: return something about 404
			return "", err
		} else {
			return "", err
		}
	}
	return link, err
}
