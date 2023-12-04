package util

type Result[T any] struct {
	Ok    bool
	Err   error
	Value T
}

func (r *Result[T]) Unwrap() (T, error) {
	if r.Ok {
		return r.Value, nil
	}
	return r.Value, r.Err
}
