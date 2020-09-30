import Vue from 'vue'
import Router from 'vue-router'
import Slug from '../components/Slug.vue'

Vue.use(Router)

export default new Router({
    routes: [
        {
            path: '/slug',
            name: 'Slug',
            component: Slug
        }
    ]
})