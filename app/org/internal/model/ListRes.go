package model

type ListReq[T any] struct {
	List  []*T
	Total int
}
