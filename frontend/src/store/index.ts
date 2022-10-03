import { createStore } from "vuex";

export default createStore({
  state: {
    sales: [],
    data: [],
    producer: "",
    isLoading: false,
    showTable: false,
  },

  getters: {
    getIsLoading(state) {
      return state.isLoading;
    },
    getShowTable(state) {
      return state.showTable;
    },
    getSales(state) {
      return state.sales;
    },
    getProducer(state) {
      return state.producer;
    },
    getData(state) {
      return state.data;
    },
  },
  mutations: {
    SET_LOADING(state) {
      state.isLoading = !state.isLoading;
    },
    SET_SHOWTABLE(state) {
      state.showTable = !state.showTable;
    },
    SET_SALES(state, payload) {
      state.sales = payload;
    },
    SET_PRODUCER(state, payload) {
      state.producer = payload;
    },
    SET_DATA(state, payload) {
      state.data = payload;
    },
  },
  actions: {
    setSales(context, payload) {
      context.commit("SET_SALES", payload);
    },
    setLoading(context) {
      context.commit("SET_LOADING");
    },
    setShowTable(context) {
      context.commit("SET_SHOWTABLE");
    },
    setData(context, payload) {
      context.commit("SET_DATA", payload);
    },
    setProducer(context, payload) {
      context.commit("SET_PRODUCER", payload);
    },
  },
  modules: {},
});
