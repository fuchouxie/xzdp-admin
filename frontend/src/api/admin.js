import request from '@/utils/request'

export function getAdminList(params) {
  return request({
    url: '/admin/AdminList',
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
export function addAdminUser(params) {
  return request({
    url: '/admin/Register',
    method: 'post',
    params: params
  })
}
