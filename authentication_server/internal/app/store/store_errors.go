package store

import "errors"

var errNotSslMode error = errors.New("value is not a valid ssl mode")

var errMultipleUsersFetchedByTargetInUpdateQuery error = errors.New("multiple users are found with update query target module")

var errMultipleUsersFetchedByTargetInDeleteQuery error = errors.New("multiple users are found with delete query target module")
