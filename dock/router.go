package dock

import (
	"errors"
	"fmt"
	"path"
)

type Router interface {
	Navigate(path string) error
	NavigateBack() error
	GetCurrentComponent() Component
}

type RouterImpl struct {
	mComponents map[string]Component
	mCurrComponent Component
	mCurrPath string
}

func NewDefaultRouter(components map[string]Component) Router {
	router := &RouterImpl{
		mComponents: components,
	}
	router.Navigate("/")
	return router
}

func (router *RouterImpl) Navigate(path string) error{
	component, ok := router.mComponents[path]
	if !ok {
		return errors.New(fmt.Sprintf("Undefined path %s", path))
	}
	router.mCurrPath = path
	router.mCurrComponent = component
	return nil
}

func (router *RouterImpl) NavigateBack() error{
	backPath := path.Dir(router.mCurrPath)
	component, ok := router.mComponents[backPath]
	if !ok {
		return errors.New(fmt.Sprintf("No parent component for current path %s", router.mCurrPath))
	}
	router.mCurrPath = backPath
	router.mCurrComponent = component
	return nil
}

func (router *RouterImpl) GetCurrentComponent() Component{
	return router.mCurrComponent
}