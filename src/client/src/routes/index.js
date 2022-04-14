import VueRouter from "vue-router";
import HomePage from "@/pages/HomePage";
import NotFound from "@/pages/NotFound";

const routes = [
  {
    name: "home",
    path: "",
    component: HomePage,
    meta: {
      title: "Home"
    }
  },
  {
    name: "not-found",
    path: "*",
    component: NotFound,
    meta: {
      title: "Page Not Found"
    }
  }
];

const router = new VueRouter({
  routes,
  mode: "history"
});

router.beforeEach(async (to, from, next) => {
  const hasTitle = to.matched
    .slice()
    .reverse()
    .find(r => r.meta && r.meta.title);

  if (hasTitle) {
    document.title = hasTitle.meta.title;
  }
  next();
});

export default router;
