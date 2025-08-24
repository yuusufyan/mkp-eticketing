package controllers

import "auth-rbac/src/services"

type MenuController struct {
	menuService services.MenuService
}

func NewMenuController(menuService services.MenuService) *MenuController {
	return &MenuController{menuService: menuService}
}
