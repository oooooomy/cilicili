import request from '../utils/request'

export function FindVideoByStatus(status) {
    return request({
        url: '/upload-service/video/findByStatus?status=' + status,
        method: 'get'
    })
}

export function SetVideoStatus(id, status) {
    return request({
        url: '/upload-service/video/' + id + '/setStatus?status=' + status,
        method: 'put'
    })
}