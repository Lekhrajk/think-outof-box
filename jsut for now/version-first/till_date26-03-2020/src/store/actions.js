/*=========================================================================================
  File Name: actions.js
  Description: Vuex Store - actions
  ----------------------------------------------------------------------------------------
  Item Name: Vuexy - Vuejs, HTML & Laravel Admin Dashboard Template
  Author: Pixinvent
  Author URL: http://www.themeforest.net/user/pixinvent
==========================================================================================*/

const actions = {

    // /////////////////////////////////////////////
    // COMPONENTS
    // /////////////////////////////////////////////

    // Vertical NavMenu
    updateVerticalNavMenuWidth({ commit }, width) {
      commit('UPDATE_VERTICAL_NAV_MENU_WIDTH', width)
    },

    // VxAutoSuggest
    updateStarredPage({ commit }, payload) {
      commit('UPDATE_STARRED_PAGE', payload)
    },

    // The Navbar
    arrangeStarredPagesLimited({ commit }, list) {
      commit('ARRANGE_STARRED_PAGES_LIMITED', list)
    },
    arrangeStarredPagesMore({ commit }, list) {
      commit('ARRANGE_STARRED_PAGES_MORE', list)
    },

    // /////////////////////////////////////////////
    // UI
    // /////////////////////////////////////////////

    toggleContentOverlay({ commit }) {
      commit('TOGGLE_CONTENT_OVERLAY')
    },
    updateTheme({ commit }, val) {
      commit('UPDATE_THEME', val)
    },

    // /////////////////////////////////////////////
    // User/Account
    // /////////////////////////////////////////////

    updateUserInfo({ commit }, payload) {
      commit('UPDATE_USER_INFO', payload)
    },
    // /////////////////////////////////////////////
    // API CALLS
    // /////////////////////////////////////////////
            // API call for performing action like poweroff, reboot,suspend,resume etc
    // actionVm(context,payload) 
    //     {
    //       let temp_id = id;
    //       var actiondata = new FormData();
    //       actiondata.set('type', name);
    //       actiondata.set('name', name);
    //       const tokenurl = '/api/vms/'+temp_id+'/actions';
    //         axios({
    //           method: 'post',
    //           url: tokenurl,
    //           data: actiondata,
    //         })
    //         .then((response) => {
    //           this.$vs.notify({
    //           color: "dark",
    //           title: 'Server reponse',
    //           text: response.data.message,
    //         })
    //         }) 
    //         .catch((error) => {
    //           this.$vs.notify({  
    //           color: 'danger',
    //           title: "Status",
    //           text: error.response.data.error
    //         })  
    //         }),
    //         this.$vs.loading({
    //           container: '#div-with-loading',
    //           scale: 0.6
    //         })
    //         setTimeout( ()=> {
    //           this.$vs.loading.close('#div-with-loading > .con-vs-loading')
    //         }, 3000);
    //     },

}

export default actions
