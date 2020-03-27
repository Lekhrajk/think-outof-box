/*=========================================================================================
  File Name: themeConfig.js
  Description: Theme configuration
  ----------------------------------------------------------------------------------------
  Item Name: Xongl user panel
  Author: Xongl cloud private limited
==========================================================================================*/

// MAIN COLORS - VUESAX THEME COLORS
let colors = {
  primary : '#007bff',
  success : '#28C76F',
  danger  : '#EA5455',
  warning : '#FF9F43',
  dark    : '#1E1E1E',
  info    : '#D7DBDD',
}

// CONFIGS
const themeConfig = {
  disableCustomizer : false,       // options[Boolean] : true, false(default)
  disableThemeTour  : false,       // options[Boolean] : true, false(default)
  footerType        : "static",    // options[String]  : static(default) / sticky / hidden
  hideScrollToTop   : false,       // options[Boolean] : true, false(default)
  mainLayoutType    : "vertical",  // options[String]  : vertical(default) / horizontal
  navbarColor       : "#007bff",      // options[String]  : HEX color / rgb / rgba / Valid HTML Color name - (default: #fff)
  navbarType        : "sticky",  // options[String]  : floating(default) / static / sticky / hidden
  routerTransition  : "zoom-fade", // options[String]  : zoom-fade / slide-fade / fade-bottom / fade / zoom-out / none(default)
  rtl               : false,       // options[Boolean] : true, false(default)
  sidebarCollapsed  : false,       // options[Boolean] : true, false(default)
  theme             : "light",     // options[String]  : "light"(default), "dark", "semi-dark"

  // Not required yet - WIP
  userInfoLocalStorageKey: "userInfo",

  // NOTE: themeTour will be disabled in screens < 1200. Please refer docs for more info.
}

import Vue from 'vue'
import Vuesax from 'vuesax'
Vue.use(Vuesax, { theme:{ colors }, rtl: themeConfig.rtl })

export default themeConfig
