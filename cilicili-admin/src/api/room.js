import request from '../utils/request'

export function SaveRoom(room) {
    return request({
        url: '/room-service/create',
        method: 'post',
        data: room
    })
}

export function FindAllRooms() {
    return request({
        url: '/room-service/findAll',
        method: 'get'
    })
}

export function FindAllRoomsByLiving(living) {
    return request({
        url: '/room-service/findByLiving/' + living,
        method: 'get'
    })
}

export function FindAllRoomsByStatus(status) {
    return request({
        url: '/room-service/findByStatus/' + status,
        method: 'get'
    })
}