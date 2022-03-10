import service from '@/utils/request'

// 获取pod list
export const getPodList = (data) => {
  return service({
    url: '/pod/getPodList',
    method: 'post',
    data
  })
}

// 获取pod详情
export const getPodDetail = (data) => {
  return service({
    url: '/pod/getPodDetail',
    method: 'post',
    data
  })
}

// 获取pod详情
export const getPodRaw = (data) => {
  return service({
    url: '/pod/getPodRaw',
    method: 'post',
    data
  })
}

// 获取pod日志
export const getPodLog = (data) => {
  return service({
    url: '/pod/getPodLog',
    method: 'post',
    data
  })
}
