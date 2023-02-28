import request from '@/utils/request'

export function getAdminList(params) {
  return request({
    url: '/admin/AdminList',
    method: 'get',
    params: params
  })
}

export function getAdminInfo(params) {
  return request({
    url: '/admin/getAdminInfo',
    method: 'get',
    params: params
  })
}

export function removeAdminUser(params) {
  return request({
    url: '/admin/Remove',
    method: 'post',
    params: params
  })
}
export function batchRemoveAdminUser(params) {
  return request({
    url: '/admin/BatchRemove',
    method: 'post',
    params: params
  })
}
export function updateAdminUser(params) {
  return request({
    url: '/admin/Update',
    method: 'post',
    params: params
  })
}

export function changePassword(params) {
  return request({
    url: '/admin/changePassword',
    method: 'post',
    params: params
  })
}
