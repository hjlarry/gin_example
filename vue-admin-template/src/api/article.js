import request from '@/utils/request'

export function fetchList(params) {
  return request({
    url: '/articles',
    method: 'get',
    params
  })
}

export function createArticle(data) {
  return request({
    url: '/articles',
    method: 'post',
    data
  })
}

export function deleteArticle(id) {
  return request({
    url: '/articles/' + id,
    method: 'delete',
  })
}

export function editArticle(data) {
  return request({
    url: '/articles/' + data.id,
    method: 'put',
    data
  })
}

export function fetchArticle(id) {
  return request({
    url: '/articles/' + id,
    method: 'get'
  })
}
