import request from "axios"
import {message} from 'ant-design-vue'

export const BASE_GATEWAY_URL = "http://localhost:8762"

const service = request.create({
    baseURL: BASE_GATEWAY_URL,
    timeout: 50000,
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    },
});

service.interceptors.response.use(
    response => {

        const res = response.data;

        //判断response状态
        if (!res.status) message.error('Error: ' + res.msg)

        return res
    },

    error => {
        message.error(error)
        console.log(error)
        return Promise.reject(error)
    }
);

export default service