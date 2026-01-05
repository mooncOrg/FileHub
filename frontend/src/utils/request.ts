import axios from 'axios'
import { toast } from 'vue-sonner'

const service = axios.create({
    baseURL: '/api',
    timeout: 60000,
})

service.interceptors.request.use(
    (config) => {
        return config
    },
    (error) => {
        return Promise.reject(error)
    }
)

service.interceptors.response.use((response) => {
    const res = response.data
    if (res.code === 500) {
        toast.error(res.msg)
        return Promise.reject(new Error(res.msg))
    }
    return res
})

export default service