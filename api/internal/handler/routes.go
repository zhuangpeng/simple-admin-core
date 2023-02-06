// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	api "github.com/suyuan32/simple-admin-core/api/internal/handler/api"
	authority "github.com/suyuan32/simple-admin-core/api/internal/handler/authority"
	captcha "github.com/suyuan32/simple-admin-core/api/internal/handler/captcha"
	core "github.com/suyuan32/simple-admin-core/api/internal/handler/core"
	department "github.com/suyuan32/simple-admin-core/api/internal/handler/department"
	dictionary "github.com/suyuan32/simple-admin-core/api/internal/handler/dictionary"
	member "github.com/suyuan32/simple-admin-core/api/internal/handler/member"
	memberrank "github.com/suyuan32/simple-admin-core/api/internal/handler/memberrank"
	menu "github.com/suyuan32/simple-admin-core/api/internal/handler/menu"
	oauth "github.com/suyuan32/simple-admin-core/api/internal/handler/oauth"
	position "github.com/suyuan32/simple-admin-core/api/internal/handler/position"
	role "github.com/suyuan32/simple-admin-core/api/internal/handler/role"
	token "github.com/suyuan32/simple-admin-core/api/internal/handler/token"
	user "github.com/suyuan32/simple-admin-core/api/internal/handler/user"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/role/create_or_update",
					Handler: role.CreateOrUpdateRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/delete",
					Handler: role.DeleteRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/list",
					Handler: role.GetRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/status",
					Handler: role.UpdateRoleStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/change-password",
					Handler: user.ChangePasswordHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/info",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/create_or_update",
					Handler: user.CreateOrUpdateUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/list",
					Handler: user.GetUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/delete",
					Handler: user.DeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/batch_delete",
					Handler: user.BatchDeleteUserHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/perm",
					Handler: user.GetUserPermCodeHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/profile",
					Handler: user.GetUserProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/profile",
					Handler: user.UpdateUserProfileHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/status",
					Handler: user.UpdateUserStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/user/logout",
					Handler: user.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/menu/create_or_update",
					Handler: menu.CreateOrUpdateMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/delete",
					Handler: menu.DeleteMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menu/list",
					Handler: menu.GetMenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/menu/role",
					Handler: menu.GetMenuByRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/param/create_or_update",
					Handler: menu.CreateOrUpdateMenuParamHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/param/list",
					Handler: menu.GetMenuParamListByMenuIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/param/delete",
					Handler: menu.DeleteMenuParamHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/captcha",
				Handler: captcha.GetCaptchaHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/create_or_update",
					Handler: api.CreateOrUpdateApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/delete",
					Handler: api.DeleteApiHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/list",
					Handler: api.GetApiListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/authority/api/create_or_update",
					Handler: authority.CreateOrUpdateApiAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/api/role",
					Handler: authority.GetApiAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/menu/create_or_update",
					Handler: authority.CreateOrUpdateMenuAuthorityHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/authority/menu/role",
					Handler: authority.GetMenuAuthorityHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/dict/create_or_update",
					Handler: dictionary.CreateOrUpdateDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dict/delete",
					Handler: dictionary.DeleteDictionaryHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dict/list",
					Handler: dictionary.GetDictionaryListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dict/detail/create_or_update",
					Handler: dictionary.CreateOrUpdateDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dict/detail/delete",
					Handler: dictionary.DeleteDictionaryDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/dict/detail/list",
					Handler: dictionary.GetDetailByDictionaryNameHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/oauth/login",
				Handler: oauth.OauthLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/oauth/login/callback",
				Handler: oauth.OauthCallbackHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/oauth/provider/create_or_update",
					Handler: oauth.CreateOrUpdateProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth/provider/delete",
					Handler: oauth.DeleteProviderHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oauth/provider/list",
					Handler: oauth.GetProviderListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/token/create_or_update",
					Handler: token.CreateOrUpdateTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/delete",
					Handler: token.DeleteTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/batch_delete",
					Handler: token.BatchDeleteTokenHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/list",
					Handler: token.GetTokenListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/status",
					Handler: token.UpdateTokenStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/token/logout",
					Handler: token.LogoutHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/core/health",
				Handler: core.HealthCheckHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/core/init/database",
				Handler: core.InitDatabaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/department/create_or_update",
					Handler: department.CreateOrUpdateDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/delete",
					Handler: department.DeleteDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/list",
					Handler: department.GetDepartmentListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/batch_delete",
					Handler: department.BatchDeleteDepartmentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/department/status",
					Handler: department.UpdateDepartmentStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/position/create_or_update",
					Handler: position.CreateOrUpdatePositionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/delete",
					Handler: position.DeletePositionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/list",
					Handler: position.GetPositionListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/batch_delete",
					Handler: position.BatchDeletePositionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/position/status",
					Handler: position.UpdatePositionStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/member/login",
				Handler: member.MemberLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/member/register",
				Handler: member.MemberRegisterHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/member/create_or_update",
					Handler: member.CreateOrUpdateMemberHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member/delete",
					Handler: member.DeleteMemberHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member/list",
					Handler: member.GetMemberListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member/batch_delete",
					Handler: member.BatchDeleteMemberHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member/status",
					Handler: member.UpdateMemberStatusHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/member_rank/create_or_update",
					Handler: memberrank.CreateOrUpdateMemberRankHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member_rank/delete",
					Handler: memberrank.DeleteMemberRankHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member_rank/list",
					Handler: memberrank.GetMemberRankListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/member_rank/batch_delete",
					Handler: memberrank.BatchDeleteMemberRankHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
