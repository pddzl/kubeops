import service from '@/utils/request'

// 获取secret list
export const getSecretList = (data) => {
  return service({
    url: '/secret/getSecretList',
    method: 'post',
    data
  })
}

// 获取secret编排
export const getSecretRaw = (data) => {
  return service({
    url: '/secret/getSecretRaw',
    method: 'post',
    data
  })
}

// 获取secret详情
export const getSecretDetail = (data) => {
  return service({
    url: '/secret/getSecretDetail',
    method: 'post',
    data
  })
}
