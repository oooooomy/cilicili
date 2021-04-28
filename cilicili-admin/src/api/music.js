import request from '../utils/request'

export function SaveMusic(data) {
    return request({
        url: '/upload-service/music',
        method: 'post',
        data: data
    })
}

export function FindAllMusic() {
    return request({
        url: '/upload-service/music',
        method: 'get'
    })
}

export function DeleteMusic(id) {
    return request({
        url: '/upload-service/music/' + id,
        method: 'delete'
    })
}