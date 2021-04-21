let JwtToken = {
    setToken(token){
        localStorage.setItem("token", token)
        console.log(token)
    },
    getToken(){
        console.log("getToken: ", localStorage.getItem("token"))
        return localStorage.getItem("token")
    },
    hasToken(){
        console.log("hasToken:", localStorage.getItem("token"))
        return localStorage.getItem("token") !== null
    },
}

export default JwtToken