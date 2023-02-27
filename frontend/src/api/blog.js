import request from '@/utils/request'

export function getBlogList(params = {}) {
  return request({
    url: '/blog/BlogList',
    method: 'get',
    params
  })
}

export function getBlogInfo(params = {}) {
  return request({
    url: '/blog/GetBlogInfo',
    method: 'get',
    params
  })
}

export function remove(params = {}) {
  return request({
    url: '/blog/Remove',
    method: 'post',
    params
  })
}

export function batchRemove(params = {}) {
  return request({
    url: '/blog/BatchRemove',
    method: 'post',
    params
  })
}

