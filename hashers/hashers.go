package hashers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"

	"github.com/pkg/errors"
)

type ErrorList []error

func (el ErrorList) Error() string {
	s := ""
	for i, err := range el {
		s += fmt.Sprintf("error %d: %s\n", i+1, err)
	}
	return s
}

// Get hashers for each of the hash algorithms listed. If a hash is not
// known, this function will return the hashers it does know, and an error.
func Get(algorithms []string) ([]hash.Hash, error) {
	errs := make(ErrorList, 0)
	hashers := make([]hash.Hash, 0, 2)
	for _, a := range algorithms {
		switch a {
		case "sha512":
			hashers = append(hashers, sha512.New())
		case "sha256":
			hashers = append(hashers, sha256.New())
		case "md5":
			hashers = append(hashers, md5.New())
		case "sha1":
			hashers = append(hashers, sha1.New())
		default:
			errs = append(errs, errors.Errorf("unknown hasher %s", a))
		}
	}

	if len(errs) > 0 {
		return hashers, errs
	}
	return hashers, nil
}
