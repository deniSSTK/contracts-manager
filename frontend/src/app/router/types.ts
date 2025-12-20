export enum RouteName {
    LOGIN = "login",
    SIGNUP = "signup",

    DASHBOARD = 'dashboard',

    ADMIN_PANEL = 'admin-panel',
    ADMIN_PANEL_TABLE = 'admin-panel-table',
    ADMIN_PANEL_INFO = 'admin-panel-info',
    ADMIN_PANEL_NEW = 'admin-panel-new',
}

export const AuthPages: RouteName[] = [
    RouteName.LOGIN,
    RouteName.SIGNUP,
]