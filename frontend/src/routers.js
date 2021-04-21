import Login from "@/components/Login";
import Register from "@/components/Register";
import TorrentTable from "@/components/TorrentTable";
import PubTorrent from "@/components/PubTorrent";

const routers = [
    {
        path: '/login',
        name: 'login',
        component: Login
    },
    {
        path: '/register',
        name: 'register',
        component: Register,
    },
    {
        path: '/torrents',
        name: "torrents",
        component: TorrentTable,
    },
    {
        path: "/pubTorrent",
        name: "pubTorrent",
        component: PubTorrent,
    },
    {
        path: '/',
        redirect: "login",
    },
]

export default routers
