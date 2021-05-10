import request from '../utils/request'

export function AdminLogin(data) {
    return request({
        url: '/account-service/admin/loginByPassword',
        method: 'post',
        data: data
    })
}

export function AdminLoginByCode(email, code) {
    return request({
        url: '/account-service/admin/loginByEmail?email=' + email + '&value=' + code,
        method: 'get'
    })
}

export function GetEmailCode(email) {
    return request({
        url: '/email-service/code?email=' + email,
        method: 'get'
    })
}