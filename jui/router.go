package jui

import (
	"errors"
	"fmt"
	"path"

	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=mocks/router.go -package=mocks . Router
type Router interface {
	Navigate(path string, args interface{}) error
	NavigateBack() error
	GetCurrentPath() string
	GetCurrentComponent() Component
	RegisterComponent(path string, component Component) error
	UnRegisterComponent(path string)
}

type DefaultRouter struct {
	mComponents map[string]Component
	mCurrComponent Component
	mCurrPath string
}

func NewDefaultRouter() *DefaultRouter {
	dr := &DefaultRouter{
		mComponents: make(map[string]Component),
		mCurrComponent: nil,
		mCurrPath: "",
	}
	return dr
}

func (dr *DefaultRouter) Navigate(path string, args interface{}) error{
	log.Debugf("Router Navigate: from %s to %s", dr.mCurrPath, path)
	component, ok := dr.mComponents[path]
	if !ok {
		return errors.New(fmt.Sprintf("Undefined path %s", path))
	}
	oldComponent := dr.mCurrComponent
	if oldComponent != nil {
		oldComponent.WillUnmount()
	}
	component.WillMount(args)
	dr.mCurrPath = path
	dr.mCurrComponent = component
	if oldComponent != nil {
		oldComponent.DidUnmount()
	}
	component.DidMount()
	return nil
}

func (dr *DefaultRouter) NavigateBack() error{
	log.Debugf("Router Navigate Back Curr(%s)", dr.mCurrPath)
	if dr.mCurrPath == "/"{
		return nil
	}
	backPath := path.Dir(dr.mCurrPath)
	_, ok := dr.mComponents[backPath]
	if !ok {
		return errors.New(fmt.Sprintf("No parent component for current path %s", dr.mCurrPath))
	}
	return dr.Navigate(backPath, nil)
}

func (dr *DefaultRouter) GetCurrentPath() string{
	return dr.mCurrPath
}

func (dr *DefaultRouter) GetCurrentComponent() Component{
	return dr.mCurrComponent
}

func (dr *DefaultRouter) RegisterComponent(path string, component Component) error{
	log.Debugf("Router Register Component (%s)", path)

	_, ok := dr.mComponents[path]
	if ok {
		return errors.New(fmt.Sprintf("Path(%s) is registered", path))
	}
	dr.mComponents[path] = component
	return nil
}

func (dr *DefaultRouter) UnRegisterComponent(path string){
	delete(dr.mComponents, path)
}