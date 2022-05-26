import service from '@/utils/request'

// 获取services list
export const getServicesList = (data) => {
  return service({
    url: '/services/getServicesList',
    method: 'post',
    data
  })
}

// 获取services in raw
export const getServicesRaw = (data) => {
  return service({
    url: '/services/getServicesRaw',
    method: 'post',
    data
  })
}

// 获取services in detail
export const getServicesDetail = (data) => {
  return service({
    url: '/services/getServicesDetail',
    method: 'post',
    data
  })
}

// 获取services关联pods
export const getServicesPods = (data) => {
  return service({
    url: '/services/getServicesPods',
    method: 'post',
    data
  })
}

// 删除services
export const deleteServices = (data) => {
  return service({
    url: '/services/deleteServices',
    method: 'post',
    data
  })
}
