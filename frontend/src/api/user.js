import request from "@/utils/request";

export function getUserList(params) {
  return request({
    url: '/user/UserList',
    method: 'get',
    params: params
  })
}

export function removeUser(params) {
  return request({
    url: '/user/Remove',
    method: 'post',
    params: params
  })
}

export function batchRemoveUser(params) {
  return request({
    url: '/user/BatchRemove',
    method: 'post',
    params: params
  })
}

export function getUserInfo(params) {
  return request({
    url: '/user/GetUserInfo',
    method: 'get',
    params: params
  })
}
