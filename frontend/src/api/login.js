import request from '@/utils/request'

export function login(username, password) {
  return request({
    url: '/admin/Login',
    method: 'post',
    data: {
      username,
      password
    }
  })
}

export function getInfo(token) {
  return request({
    url: '/admin/info',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/admin/Logout',
    method: 'post'
  })
}

export function code(phone) {
  return request({
    url: '/admin/Code',
    method: 'post',
    data : phone
  })
}

export function codeLogin(phone,code) {
  return request({
    url: '/admin/CodeLogin',
    method: 'post',
    data: {
      phone,
      code
    }
  })
}



