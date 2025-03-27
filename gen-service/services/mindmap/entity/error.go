package entity

import "errors"

var (
	ErrMindmapDataNotValid = errors.New("mindmap data cannot be blank")
	ErrUserIdNotValid      = errors.New("user id is not valid")
	ErrMindmapNotFound     = errors.New("mindmap not found")
	ErrCannotCreateMindmap = errors.New("cannot create mindmap")
	ErrCannotUpdateMindmap = errors.New("cannot update mindmap")
	ErrCannotDeleteMindmap = errors.New("cannot delete mindmap")
	ErrCannotListMindmap   = errors.New("cannot list mindmaps")
	ErrCannotGetMindmap    = errors.New("cannot get mindmap details")
	ErrRequesterIsNotOwner = errors.New("no permission, only mindmap owner can do this")
)
