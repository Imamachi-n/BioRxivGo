import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    articles: [],
  },
  mutations: {
    mutateArticles(state, articles) {
      state.articles = articles;
    },
  },
  actions: {
    async getArticlesAllAction({ commit }, articles) {
      const { data } = await axios.get("http://localhost:9000/api/articles");
      console.log(data);
      commit("mutateArticles", data);
    },
  },
  getters: {
    getArticles(state) {
      return state.articles;
    },
  },
});
