import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import DeckView from "@/views/DeckView.vue";
import CreateDeckView from "@/views/CreateDeckView.vue";
import EditDeckView from "@/views/EditDeckView.vue";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/deck/:id/:name",
      component: DeckView,
      props: true
    },
    {
      path: "/create-deck",
      component: CreateDeckView,
    },
    {
      path: "/edit-deck/:id/:name",
      component: EditDeckView,
      props: true
    }
  ],
});

export default router;
