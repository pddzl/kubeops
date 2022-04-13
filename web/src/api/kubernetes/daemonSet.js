import service from '@/utils/request'

// 获取daemonSet list
export const getDaemonSetList = (data) => {
  return service({
    url: '/daemonSet/getDaemonSetList',
    method: 'post',
    data
  })
}

// 获取daemonSet in raw
export const getDaemonSetRaw = (data) => {
  return service({
    url: '/daemonSet/getDaemonSetRaw',
    method: 'post',
    data
  })
}

// 获取daemonSet in detail
export const getDaemonSetDetail = (data) => {
  return service({
    url: '/daemonSet/getDaemonSetDetail',
    method: 'post',
    data
  })
}
