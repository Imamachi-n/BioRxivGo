import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    test: "",
  },
  mutations: {
    testMutation(state, test) {
      state.test = test;
    },
  },
  actions: {
    testAction({ commit }, test) {
      commit("testMutation", test);
    },
  },
  getters: {
    getTest(state) {
      return state.test;
    },
  },
});
