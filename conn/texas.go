package conn

import (
)

//go:generate mockgen -destination=mocks/texas.go -package=mocks . TexasClient
type TexasClient interface{
}